package server

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func (s *Server) RegisterRoutes() http.Handler {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	fileServer := http.FileServer(http.Dir("./src"))
	r.Get("/api", ApiRequestHandler)
	r.Get("/search",SearchRequestHandler)
    r.Get("/highlighted-verses", HighlightedVerseHandler)
	r.Handle("/*", fileServer)

	return r
}
