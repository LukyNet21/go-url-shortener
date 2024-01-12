package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

type Url struct {
	Id       string    `json:"id"`
	Name     string    `json:"name"`
	Url      string    `json:"url"`
	ShortUrl string    `json:"shortUrl"`
	Created  time.Time `json:"created"`
}

var urls = []Url{
	{"1", "Google", "https://www.google.com", "ggl", time.Date(2019, 9, 13, 11, 45, 23, 0, time.Local)},
	{"2", "Facebook", "https://www.facebook.com", "fb", time.Date(2019, 9, 13, 11, 45, 23, 0, time.Local)},
	{"3", "Youtube", "https://www.youtube.com", "ytb", time.Date(2019, 9, 13, 11, 45, 23, 0, time.Local)},
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/shorten", shorten).Methods("POST")
	r.HandleFunc("/delete/{id}", deleteUrl).Methods("DELETE")
	r.HandleFunc("/urls", listUrls).Methods("GET")
	r.HandleFunc("/{url}", redirect).Methods("GET")

	fmt.Println("Server started on port 8080")
	http.ListenAndServe(":8080", r)
}

func shorten(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var u Url
	_ = json.NewDecoder(r.Body).Decode(&u)

	u.Id = randomString(80)
	u.ShortUrl = randomString(6)
	dt := time.Now()
	u.Created = dt

	for _, n := range urls {
		if n.Url == u.Url {
			w.WriteHeader(http.StatusConflict)
			w.Write([]byte("Url already exists"))
			return
		}
	}
	urls = append(urls, u)

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(u)
}

func redirect(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	url := vars["url"]

	for _, u := range urls {
		if u.ShortUrl == url {
			http.Redirect(w, r, u.Url, http.StatusMovedPermanently)
			return
		}
	}

	w.WriteHeader(http.StatusNotFound)
}

func listUrls(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(urls)
}

func randomString(n int) string {
	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

	s := make([]rune, n)
	for i := range s {
		s[i] = letters[rand.Intn(len(letters))]
	}

	for _, u := range urls {
		if u.Id == string(s) || u.ShortUrl == string(s) {
			return randomString(n)
		}
	}

	return string(s)
}

func deleteUrl(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	for i, u := range urls {
		if u.Id == id {
			urls = append(urls[:i], urls[i+1:]...)
			w.WriteHeader(http.StatusOK)
			return
		}
	}
	w.WriteHeader(http.StatusNotFound)
}
