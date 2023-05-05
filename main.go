package main

import (
	"test-echo/config"
	"test-echo/docs"
	"test-echo/route"

	"github.com/labstack/echo/v4"
)

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @termsOfService http://swagger.io/terms/

func main() {

	// programmatically set swagger info
	docs.SwaggerInfo.Title = "Test Case"
	docs.SwaggerInfo.Description = "This is a Test Case."
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = config.Getenv("SWAGGER_HOST", "localhost:8080")
	docs.SwaggerInfo.Schemes = []string{"http", "https"}

	// database connection
	config.ConnectDataBase()

	// router
	r := echo.New()
	route.SetupRouter(r)

	r.Logger.Fatal(r.Start(":8080"))
}
