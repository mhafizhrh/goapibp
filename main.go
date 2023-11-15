package main

import (
	"goapibp/controller"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func init() {
	godotenv.Load(".env")
}

func main() {
	fiberApp := fiber.New()
	appVersion := os.Getenv("MAJOR_APP_VERSION")
	controller.Route(fiberApp.Group("/" + appVersion))

	fiberApp.Listen(":7000")
}
