package main

import (
	"go-api/router"

	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := router.App

	e.Use(middleware.Logger())

	e.Logger.Fatal(e.Start(":3333"))
}
