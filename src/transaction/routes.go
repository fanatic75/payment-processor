package transaction

import "github.com/go-chi/chi/v5"

func SetupRoutes() *chi.Mux {
	r := chi.NewRouter()
	r.Post("/", CreateTransaction)
	return r
}
