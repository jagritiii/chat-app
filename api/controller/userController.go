package controller

import (
	"chat-app/api/handler"
	"github.com/labstack/echo/v4"
)

func SetupUserRoutes(e *echo.Echo) {
	userRoutes := e.Group("/users")
	{
		userRoutes.GET("", handler.GetUsers)
		userRoutes.POST("/signup", handler.SignUp)
		userRoutes.POST("/login", handler.Login)
		userRoutes.GET("/users", handler.GetUsers)
		userRoutes.GET("/user/:user_id", handler.GetUser)

	}
}
