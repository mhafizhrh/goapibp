package handler

import (
	"goapibp/helper"
	"goapibp/model"
	"goapibp/response"
)

type Auth interface {
	Login(request *model.AuthLoginRequest) response.CustomResponse
}

type auth struct {
	jwtHelper helper.JWT
}

func NewAuth(jwtHelper helper.JWT) Auth {
	return &auth{
		jwtHelper: jwtHelper,
	}
}

func (handler auth) Login(request *model.AuthLoginRequest) response.CustomResponse {
	token := handler.jwtHelper.CreateToken()
	return response.NewResponse(200, response.AuthLoginResponse{
		Token: token,
	})
}
