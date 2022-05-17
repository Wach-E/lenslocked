package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprint(w, "<h1>Welcome to my first awesome website!</h1>")
}

func contactHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprint(w, "<h1>Contact Page</h1><p>To get in touch, email me at <a href=\"mailto:immacurte1@gmail.com\">Wach Email</a>.</p>")
}

func faqHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprint(w, `<h1>FAQ Page</h1>
  <ul>
	<li>
	  <b>Is there a free version?</b>
	  Yes! We offer a free trial for 30 days on any paid plans.
	</li>
	<li>
	  <b>What are your support hours?</b>
	  We have support staff answering emails 24/7, though response
	  times may be a bit slower on weekends.
	</li>
	<li>
	  <b>How do I contact support?</b>
	  Email us - <a href="mailto:support@lenslocked.com">support@lenslocked.com</a>
	</li>
  </ul>
  `)
}

func getUser(w http.ResponseWriter, r *http.Request) {
	userID := chi.URLParam(r, "userID")
	// ctx := r.Context()
	// key := ctx.Value("key").(string)

	// respond to the client
	fmt.Fprintf(w, "hi %v", userID)
	// fmt.Fprintf(w, "hi %v, %v", userID, key)
	// w.Write([]byte(fmt.Sprintf("hi %v, %v", userID, key)))
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
