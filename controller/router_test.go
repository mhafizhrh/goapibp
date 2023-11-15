package controller_test

import (
	"goapibp/controller"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

func TestRouter(t *testing.T) {
	fiberApp := fiber.New()
	controller.Route(fiberApp)

	assert.Equal(t, "/auth/login", fiberApp.GetRoute("auth.login").Path)
}
