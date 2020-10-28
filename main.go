package main

import (
	"log"
	"net/http"
	"os"

	"github.com/mbogunovic/working/handlers"
)

func main() {
	l := log.New(os.Stdout, "product-api", log.LstdFlags)
	hh := handlers.NewHello(l)

	http.HandleFunc("/")

	http.HandleFunc("/goodbye", func(rw http.ResponseWriter, r *http.Request) {
		log.Println("Goodbye world!")
	})

	http.ListenAndServe(":9090", nil)
}
