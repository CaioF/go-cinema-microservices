package router

import (
	"github.com/go-chi/chi"
	"github.com/caiof/go-cinema-microservices/auth/httphandler"
	"github.com/caiof/go-cinema-microservices/auth/storage/mongodb"
	"github.com/caiof/go-cinema-microservices/auth/usecase/auth"
)

func InitRouter(r *chi.Mux, s *mongodb.Storage) *chi.Mux {
	authHandler := initDependencies(s)
	r.Post("/login", authHandler.Login())
	r.Post("/authenticate", authHandler.Authentication())
	r.Get("/healthy", authHandler.GetReadiness())
	r.Get("/healthz", authHandler.GetLiveness())

	return r
}

func initDependencies(storage *mongodb.Storage) httphandler.AuthHandler {
	authenticator := auth.NewAuthentication(storage)
	authHandler := httphandler.NewAuthHandler(authenticator, storage)
	return *authHandler
}
