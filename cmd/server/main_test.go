package main

import (
	"fmt"
	"net/http"
	"testing"
)

func Test_mainHandler(t *testing.T) {
	go main()
	srv := "http://localhost:8080"
	req, _ := http.NewRequest("POST", fmt.Sprintf("%s/update/unknown/testCounter/100", srv), nil)
	req.Header.Set("Content-Type", "text/plain")
	http.DefaultClient.Do(req)
}
