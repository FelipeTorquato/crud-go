package service

import (
	"crud/src/configuration/rest_err"
	"crud/src/model"
	"crud/src/tests/mocks"
	"github.com/golang-jwt/jwt/v5"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.uber.org/mock/gomock"
	"os"
	"testing"
)

func TestUserDomainService_LoginUserServices(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repository := mocks.NewMockUserRepository(ctrl)
	service := &userDomainService{repository}

	t.Run("when_calling_repository_returns_error", func(t *testing.T) {
		id := primitive.NewObjectID().Hex()

		userDomain := model.NewUserDomain("test@test.com", "test", "test", 43)
		userDomain.SetID(id)

		repository.EXPECT().FindUserByEmailAndPassword(
			userDomain.GetEmail(), gomock.Any()).Return(
			nil, rest_err.NewInternalServerError("error trying to find user by email and password"))

		user, token, err := service.LoginUserServices(userDomain)

		assert.Nil(t, user)
		assert.Empty(t, token)
		assert.NotNil(t, err)
		assert.EqualValues(t, err.Message, "error trying to find user by email and password")
	})

	t.Run("when_calling_create_token_returns_error", func(t *testing.T) {
		userDomainMock := mocks.NewMockUserDomainInterface(ctrl)

		userDomainMock.EXPECT().GetEmail().Return("test@test.com")
		userDomainMock.EXPECT().GetPassword().Return("test")
		userDomainMock.EXPECT().EncryptPassword()

		userDomainMock.EXPECT().GenerateToken().Return(
			"", rest_err.NewInternalServerError("error trying to create token"))

		repository.EXPECT().FindUserByEmailAndPassword(
			"test@test.com", "test").Return(
			userDomainMock, nil)

		user, token, err := service.LoginUserServices(userDomainMock)

		assert.Nil(t, user)
		assert.Empty(t, token)
		assert.NotNil(t, err)
		assert.EqualValues(t, err.Message, "error trying to create token")
	})

	t.Run("when_user_and_password_is_valid_return_success", func(t *testing.T) {
		id := primitive.NewObjectID().Hex()
		secret := "test"
		os.Setenv("JWT_SECRET_KEY", secret)
		defer os.Clearenv()

		userDomain := model.NewUserDomain("test@test.com", "test", "test", 43)
		userDomain.SetID(id)

		repository.EXPECT().FindUserByEmailAndPassword(
			userDomain.GetEmail(), gomock.Any()).Return(
			userDomain, nil)

		userDomainReturn, token, err := service.LoginUserServices(userDomain)

		assert.Nil(t, err)
		assert.EqualValues(t, userDomain.GetID(), id)
		assert.EqualValues(t, userDomain.GetEmail(), userDomainReturn.GetEmail())
		assert.EqualValues(t, userDomain.GetPassword(), userDomainReturn.GetPassword())
		assert.EqualValues(t, userDomain.GetName(), userDomainReturn.GetName())
		assert.EqualValues(t, userDomain.GetAge(), userDomainReturn.GetAge())

		tokenReturn, _ := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); ok {
				return []byte(secret), nil
			}

			return nil, rest_err.NewBadRequestError("invalid jwt token")
		})

		_, ok := tokenReturn.Claims.(jwt.MapClaims)
		if !ok || !tokenReturn.Valid {
			t.FailNow()
			return
		}
	})
}
