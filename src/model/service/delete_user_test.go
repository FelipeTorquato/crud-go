package service

import (
	"crud/src/configuration/rest_err"
	"crud/src/tests/mocks"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.uber.org/mock/gomock"
	"testing"
)

func TestUserDomainService_DeleteUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repository := mocks.NewMockUserRepository(ctrl)
	service := NewUserDomainService(repository)

	t.Run("when_sending_a_valid_userId_returns_success", func(t *testing.T) {
		id := primitive.NewObjectID().Hex()

		repository.EXPECT().DeleteUser(id).Return(nil)

		err := service.DeleteUser(id)
		assert.Nil(t, err)
	})

	t.Run("when_sending_an_invalid_userId_returns_error", func(t *testing.T) {
		id := primitive.NewObjectID().Hex()

		repository.EXPECT().DeleteUser(id).Return(
			rest_err.NewInternalServerError("error trying to delete user"))

		err := service.DeleteUser(id)
		assert.NotNil(t, err)
	})
}
