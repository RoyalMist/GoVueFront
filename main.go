//go:generate statik -src=./public/dist
package main

import (
	_ "GoVueFront/statik"
	"github.com/go-chi/chi"
	"github.com/unrolled/render"
	"log"
	"net/http"

	"github.com/rakyll/statik/fs"
)

type User struct {
	Name string
	Age  int8
}

func main() {
	sfs, err := fs.New()
	if err != nil {
		log.Fatal(err)
	}

	mux := chi.NewRouter()
	r := render.New()

	mux.Get("/*", func(writer http.ResponseWriter, request *http.Request) {
		fileServer := http.StripPrefix("/", http.FileServer(sfs))
		fileServer.ServeHTTP(writer, request)
	})

	mux.Get("/api", func(writer http.ResponseWriter, request *http.Request) {
		_ = r.JSON(writer, http.StatusOK, User{
			Name: "Joe",
			Age:  34,
		})
	})

	err = http.ListenAndServe("localhost:8080", mux)
	if err != nil {
		log.Fatal(err)
	}
}
