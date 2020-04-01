package controller

import (
	"go-api/model"
	"net/http"

	"github.com/labstack/echo/v4"
)

//Home é a pagina inicial da minha aplicacao
func Home(c echo.Context) error {
	return c.String(http.StatusOK, "Hello World!!")
}

func InserirUsuario(c echo.Context) error {
	nome := c.FormValue("nome")
	email := c.FormValue("email")

	var usuario model.Usuarios
	usuario.Nome = nome
	usuario.Email = email

	if nome != "" && email != "" {

		if _, err := model.UsuarioModel.Insert(usuario); err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{
				"message": "Não foi possivel inserir o registro, verifique os campos envidos.",
			})
		}

		return c.JSON(http.StatusCreated, map[string]string{
			"message": "Registro inserido com sucesso.",
		})

	}

	return c.JSON(http.StatusBadRequest, map[string]string{
		"message": "Os campos devem ser preenchidos.",
	})
}
