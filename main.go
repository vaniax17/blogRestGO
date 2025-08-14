package blogRest

import (
	"blogRest/src/db"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	if err := e.Start(":8080"); err != nil {
		e.Logger.Fatal(err)
		db.Close()
	}

}
