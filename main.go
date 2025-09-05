package main

import "html/template"
import "net/http"
import "embed"

//go:embed views/**/*.html
var views embed.FS

func main() {

	mux := http.NewServeMux()

	var base = template.Must(template.New("root").ParseFS(views, "views/layouts/*.html", "views/partials/*.html"))

	index := template.Must(template.Must(base.Clone()).ParseFS(views, "views/pages/index.html"))

	about := template.Must(template.Must(base.Clone()).ParseFS(views, "views/pages/about.html"))

	products := template.Must(template.Must(base.Clone()).ParseFS(views, "views/pages/products.html"))

	mux.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		index.ExecuteTemplate(w, "base", nil)
	})

	mux.HandleFunc("GET /about", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		about.ExecuteTemplate(w, "base", nil)
	})

	mux.HandleFunc("GET /products", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		products.ExecuteTemplate(w, "base", nil)
	})

	http.ListenAndServe(":8080", mux)
}
