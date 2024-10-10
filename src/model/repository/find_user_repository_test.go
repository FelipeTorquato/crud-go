package repository

import (
	"crud/src/model/repository/entity"
	"fmt"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/integration/mtest"
	"os"
	"testing"
)

func TestUserRepository_FindUserByEmail(t *testing.T) {
	databaseName := "user_database_test"
	collectionName := "user_collection_test"
	err := os.Setenv("MONGODB_USER_DB", collectionName)
	if err != nil {
		return
	}

	defer os.Clearenv()
	mtestDb := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))

	mtestDb.Run("when_sending_a_valid_email_return_success", func(mt *mtest.T) {
		userEntity := entity.UserEntity{
			ID:       primitive.NewObjectID(),
			Email:    "test@test.com",
			Password: "test",
			Name:     "test",
			Age:      23,
		}
		mt.AddMockResponses(mtest.CreateCursorResponse(1,
			fmt.Sprintf("%s.%s", databaseName, collectionName),
			mtest.FirstBatch,
			convertEntityToBson(userEntity)))

		database := mt.Client.Database(databaseName)

		repo := NewUserRepository(database)
		userDomain, err := repo.FindUserByEmail(userEntity.Email)

		assert.Nil(t, err)
		assert.EqualValues(t, userDomain.GetID(), userEntity.ID.Hex())
		assert.EqualValues(t, userDomain.GetEmail(), userEntity.Email)
		assert.EqualValues(t, userDomain.GetName(), userEntity.Name)
		assert.EqualValues(t, userDomain.GetAge(), userEntity.Age)
		assert.EqualValues(t, userDomain.GetPassword(), userEntity.Password)
	})

	mtestDb.Run("returns_error_when_mongodb_returns_error", func(mt *mtest.T) {
		mt.AddMockResponses(bson.D{
			{Key: "ok", Value: 0},
		})

		database := mt.Client.Database(databaseName)

		repo := NewUserRepository(database)
		userDomain, err := repo.FindUserByEmail("test")

		assert.NotNil(t, err)
		assert.Nil(t, userDomain)
	})

	mtestDb.Run("returns_no_document_found", func(mt *mtest.T) {
		mt.AddMockResponses(mtest.CreateCursorResponse(0,
			fmt.Sprintf("%s.%s", databaseName, collectionName),
			mtest.FirstBatch))

		database := mt.Client.Database(databaseName)

		repo := NewUserRepository(database)
		userDomain, err := repo.FindUserByEmail("test")

		assert.NotNil(t, err)
		assert.Nil(t, userDomain)
	})
}

func TestUserRepository_FindUserByID(t *testing.T) {
	databaseName := "user_database_test"
	collectionName := "user_collection_test"
	err := os.Setenv("MONGODB_USER_DB", collectionName)
	if err != nil {
		return
	}

	defer os.Clearenv()
	mtestDb := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))

	mtestDb.Run("when_sending_a_valid_id_return_success", func(mt *mtest.T) {
		userEntity := entity.UserEntity{
			ID:       primitive.NewObjectID(),
			Email:    "test@test.com",
			Password: "test",
			Name:     "test",
			Age:      23,
		}
		mt.AddMockResponses(mtest.CreateCursorResponse(1,
			fmt.Sprintf("%s.%s", databaseName, collectionName),
			mtest.FirstBatch,
			convertEntityToBson(userEntity)))

		database := mt.Client.Database(databaseName)

		repo := NewUserRepository(database)
		userDomain, err := repo.FindUserByID(userEntity.ID.Hex())

		assert.Nil(t, err)
		assert.EqualValues(t, userDomain.GetID(), userEntity.ID.Hex())
		assert.EqualValues(t, userDomain.GetEmail(), userEntity.Email)
		assert.EqualValues(t, userDomain.GetName(), userEntity.Name)
		assert.EqualValues(t, userDomain.GetAge(), userEntity.Age)
		assert.EqualValues(t, userDomain.GetPassword(), userEntity.Password)
	})

	mtestDb.Run("returns_error_when_mongodb_returns_error", func(mt *mtest.T) {
		mt.AddMockResponses(bson.D{
			{Key: "ok", Value: 0},
		})

		database := mt.Client.Database(databaseName)

		repo := NewUserRepository(database)
		userDomain, err := repo.FindUserByID(primitive.NewObjectID().Hex())

		assert.NotNil(t, err)
		assert.Nil(t, userDomain)
	})

	mtestDb.Run("returns_no_document_found", func(mt *mtest.T) {
		mt.AddMockResponses(mtest.CreateCursorResponse(0,
			fmt.Sprintf("%s.%s", databaseName, collectionName),
			mtest.FirstBatch))

		database := mt.Client.Database(databaseName)

		repo := NewUserRepository(database)
		userDomain, err := repo.FindUserByID(primitive.NewObjectID().Hex())

		assert.NotNil(t, err)
		assert.Nil(t, userDomain)
	})
}

func TestUserRepository_FindUserByEmailAndPassword(t *testing.T) {
	databaseName := "user_database_test"
	collectionName := "user_collection_test"
	err := os.Setenv("MONGODB_USER_DB", collectionName)
	if err != nil {
		return
	}

	defer os.Clearenv()
	mtestDb := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))

	mtestDb.Run("when_sending_a_valid_email_and_password_return_success", func(mt *mtest.T) {
		userEntity := entity.UserEntity{
			ID:       primitive.NewObjectID(),
			Email:    "test@test.com",
			Password: "test",
			Name:     "test",
			Age:      23,
		}
		mt.AddMockResponses(mtest.CreateCursorResponse(1,
			fmt.Sprintf("%s.%s", databaseName, collectionName),
			mtest.FirstBatch,
			convertEntityToBson(userEntity)))

		database := mt.Client.Database(databaseName)

		repo := NewUserRepository(database)
		userDomain, err := repo.FindUserByEmailAndPassword(userEntity.Email, userEntity.Password)

		assert.Nil(t, err)
		assert.EqualValues(t, userDomain.GetID(), userEntity.ID.Hex())
		assert.EqualValues(t, userDomain.GetEmail(), userEntity.Email)
		assert.EqualValues(t, userDomain.GetName(), userEntity.Name)
		assert.EqualValues(t, userDomain.GetAge(), userEntity.Age)
		assert.EqualValues(t, userDomain.GetPassword(), userEntity.Password)
	})

	mtestDb.Run("returns_error_when_mongodb_returns_error", func(mt *mtest.T) {
		mt.AddMockResponses(bson.D{
			{Key: "ok", Value: 0},
		})

		database := mt.Client.Database(databaseName)

		repo := NewUserRepository(database)
		userDomain, err := repo.FindUserByEmailAndPassword("test", "sdfsdfsdfg")

		assert.NotNil(t, err)
		assert.Nil(t, userDomain)
	})

	mtestDb.Run("returns_no_document_found", func(mt *mtest.T) {
		mt.AddMockResponses(mtest.CreateCursorResponse(0,
			fmt.Sprintf("%s.%s", databaseName, collectionName),
			mtest.FirstBatch))

		database := mt.Client.Database(databaseName)

		repo := NewUserRepository(database)
		userDomain, err := repo.FindUserByEmailAndPassword("test", "asasasasa")

		assert.NotNil(t, err)
		assert.Nil(t, userDomain)
	})
}

func convertEntityToBson(userEntity entity.UserEntity) bson.D {
	return bson.D{
		{Key: "_id", Value: userEntity.ID},
		{Key: "email", Value: userEntity.Email},
		{Key: "password", Value: userEntity.Password},
		{Key: "name", Value: userEntity.Name},
		{Key: "age", Value: userEntity.Age},
	}
}
