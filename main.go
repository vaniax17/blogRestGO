package main

import (
	"blogRest/src/api"
	"blogRest/src/db"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {

	err := godotenv.Load(".env")
	if err != nil {
		panic("Error loading .env file")
	}
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	db.Init()
	api.RoutesInit(e)

	if err := e.Start(":8080"); err != nil {
		e.Logger.Fatal(err)
		db.Close()
	}

}
