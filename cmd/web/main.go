package main

import (
	_ "embed"
	"flag"
	"fmt"
	"log"
	"net/http"
	"website/internal/router"

	"github.com/go-chi/chi/v5"
)

func main() {
	port := flag.String("port", "8080", "Port for SSH Server")
	host := flag.String("host", "localhost", "Host for SSH Server")
	flag.Parse()

	r := chi.NewRouter()

	r.Mount("/", router.PageRoutes())
	r.Mount("/ds", router.DataStarRoutes())

	log.Printf("Starting server on http://%s:%s", *host, *port)

	if err := http.ListenAndServe(fmt.Sprintf("%s:%s", *host, *port), r); err != nil {
		panic(err)
	}
}
