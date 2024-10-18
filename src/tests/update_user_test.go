package tests

import (
	"context"
	"crud/src/controller/model/request"
	"crud/src/model/repository/entity"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
)

func TestUpdateUser(t *testing.T) {
	recorder := httptest.NewRecorder()
	ctx := GetTestGinContext(recorder)
	id := primitive.NewObjectID()

	_, err := Database.
		Collection("test_user").
		InsertOne(context.Background(), bson.M{"_id": id, "name": "old_name", "age": 18, "email": "test@test.com"})

	if err != nil {
		t.Fatal(err)
		return
	}

	param := []gin.Param{
		{
			Key:   "userId",
			Value: id.Hex(),
		},
	}

	userRequest := request.UserUpdateRequest{
		Name: "new_name",
		Age:  int8(78),
	}

	b, _ := json.Marshal(userRequest)
	stringReader := io.NopCloser(strings.NewReader(string(b)))
	MakeRequest(ctx, param, url.Values{}, "PUT", stringReader)
	UserController.UpdateUser(ctx)

	assert.EqualValues(t, http.StatusOK, recorder.Result().StatusCode)

	userEntity := entity.UserEntity{}

	filter := bson.D{{Key: "_id", Value: id}}
	_ = Database.
		Collection("test_user").
		FindOne(context.Background(), filter).Decode(&userEntity)

	assert.EqualValues(t, userRequest.Name, userEntity.Name)
	assert.EqualValues(t, userRequest.Age, userEntity.Age)
}
