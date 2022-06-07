package main

import (
	"context"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func executeTemplate(w http.ResponseWriter, filepath string) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	// tpl, err := template.ParseFiles("templates\\home.gohtml")
	tpl, err := template.ParseFiles(filepath)
	if err != nil {
		log.Printf("parsinng template: %v", err)
		http.Error(w, "There was an error parsing the template.",
			http.StatusInternalServerError)
		return
	}

	err = tpl.Execute(w, nil)
	if err != nil {
		log.Printf("executing template: %v", err)
		http.Error(w, "There was an error executing the template.",
			http.StatusInternalServerError)
		return
	}
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	tplPath := filepath.Join("templates", "home.gohtml")
	executeTemplate(w, tplPath)
}

func contactHandler(w http.ResponseWriter, r *http.Request) {
	tplPath := filepath.Join("templates", "contact.gohtml")
	executeTemplate(w, tplPath)
}

func faqHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	tplPath := filepath.Join("templates", "faq.gohtml")
	executeTemplate(w, tplPath)
}

func getUser(w http.ResponseWriter, r *http.Request) {
	userID := chi.URLParam(r, "userID")

	ctx := context.WithValue(r.Context(), userID, 123)
	// fmt.Println(ctx.Value(key))
	key := ctx.Value(userID)

	// respond to the client
	// fmt.Fprintf(w, "hi %v", userID)
	fmt.Fprintf(w, "hi %v, %v", userID, key)
}

func main() {
	r := chi.NewRouter()

	r.Get("/", homeHandler)
	r.Route("/contact", func(r chi.Router) {
		r.Use(middleware.Logger)
		r.Get("/", contactHandler)
	})

	r.Get("/faq", faqHandler)

	r.Route("/users", func(r chi.Router) {
		// Subrouters:
		r.Route("/{userID}", func(r chi.Router) {
			r.Use(middleware.Logger)
			r.Get("/", getUser) // GET /users/123
		})
	})

	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Page not found", http.StatusNotFound)
	})

	fmt.Println("Starting the server on :3000...")
	http.ListenAndServe(":3000", r)
}
