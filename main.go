package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/kidsan/contacts-api/contact"
	"github.com/kidsan/contacts-api/logger"
)

func main() {
	logger := logger.NewLogger()
	r := chi.NewRouter()
	contactRoutes := contact.BuildRouter()
	r.Mount("/contacts", contactRoutes)
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("."))
	})

	logger.Info("Application listening on port 3000")
	http.ListenAndServe(":3000", r)
}
