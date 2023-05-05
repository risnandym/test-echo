package route

import (
	"test-echo/config"
	"test-echo/handler"
	"test-echo/repository"
	"test-echo/service"

	"github.com/labstack/echo/v4"

	echoSwagger "github.com/swaggo/echo-swagger" // echo-swagger middleware
)

func SetupRouter(e *echo.Echo) {

	userRepository := repository.NewUserRepository(config.DB)
	userService := service.NewUserService(userRepository)
	userHandler := handler.NewUserHandler(userService)
	User(e, userHandler)

	e.GET("/swagger/*any", echoSwagger.WrapHandler)
}
