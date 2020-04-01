package controller

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

//Home é a pagina inicial da minha aplicacao
func Home(c echo.Context) error {
	return c.String(http.StatusOK, "Hello World!!")
}
