package main

import (
	"math/rand"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

var Url_Store = make(map[string]string)

const charset = "qwertyuiopasdfghjklzxcvbnmQWERTYUIOPASDFGHJKLZXCVBNM"

func generateShorturl(length int) string {
	rand.NewSource(time.Now().UnixMicro())
	short_arr := make([]byte, length)
	for i := 0; i < len(short_arr); i++ {
		short_arr[i] = charset[rand.Intn(len(charset))]

	}
	return string(short_arr)

}

func short_url_deal(w http.ResponseWriter, r *http.Request) {
	longU := r.FormValue("url")
	shortU := generateShorturl(rand.Intn(6))
	Url_Store[shortU] = longU
	w.Write([]byte("Shortened URL: localhost:8080/" + shortU))
}

func requestHandler(w http.ResponseWriter, r *http.Request) {
	t := mux.Vars(r)
	short_url := t["shortU"]
	longUrl, exists := Url_Store[short_url]
	if exists {
		http.Redirect(w, r, longUrl, http.StatusAccepted)
	} else {
		http.Error(w, "Not found", http.StatusNotFound)
	}

}
func main() {
	run := mux.NewRouter()
	run.HandleFunc("/shorten", short_url_deal).Methods("POST")
	run.HandleFunc("/{shorturl}", requestHandler).Methods("GET")
	http.ListenAndServe(":8080", run)
}
