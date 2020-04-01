package router

import (
	"go-api/controller"

	"github.com/labstack/echo/v4"
)

// App é uma instancia de Echo
var App *echo.Echo

func init() {
	App = echo.New()

	App.GET("/", controller.Home)
}
