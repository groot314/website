package main

import (
	_ "embed"
	"fmt"
	"log"
	"net/http"
	"website/internal/router"

	"github.com/go-chi/chi/v5"
)

const port = 8080

func main() {
	r := chi.NewRouter()

	r.Mount("/", router.PageRoutes())
	r.Mount("/ds", router.DataStarRoutes())

	log.Printf("Starting server on http://localhost:%d", port)

	if err := http.ListenAndServe(fmt.Sprintf(":%d", port), r); err != nil {
		panic(err)
	}
}
