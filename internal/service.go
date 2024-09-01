package internal

import (
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/spf13/viper"
)

type Service struct {
}

func (*Service) Run() {
	r := createRouter()
	port := viper.GetInt32("PORT")
	server := http.Server{
		Addr:    fmt.Sprintf(":%d", port),
		Handler: r,
	}
	log.Printf("Listening on port :%d\n", port)
	if err := server.ListenAndServe(); err != nil {
		log.Printf("Failed to start http server: %v", err)
	}
}

func createRouter() *chi.Mux {
	r := chi.NewRouter()
	r.Get("/persons", func(w http.ResponseWriter, r *http.Request) {
		persons, err := GetPersons()
		if err != nil {
			fmt.Fprintf(w, "Failed to fetch persons from db: %v", err)
		}
		for _, person := range persons {
			fmt.Fprintf(w, "Name: %s, Id: %d\n", person.Name, person.Id)
		}
	})
	return r
}
