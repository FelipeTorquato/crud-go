package controller

import (
	"crud/src/configuration/rest_err"
	"crud/src/controller/model/request"
	"crud/src/model"
	"crud/src/tests/mocks"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.uber.org/mock/gomock"
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
)

func TestUserControllerInterface_UpdateUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	service := mocks.NewMockUserDomainService(ctrl)
	controller := NewUserControllerInterface(service)

	t.Run("validation_body_got_error", func(t *testing.T) {
		recorder := httptest.NewRecorder()
		context := GetTestGinContext(recorder)

		userRequest := request.UserUpdateRequest{
			Name: "a",
			Age:  -1,
		}

		b, _ := json.Marshal(userRequest)
		stringReader := io.NopCloser(strings.NewReader(string(b)))

		MakeRequest(context, []gin.Param{}, url.Values{}, "PUT", stringReader)
		controller.UpdateUser(context)

		assert.EqualValues(t, http.StatusBadRequest, recorder.Code)
	})

	t.Run("validation_userId_got_error", func(t *testing.T) {
		recorder := httptest.NewRecorder()
		context := GetTestGinContext(recorder)

		userRequest := request.UserUpdateRequest{
			Name: "test123",
			Age:  23,
		}

		param := []gin.Param{
			{
				Key:   "userId",
				Value: "userId",
			},
		}

		b, _ := json.Marshal(userRequest)
		stringReader := io.NopCloser(strings.NewReader(string(b)))

		MakeRequest(context, param, url.Values{}, "PUT", stringReader)
		controller.UpdateUser(context)

		assert.EqualValues(t, http.StatusBadRequest, recorder.Code)
	})

	t.Run("object_is_valid_service_returns_error", func(t *testing.T) {
		recorder := httptest.NewRecorder()
		context := GetTestGinContext(recorder)
		id := primitive.NewObjectID().Hex()

		userRequest := request.UserUpdateRequest{
			Name: "test1234",
			Age:  12,
		}

		param := []gin.Param{
			{
				Key:   "userId",
				Value: id,
			},
		}

		domain := model.NewUserUpdateDomain(
			userRequest.Name,
			userRequest.Age,
		)

		b, _ := json.Marshal(userRequest)
		stringReader := io.NopCloser(strings.NewReader(string(b)))

		service.EXPECT().UpdateUser(id, domain).Return(
			rest_err.NewInternalServerError("error test"))

		MakeRequest(context, param, url.Values{}, "POST", stringReader)
		controller.UpdateUser(context)

		assert.EqualValues(t, http.StatusInternalServerError, recorder.Code)
	})

	t.Run("object_is_valid_service_returns_success", func(t *testing.T) {
		recorder := httptest.NewRecorder()
		context := GetTestGinContext(recorder)
		id := primitive.NewObjectID().Hex()

		userRequest := request.UserUpdateRequest{
			Name: "test1234",
			Age:  12,
		}

		param := []gin.Param{
			{
				Key:   "userId",
				Value: id,
			},
		}

		domain := model.NewUserUpdateDomain(
			userRequest.Name,
			userRequest.Age,
		)

		b, _ := json.Marshal(userRequest)
		stringReader := io.NopCloser(strings.NewReader(string(b)))

		service.EXPECT().UpdateUser(id, domain).Return(
			nil)

		MakeRequest(context, param, url.Values{}, "POST", stringReader)
		controller.UpdateUser(context)

		assert.EqualValues(t, http.StatusOK, recorder.Code)
	})
}
