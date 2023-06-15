package main

import (
	"fmt"
	"github.com/KvevriGit/go-metrics/cmd/server/internal"
	"github.com/caarlos0/env"
	"github.com/go-chi/chi/v5"
	"html/template"
	"log"
	"net/http"
)

// var GlobalStorage internal.MemStorage = internal.MemStorage{}

var GlobalStorage internal.MemStorage = internal.MemStorage{Values: make(map[string]float64)}

type EnvConfig struct {
	address string `env:"ADDRESS"`
}

func ErrorComp(err error) func(err2 error) bool {
	return func(err2 error) bool {
		return err.Error() == err2.Error()
	}
}

func saveMetricHandler(res http.ResponseWriter, req *http.Request) {
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

func listAllMetricsHandler(res http.ResponseWriter, req *http.Request) {
	templateHTML, _ := template.New("data").Parse("<body>{{range $index, $element :=.Values}}<h1>{{$index}} = {{$element}}</h1>{{end}}</body>")
	templateHTML.Execute(res, GlobalStorage)
}

func getSpecificMetricHandler(res http.ResponseWriter, req *http.Request) {
	//typeV := chi.URLParam(req, "type")
	nameV := chi.URLParam(req, "name")

	res.Header().Set("Content-Type", "text/plain")
	if valueM, ok := GlobalStorage.Values[nameV]; ok {
		res.WriteHeader(http.StatusOK)
		res.Write([]byte(fmt.Sprintf("%f", valueM)))

	} else {
		res.WriteHeader(http.StatusNotFound)
	}
}

func main() {
	var cfg EnvConfig
	err := env.Parse(&cfg)
	if err != nil {
		log.Fatal(err)
	}

	r := chi.NewRouter()
	r.Route("/", func(r chi.Router) {
		r.Get("/value/{type}/{name}", getSpecificMetricHandler)
		r.Get("/", listAllMetricsHandler)
		//r.Post("/update", saveMetricHandler)
	})
	r.NotFound(saveMetricHandler)
	err = http.ListenAndServe(`:8080`, r)
	if err != nil {
		panic(err)
	}
}
