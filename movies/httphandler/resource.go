package httphandler

import "github.com/caiof/go-cinema-microservices/movies/model"

type (
	// For Get - /movies
	MoviesResource struct {
		Data []model.Movie `json:"data"`
	}
	// For Post/Put - /movies
	MovieResource struct {
		Data model.Movie `json:"data"`
	}
)
