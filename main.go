//go:generate statik -src=./public/dist
package main

import (
	_ "GoVueFront/statik"
	"log"
	"net/http"

	"github.com/rakyll/statik/fs"
)

func main() {
	sfs, err := fs.New()
	if err != nil {
		log.Fatal(err)
	}

	http.Handle("/", http.StripPrefix("/", http.FileServer(sfs)))
	_ = http.ListenAndServe("localhost:8080", nil)
}
