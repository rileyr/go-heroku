package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/rakyll/statik/fs"
	"github.com/rileyr/go-heroku/home"
	_ "github.com/rileyr/go-heroku/statik"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	virtualFS, err := fs.New()
	if err != nil {
		log.Fatal(err)
	}

	hh := home.NewHandler(virtualFS)

	http.Handle("/", http.HandlerFunc(hh.Home))
	http.ListenAndServe(fmt.Sprintf(":%s", port), nil)
}
