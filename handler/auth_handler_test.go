package handler_test

import (
	"goapibp/handler"
	helperMocks "goapibp/helper/mocks"
	"goapibp/model"
	"goapibp/response"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAuthLogin(t *testing.T) {
	t.Run("OK", func(t *testing.T) {
		request := &model.AuthLoginRequest{
			Username: "veryhardusername",
			Password: "12345678",
		}
		jwtHelper := new(helperMocks.JWT)
		token := "jkasjf9w0ejfg203ifn"
		jwtHelper.On("CreateToken").Return(token)
		authHandler := handler.NewAuth(jwtHelper)

		handlerResp := authHandler.Login(request)

		assert.Equal(t, response.NewResponse(200, response.AuthLoginResponse{Token: token}), handlerResp)
	})
}
