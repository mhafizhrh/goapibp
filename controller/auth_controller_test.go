package controller_test

import (
	"bytes"
	"encoding/json"
	"goapibp/controller"
	"goapibp/handler"
	"goapibp/helper"
	helperMocks "goapibp/helper/mocks"
	"goapibp/model"
	"goapibp/response"
	"io"
	"log"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

func TestAuthLogin(t *testing.T) {
	t.Run("OK", func(t *testing.T) {
		// Initial model
		request := &model.AuthLoginRequest{
			Username: "veryhardusername",
			Password: "12345678",
		}

		// Load fiber instance
		fiberApp := fiber.New()

		// Load mock
		// authHandlerMock := new(handlerMocks.Auth)
		jwtHelperMock := new(helperMocks.JWT)

		// Mock Methods
		// authHandlerMock.On("Login", request).Return(response.NewResponse(200, "hello oworld"))
		token := "afkasdasd"
		jwtHelperMock.On("CreateToken").Return(token)

		authHandler := handler.NewAuth(jwtHelperMock)

		// Load controller instance
		authController := controller.NewAuth(authHandler)
		authController.Router(fiberApp)

		// Test HTTP Request
		reqBody := new(bytes.Buffer)
		err := json.NewEncoder(reqBody).Encode(request)
		httpReq := httptest.NewRequest("POST", "/auth/login", reqBody)
		httpReq.Header.Add("Content-Type", "application/json")
		resp, err := fiberApp.Test(httpReq)
		if err != nil {
			log.Panic(err)
		}
		respBody := new(response.AuthLoginResponse)
		err = json.NewDecoder(resp.Body).Decode(respBody)
		if err != nil {
			log.Panic(err)
		}
		assert.Equal(t, 200, resp.StatusCode)
		expectResponse := &response.AuthLoginResponse{
			Token: token,
		}
		assert.Equal(t, expectResponse, respBody)
	})

	t.Run("UnprocessableEntity", func(t *testing.T) {
		// Initial model
		request := &model.AuthLoginRequest{
			Username: "veryhardusername",
			Password: "12345678",
		}

		// Load fiber instance
		fiberApp := fiber.New()

		// Load mock
		// authHandlerMock := new(handlerMocks.Auth)
		jwtHelperMock := new(helperMocks.JWT)

		// Mock Methods
		// authHandlerMock.On("Login", request).Return(response.NewResponse(200, "hello oworld"))
		token := "afkasdasd"
		jwtHelperMock.On("CreateToken").Return(token)

		authHandler := handler.NewAuth(jwtHelperMock)

		// Load controller instance
		authController := controller.NewAuth(authHandler)
		authController.Router(fiberApp)

		// Test HTTP Request
		reqBody := new(bytes.Buffer)
		err := json.NewEncoder(reqBody).Encode(request)
		httpReq := httptest.NewRequest("POST", "/auth/login", reqBody)
		resp, err := fiberApp.Test(httpReq)
		if err != nil {
			log.Panic(err)
		}
		respBodyByte, err := io.ReadAll(resp.Body)
		if err != nil {
			log.Panic(err)
		}
		assert.Equal(t, "Unprocessable Entity", string(respBodyByte))
	})
}

func TestAuthRouter(t *testing.T) {
	jwtHelper := helper.NewJWT()
	authHandler := handler.NewAuth(jwtHelper)
	authController := controller.NewAuth(authHandler)

	fiberApp := fiber.New()
	authController.Router(fiberApp)
}
