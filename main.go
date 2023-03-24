package main

import (
	"fmt"
	"log"
	"os"
	"net/http"
	"github.com/go-chi/chi/v5"
	"go-server/api"
)

func main() {

	fmt.Println("Running main")

	r := chi.NewRouter()

	fmt.Println("Enabling routes")

	r.Route("/api/blog", func(r chi.Router) {
		r.Get("/getLatestPosts", api.FetchPosts)
		r.Route("/posts", func(r chi.Router) {
			r.Get("/", api.FetchPosts)
		});
	})

	FileServer(r)

	fmt.Println("Starting server on port 8080")

	if err:=http.ListenAndServe(":8080", r); err != nil {
		log.Fatal(err)	
		fmt.Printf("An error occurred")
	}
}

func FileServer(router *chi.Mux) {
	root := "./public"
	fs := http.FileServer(http.Dir(root))

	router.Get("/*", func(w http.ResponseWriter, r *http.Request) {
		if _, err := os.Stat(root + r.RequestURI); os.IsNotExist(err) {
			http.StripPrefix(r.RequestURI, fs).ServeHTTP(w, r)
		} else {
			fs.ServeHTTP(w, r)
		}
	})
}
