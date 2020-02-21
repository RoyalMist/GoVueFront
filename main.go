//go:generate statik -src=./public/dist
package main

import (
	_ "GoVueFront/statik"
	"errors"
	"github.com/labstack/echo/v4"
	"github.com/rakyll/statik/fs"
	"log"
	"net/http"
)

type User struct {
	Name string
	Age  int8
}

func main() {
	e := echo.New()
	sfs, err := fs.New()
	if err != nil {
		log.Fatal(err)
	}

	e.GET("/*", echo.WrapHandler(http.StripPrefix("/", http.FileServer(sfs))))
	e.GET("/api", getHandler)
	e.POST("/api", postHandler)

	e.Logger.Fatal(e.Start("localhost:8080"))
}

func postHandler(c echo.Context) error {
	u := new(User)
	if err := c.Bind(u); err != nil {
		return errors.New("boom")
	}

	return c.JSON(http.StatusOK, u)
}

func getHandler(c echo.Context) error {
	return c.JSON(http.StatusOK, User{
		Name: "John Doe",
		Age:  35,
	})
}
