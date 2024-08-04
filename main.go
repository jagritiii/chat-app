package main

import (
	"chat-app/api/controller"
	middlewares "chat-app/pkg/middleware"
	"os"

	"github.com/labstack/echo/v4"
)

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		port = "8000"
	}

	e := echo.New()
	// e.Use(middleware.Logger())

	controller.Setupe(e)
	authRoutes := e.Group("")
	authRoutes.Use(middlewares.AuthenticationMiddleware)

	e.Logger.Fatal(e.Start(":" + port))
}
