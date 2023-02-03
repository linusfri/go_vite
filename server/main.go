package main

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
)

func main() {
	r := chi.NewRouter()
	db := NewDatabase()

	defer db.Client.Close()

	r.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{"http://localhost:5173"},
	}))

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello"))
	})

	r.Route("/person", func(r chi.Router) {
		r.Get("/", getPerson)
	})

	http.ListenAndServe(":8080", r)
}

func getPerson(w http.ResponseWriter, r *http.Request) {
	type Person struct {
		Name string
	}

	res, _ := json.Marshal(Person{Name: "pelle"})
	w.Write([]byte(res))
}
