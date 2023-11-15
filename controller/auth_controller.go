package controller

import (
	"goapibp/handler"
	"goapibp/middleware"
	"goapibp/model"

	"github.com/gofiber/fiber/v2"
)

type AuthImpl interface {
	Router(fiberRouter fiber.Router)
}

type auth struct {
	// limiterMiddleware middleware.Auth
	authHandler handler.Auth
}

func NewAuth(authHandler handler.Auth) AuthImpl {
	return &auth{
		// limiterMiddleware: limiterMiddleware,
		authHandler: authHandler,
	}
}
func (controller auth) Router(fiberRouter fiber.Router) {
	auth := fiberRouter.Group("/auth")
	auth.Post("/login", middleware.Limiter, controller.Login).Name("auth.login")
}

func (controller auth) Login(fiberCtx *fiber.Ctx) error {
	request := new(model.AuthLoginRequest)
	err := fiberCtx.BodyParser(request)
	if err != nil {
		return err
	}
	response := controller.authHandler.Login(request)

	return response.ToFiberResponse(fiberCtx)
}
