package main

import (
	"net/http"
	"strings"
)

type MemStorage struct {
	strg []string
}

func mainPage(w http.ResponseWriter, r *http.Request) {
	rawUrl := r.URL.String()
	splitURL := strings.Split(rawUrl, "/")
}

func main() {
	http.ListenAndServe("localhost:8080", nil)
	mx := http.NewServeMux()
	mx.HandleFunc("/", mainPage)
}
