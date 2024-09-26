package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strings"
	"text/template"
)

var urlStore = make(map[string]string)

func main() {

	http.HandleFunc("/", home)
	http.HandleFunc("/shorten", shorten)
	http.HandleFunc("/short/", redirect)

	fmt.Println("Server is running at http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}

func home(w http.ResponseWriter, r *http.Request) {
	render(w, "templates/form.html", nil)
}

func shorten(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	longURL := r.FormValue("url")
	shortKey := generateKey()
	urlStore[shortKey] = longURL

	shortenedURL := fmt.Sprintf("http://localhost:8080/short/%s", shortKey)
	render(w, "templates/result.html", map[string]string{"ShortURL": shortenedURL})
}

func redirect(w http.ResponseWriter, r *http.Request) {
	shortKey := strings.TrimPrefix(r.URL.Path, "/short/")
	if shortKey == "" {
		http.Error(w, "Shortened key is missing", http.StatusBadRequest)
		return
	}

	originalURL, found := urlStore[shortKey]
	if !found {
		http.Error(w, "Shortened key not found", http.StatusNotFound)
		return
	}

	// Redirect the user to the original URL
	http.Redirect(w, r, originalURL, http.StatusMovedPermanently)
}

func render(w http.ResponseWriter, filename string, data interface{}) {
	tmpl := template.Must(template.ParseFiles(filename))

	err := tmpl.Execute(w, data)
	if err != nil {
		log.Print(err)
		http.Error(w, "Sorry, something went wrong", http.StatusInternalServerError)
	}
}

func generateKey() string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	const keyLength = 6

	shortKey := make([]byte, keyLength)
	for i := range shortKey {
		shortKey[i] = charset[rand.Intn(len(charset))]
	}

	return string(shortKey)
}
