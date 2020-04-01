package controller

import (
	"go-api/model"
	"net/http"
	"strconv"

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

func ListarUsuarios(c echo.Context) error {
	var usuarios []model.Usuarios

	if err := model.UsuarioModel.Find().All(&usuarios); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": "Erro ao listar usuarios",
		})
	}

	return c.JSON(http.StatusOK, usuarios)
}

func DeletarUsuario(c echo.Context) error {
	usuarioID, _ := strconv.Atoi(c.Param("id"))

	resultado := model.UsuarioModel.Find("id=?", usuarioID)
	if count, _ := resultado.Count(); count > 0 {
		if err := resultado.Delete(); err != nil {
			return c.JSON(http.StatusNotFound, map[string]string{
				"message": "Falha ao excluir Usuário",
			})
		}
		return c.JSON(http.StatusAccepted, map[string]string{
			"message": "Usuário foi excluido com sucesso.",
		})
	}

	return c.JSON(http.StatusFound, map[string]string{
		"message": "Não foi possivel encontrar usuário para excluir",
	})
}

func CarregarUsuario(c echo.Context) error {
	usuarioID, _ := strconv.Atoi(c.Param("id"))
	var usuario model.Usuarios

	if err := model.UsuarioModel.Find("id=?", usuarioID).One(&usuario); err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{
			"message": "Registro nao encontrado",
		})
	}
	return c.JSON(http.StatusOK, usuario)
}

func AlterarUsuario(c echo.Context) error {
	usuarioID, _ := strconv.Atoi(c.Param("id"))
	nome := c.FormValue("nome")
	email := c.FormValue("email")

	var usuario = model.Usuarios{
		ID:    usuarioID,
		Nome:  nome,
		Email: email,
	}

	resultado := model.UsuarioModel.Find("id=?", usuarioID)
	if count, _ := resultado.Count(); count > 0 {
		if err := resultado.Update(&usuario); err != nil {
			return c.JSON(http.StatusNotFound, map[string]string{
				"message": "Falha ao atualizar Usuário",
			})
		}
		return c.JSON(http.StatusAccepted, map[string]string{
			"message": "Usuário atualizado com sucesso.",
		})
	}

	return c.JSON(http.StatusFound, map[string]string{
		"message": "Não foi possivel encontrar usuário para excluir",
	})

	// if err := model.UsuarioModel.Find("id=?", usuarioID).One(&usuario); err != nil {
	// 	return c.JSON(http.StatusNotFound, map[string]string{
	// 		"message": "Registro nao encontrado",
	// 	})
	// }

}
