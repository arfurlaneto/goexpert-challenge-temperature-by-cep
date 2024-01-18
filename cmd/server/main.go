package main

import (
	"net/http"

	"github.com/arfurlaneto/goexpert-challenge-temperature-by-cep/internal/handlers"
	"github.com/arfurlaneto/goexpert-challenge-temperature-by-cep/internal/services"
	"github.com/arfurlaneto/goexpert-challenge-temperature-by-cep/internal/usecases"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load(".env")

	r := chi.NewRouter()
	r.Use(middleware.Recoverer)

	getTemperatureByCepHandler := handlers.NewGetTemperatureByCepHandler(
		usecases.NewGetTemperatureByCepUseCase(
			services.NewViaCepService(),
			services.NewWeatherApiService(),
		),
	)

	r.Get("/healthz", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	})

	r.Get("/", getTemperatureByCepHandler.Handle)

	err := http.ListenAndServe(":8080", r)
	if err != nil {
		panic(err)
	}
}
