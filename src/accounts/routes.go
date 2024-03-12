package accounts

import (
	"github.com/go-chi/chi/v5"
)

func SetupRoutes() *chi.Mux {
	r := chi.NewRouter()
	r.Post("/", CreateAccount)
	r.Get("/{id}", GetAccount)
	return r
}
