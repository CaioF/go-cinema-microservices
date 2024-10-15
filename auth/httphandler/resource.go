package httphandler

import "github.com/caiof/go-cinema-microservices/auth/model"

type (
	// For Post/Put - /users
	UserResource struct {
		Data model.User `json:"data"`
	}
)
