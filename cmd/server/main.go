package main

import (
	"github.com/KvevriGit/go-metrics/cmd/server/internal"
	"net/http"
)

// var GlobalStorage internal.MemStorage = internal.MemStorage{}

var GlobalStorage internal.MemStorage = internal.MemStorage{Values: make(map[string]float64)}

func ErrorComp(err error) func(err2 error) bool {
	return func(err2 error) bool {
		return err.Error() == err2.Error()
	}
}

func mainHandler(res http.ResponseWriter, req *http.Request) {
	body := ""
	err := GlobalStorage.SaveMetric(req.URL.Path)
	ThatError := ErrorComp(err)
	switch true {
	case err == nil:
		res.WriteHeader(http.StatusOK)
	case ThatError(internal.ErrMap[404]):
		res.WriteHeader(http.StatusNotFound)
	case ThatError(internal.ErrMap[501]):
		res.WriteHeader(http.StatusNotImplemented)
	case ThatError(internal.ErrMap[400]):
		res.WriteHeader(http.StatusBadRequest)
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
