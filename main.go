package main

import (
	"github.com/seungsu12/Sailboat-Message/router"
	"log"
	"net/http"
)

func main() {

	r := router.New()
	log.Print("Sailboat Application Server Started")

	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatal(err)
	}
}
