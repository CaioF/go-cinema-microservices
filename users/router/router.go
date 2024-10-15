package router

import (
	"github.com/go-chi/chi"
	"github.com/caiof/go-cinema-microservices/users/httphandler"
	"github.com/caiof/go-cinema-microservices/users/storage"
)

func InitRouter(r *chi.Mux, s storage.Storage) *chi.Mux {

	r.Route("/users", func(r chi.Router) {
		r.Get("/", httphandler.GetUsers(s))
		r.Post("/", httphandler.CreateUser(s))
		r.Delete("/{id}", httphandler.DeleteUser(s))
	})

	r.Get("/healthy", httphandler.GetReadiness(s))
	r.Get("/healthz", httphandler.GetLiveness())

	return r
}
