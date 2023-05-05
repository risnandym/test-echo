package middleware

import (
	"test-echo/config"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func IsLogIn(e *echo.Echo) middleware.JWTConfig {
	config := middleware.JWTConfig{
		SigningKey: []byte(config.Getenv("API_SECRET", "rahasiasekali")),
	}
	return config
}

// func jwtError(c echo.Context, err error) error {
// 	if err.Error() == "Missing or malformed JWT" {
// 		return c.JSON(http.StatusBadRequest, entity.ErrorResponse{
// 			Code:    http.StatusBadRequest,
// 			Message: "Missing or malformed JWT",
// 			Error:   err.Error(),
// 		})
// 	}
// 	return echo.NewHTTPError(http.StatusUnauthorized, "Please provide valid credentials")
// }
