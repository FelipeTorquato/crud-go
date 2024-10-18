package tests

import (
	"crud/src/controller/model/request"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"io"
	"math/rand"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
)

func TestLoginUser(t *testing.T) {

	t.Run("user_and_password_is_not_valid", func(t *testing.T) {
		recorderCreateUser := httptest.NewRecorder()
		ctxCreateUser := GetTestGinContext(recorderCreateUser)

		recorderLoginUser := httptest.NewRecorder()
		ctxLoginUser := GetTestGinContext(recorderLoginUser)

		email := fmt.Sprintf("%d@test.com", rand.Int())
		password := fmt.Sprintf("%d$#@$#@", rand.Int())

		userCreateRequest := request.UserRequest{
			Email:    email,
			Password: password,
			Name:     "test12",
			Age:      32,
		}

		bCreate, _ := json.Marshal(userCreateRequest)
		stringReaderCreate := io.NopCloser(strings.NewReader(string(bCreate)))
		MakeRequest(ctxCreateUser, []gin.Param{}, url.Values{}, "POST", stringReaderCreate)
		UserController.CreateUser(ctxCreateUser)

		userLoginRequest := request.UserLogin{
			Email:    "tstesdsd@tesst.com",
			Password: "ds2h323jd9ksk",
		}

		bLogin, _ := json.Marshal(userLoginRequest)
		stringReaderLogin := io.NopCloser(strings.NewReader(string(bLogin)))
		MakeRequest(ctxLoginUser, []gin.Param{}, url.Values{}, "POST", stringReaderLogin)
		UserController.LoginUser(ctxLoginUser)

		assert.EqualValues(t, http.StatusBadRequest, recorderLoginUser.Result().StatusCode)
	})

	t.Run("user_and_password_is_valid", func(t *testing.T) {
		recorderCreateUser := httptest.NewRecorder()
		ctxCreateUser := GetTestGinContext(recorderCreateUser)

		recorderLoginUser := httptest.NewRecorder()
		ctxLoginUser := GetTestGinContext(recorderLoginUser)

		email := fmt.Sprintf("%d@test.com", rand.Int())
		password := fmt.Sprintf("%d$#@$#@", rand.Int())

		userCreateRequest := request.UserRequest{
			Email:    email,
			Password: password,
			Name:     "test12",
			Age:      32,
		}

		bCreate, _ := json.Marshal(userCreateRequest)
		stringReaderCreate := io.NopCloser(strings.NewReader(string(bCreate)))
		MakeRequest(ctxCreateUser, []gin.Param{}, url.Values{}, "POST", stringReaderCreate)
		UserController.CreateUser(ctxCreateUser)

		userLoginRequest := request.UserLogin{
			Email:    email,
			Password: password,
		}

		bLogin, _ := json.Marshal(userLoginRequest)
		stringReaderLogin := io.NopCloser(strings.NewReader(string(bLogin)))
		MakeRequest(ctxLoginUser, []gin.Param{}, url.Values{}, "POST", stringReaderLogin)
		UserController.LoginUser(ctxLoginUser)

		assert.EqualValues(t, http.StatusOK, recorderLoginUser.Result().StatusCode)
		assert.NotEmpty(t, recorderLoginUser.Result().Header.Get("Authorization"))
	})
}
