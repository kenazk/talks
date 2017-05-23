package main

import (
	"html/template"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

func main() {
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	http.HandleFunc("/", serveTemplate)

	log.Println("Listening...")
	http.ListenAndServe(":80", Log(http.DefaultServeMux))
}

func Log(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s %s", r.RemoteAddr, r.Method, r.URL)
		handler.ServeHTTP(w, r)
	})
}

func serveTemplate(w http.ResponseWriter, r *http.Request) {
	lp := filepath.Join("templates", "layout.html")
	fp := filepath.Join("templates", filepath.Clean(r.URL.Path))

	if r.URL.Path == "/" {
		fp = "templates/index.html"
	}

	tmpl, err := template.ParseFiles(lp, fp)
	if err != nil {
		// Log the detailed error
		log.Println(err.Error())
		// Return a generic "Internal Server Error" message
		http.Error(w, http.StatusText(500), 500)
		return
	}

	hostname, _ := os.Hostname()

	h := struct {
		Hostname string
	}{
		Hostname: hostname,
	}

	if err := tmpl.ExecuteTemplate(w, "layout", h); err != nil {
		log.Println(err.Error())
		http.Error(w, http.StatusText(500), 500)
	}
}
