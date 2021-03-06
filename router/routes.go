package router

import (
	"go-api/controller"

	"github.com/labstack/echo/v4"
)

// App é uma instancia de Echo
var App *echo.Echo

func init() {
	App = echo.New()

	// Pagina inicial da Aplicacao.
	App.GET("/", controller.Home)

	api := App.Group("/api/v1/usuarios")
	api.POST("/", controller.InserirUsuario)
	api.GET("/", controller.ListarUsuarios)
	api.GET("/:id", controller.CarregarUsuario)
	api.PUT("/:id", controller.AlterarUsuario)
	api.DELETE("/:id", controller.DeletarUsuario)

}
