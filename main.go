package main

import (
	"go-api/router"

	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo/v4/middleware"
	_ "upper.io/db.v3/mysql"
)

func main() {
	e := router.App

	e.Use(middleware.Logger())

	e.Logger.Fatal(e.Start(":3333"))
}
