package lib

import (
	"log"

	"upper.io/db.v3"
	"upper.io/db.v3/mysql"
)

var settings = mysql.ConnectionURL{
	Host:     "localhost",
	User:     "api",
	Password: "api",
	Database: "api-go",
}

// Sess que é uma variavel que faz conexao com banco de dados.
var Sess db.Database

func init() {
	var err error

	Sess, err = mysql.Open(settings)
	if err != nil {
		log.Fatal(err.Error())
	}
}
