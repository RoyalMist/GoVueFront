//go:generate statik -src=./public/dist
package main

import (
	_ "GoVueFront/statik"
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
	e.GET("/api", func(c echo.Context) error {
		return c.JSON(http.StatusOK, User{
			Name: "John Doe",
			Age:  35,
		})
	})

	e.Logger.Fatal(e.Start("localhost:8080"))
}
