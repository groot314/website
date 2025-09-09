package router

import (
	"website/internal/content"

	"github.com/go-chi/chi/v5"
)

func DataStarRoutes() chi.Router {
	r := chi.NewRouter()

	r.Get("/console-hello", content.HelloConsole)
	r.Get("/about", content.GetAbout)
	r.Get("/home", content.GetHomeContent)

	return r
}

func PageRoutes() chi.Router {
	r := chi.NewRouter()

	r.Get("/", content.GetHome)

	return r
}
