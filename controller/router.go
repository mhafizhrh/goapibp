package controller

import (
	"goapibp/handler"
	"goapibp/helper"

	"github.com/gofiber/fiber/v2"
)

func Route(fiberRouter fiber.Router) {
	// Load Helper
	jwtHelper := helper.NewJWT()

	// Load Handler
	authHandler := handler.NewAuth(jwtHelper)

	// // Load Middleware
	// limiterMiddleware := middleware.NewLimiter()

	NewAuth(authHandler).Router(fiberRouter)
}
