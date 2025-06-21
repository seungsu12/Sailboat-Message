package main

import (
	"embed"
	"github.com/seungsu12/Sailboat-Message/router"
	"log"
	"net/http"
)

var staticFS embed.FS

func main() {

	fs := http.FileServer(http.FS(staticFS))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	r := router.New()
	log.Print("Sailboat Application Server Started")

	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatal(err)
	}
}
