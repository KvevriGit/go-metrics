package main

import (
	"github.com/KvevriGit/go-metrics/cmd/server/internal"
	"net/http"
)

// var GlobalStorage internal.MemStorage = internal.MemStorage{}

var GlobalStorage internal.MemStorage = internal.MemStorage{Values: make(map[string]float64)}

func mainHandler(res http.ResponseWriter, req *http.Request) {
	body := ""
	err := GlobalStorage.SaveMetric(req.URL.Path)
	if err != nil {
		switch err.Error() {
		case "no name":
			res.WriteHeader(http.StatusNotFound)
		default:
			res.WriteHeader(http.StatusBadRequest)
		}
	} else {
		res.WriteHeader(http.StatusOK)
	}
	res.Write([]byte(body))
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc(`/`, mainHandler)
	err := http.ListenAndServe(`:8080`, mux)
	if err != nil {
		panic(err)
	}
}
