package repository

import (
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/integration/mtest"
	"os"
	"testing"
)

func TestUserRepository_DeleteUser(t *testing.T) {
	databaseName := "user_database_test"
	collectionName := "user_collection_test"
	err := os.Setenv("MONGODB_USER_DB", collectionName)
	if err != nil {
		return
	}

	defer os.Clearenv()
	mtestDb := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))

	mtestDb.Run("when_sending_a_valid_userId_returns_success", func(mt *mtest.T) {
		mt.AddMockResponses(bson.D{
			{Key: "ok", Value: 1},
			{Key: "n", Value: 1},
			{Key: "acknowledged", Value: true},
		})
		database := mt.Client.Database(databaseName)

		repo := NewUserRepository(database)
		err := repo.DeleteUser("test")

		assert.Nil(t, err)

	})

	mtestDb.Run("return_error_from_database", func(mt *mtest.T) {
		mt.AddMockResponses(bson.D{
			{Key: "ok", Value: 0},
		})

		database := mt.Client.Database(databaseName)

		repo := NewUserRepository(database)
		err := repo.DeleteUser("test")

		assert.NotNil(t, err)
	})
}
