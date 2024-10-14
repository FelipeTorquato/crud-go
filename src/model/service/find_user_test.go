package service

import (
	"crud/src/configuration/rest_err"
	"crud/src/model"
	"crud/src/tests/mocks"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.uber.org/mock/gomock"
	"math/rand"
	"strconv"
	"testing"
)

func TestUserDomainService_FindUserByIDServices(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repository := mocks.NewMockUserRepository(ctrl)
	service := NewUserDomainService(repository)

	t.Run("when_exists_an_user_returns_success", func(t *testing.T) {
		id := primitive.NewObjectID().Hex()

		userDomain := model.NewUserDomain("test@test.com", "test", "test", 43)
		userDomain.SetID(id)
		repository.EXPECT().FindUserByID(id).Return(userDomain, nil)

		userDomainReturn, err := service.FindUserByIDServices(id)

		assert.Nil(t, err)
		assert.EqualValues(t, userDomain.GetID(), id)
		assert.EqualValues(t, userDomain.GetEmail(), userDomainReturn.GetEmail())
		assert.EqualValues(t, userDomain.GetPassword(), userDomainReturn.GetPassword())
		assert.EqualValues(t, userDomain.GetName(), userDomainReturn.GetName())
		assert.EqualValues(t, userDomain.GetAge(), userDomainReturn.GetAge())
	})

	t.Run("when_does_not_exists_an_user_returns_error", func(t *testing.T) {
		id := primitive.NewObjectID().Hex()

		repository.EXPECT().FindUserByID(id).Return(nil, rest_err.NewNotFoundError("user not found"))
		userDomainReturn, err := service.FindUserByIDServices(id)

		assert.NotNil(t, err)
		assert.Nil(t, userDomainReturn)
	})
}

func TestUserDomainService_FindUserByEmailServices(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repository := mocks.NewMockUserRepository(ctrl)
	service := NewUserDomainService(repository)

	t.Run("when_exists_an_user_returns_success", func(t *testing.T) {
		id := primitive.NewObjectID().Hex()
		email := "test@test.com"

		userDomain := model.NewUserDomain("test@test.com", "test", "test", 43)
		userDomain.SetID(id)
		repository.EXPECT().FindUserByEmail(email).Return(userDomain, nil)

		userDomainReturn, err := service.FindUserByEmailServices(email)

		assert.Nil(t, err)
		assert.EqualValues(t, userDomain.GetID(), id)
		assert.EqualValues(t, userDomain.GetEmail(), userDomainReturn.GetEmail())
		assert.EqualValues(t, userDomain.GetPassword(), userDomainReturn.GetPassword())
		assert.EqualValues(t, userDomain.GetName(), userDomainReturn.GetName())
		assert.EqualValues(t, userDomain.GetAge(), userDomainReturn.GetAge())
	})

	t.Run("when_does_not_exists_an_user_returns_error", func(t *testing.T) {
		email := "test@test.com"

		repository.EXPECT().FindUserByEmail(email).Return(nil, rest_err.NewNotFoundError("user not found"))
		userDomainReturn, err := service.FindUserByEmailServices(email)

		assert.NotNil(t, err)
		assert.Nil(t, userDomainReturn)
	})
}

func TestUserDomainService_FindUserByEmailAndPasswordServices(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repository := mocks.NewMockUserRepository(ctrl)
	service := &userDomainService{repository}

	t.Run("when_exists_an_user_returns_success", func(t *testing.T) {
		id := primitive.NewObjectID().Hex()
		email := "test@test.com"
		password := strconv.FormatInt(rand.Int63(), 10)

		userDomain := model.NewUserDomain("test@test.com", "test", "test", 43)
		userDomain.SetID(id)
		repository.EXPECT().FindUserByEmailAndPassword(email, password).Return(userDomain, nil)

		userDomainReturn, err := service.findUserByEmailAndPasswordServices(email, password)

		assert.Nil(t, err)
		assert.EqualValues(t, userDomain.GetID(), id)
		assert.EqualValues(t, userDomain.GetEmail(), userDomainReturn.GetEmail())
		assert.EqualValues(t, userDomain.GetPassword(), userDomainReturn.GetPassword())
		assert.EqualValues(t, userDomain.GetName(), userDomainReturn.GetName())
		assert.EqualValues(t, userDomain.GetAge(), userDomainReturn.GetAge())
	})

	t.Run("when_does_not_exists_an_user_returns_error", func(t *testing.T) {
		email := "test@test.com"
		password := strconv.FormatInt(rand.Int63(), 10)

		repository.EXPECT().FindUserByEmailAndPassword(email, password).Return(nil, rest_err.NewNotFoundError("user not found"))
		userDomainReturn, err := service.findUserByEmailAndPasswordServices(email, password)

		assert.NotNil(t, err)
		assert.Nil(t, userDomainReturn)
	})
}
