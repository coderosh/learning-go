package main

import (
	"gshort/db"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	db.Init()

	var address = ":8080"

	app := echo.New()
	app.Static("/public/", "public")

	app.Pre(middleware.AddTrailingSlash())

	app.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "${host} ${method} ${uri}\n",
	}))

	app.Use(middleware.Recover())

	Routes(app)

	app.HTTPErrorHandler = customHTTPErrorHandler

	app.Logger.Fatal(app.Start(address))

}

func customHTTPErrorHandler(err error, c echo.Context) {
	code := http.StatusInternalServerError
	if he, ok := err.(*echo.HTTPError); ok {
		code = he.Code
	}

	log.Println(code, err)

	// render file instead
	c.JSON(code, map[string]string{
		"message": err.Error(),
	})
}
