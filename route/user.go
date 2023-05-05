package route

import (
	"test-echo/config"
	"test-echo/handler"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func User(e *echo.Echo, userHandler *handler.UserHandler) {

	r := e.Group("/users")
	r.POST("/", userHandler.CreateUser)
	r.GET("/", userHandler.GetAllUser)
	r.POST("/login", handler.Login)

	config := middleware.JWTConfig{
		SigningKey: []byte(config.Config("SECRET")),
	}

	r.Use(middleware.JWTWithConfig(config))
	r.GET("/:id", userHandler.GetUserByID)
	r.PUT("/:id", userHandler.UpdateUser)
	r.DELETE("/:id", userHandler.DeleteUser)

}
