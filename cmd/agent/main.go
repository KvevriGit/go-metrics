package main

import (
	"errors"
	"fmt"
	"net/http"
	"reflect"
	"runtime"
	"strconv"
)

func allocate() {
	//  0.25MB
	_ = make([]byte, int((1<<20)*0.25))
}

func getAttr(obj interface{}, fieldName string) reflect.Value {
	pointToStruct := reflect.ValueOf(obj) // addressable
	curStruct := pointToStruct.Elem()
	if curStruct.Kind() != reflect.Struct {
		panic("not struct")
	}
	curField := curStruct.FieldByName(fieldName) // type: reflect.Value
	if !curField.IsValid() {
		panic("not found:" + fieldName)
	}
	return curField
}

var tributes []string = []string{"Alloc",
	"BuckHashSys", "Frees", "GCCPUFraction",
	"GCSys", "HeapAlloc", "HeapIdle",
	"HeapInuse", "HeapObjects", "HeapReleased",
	"HeapSys", "LastGC", "Lookups", "MCacheInuse",
	"MCacheSys", "MSpanInuse", "MSpanSys", "Mallocs",
	"NextGC", "NumForcedGC", "NumGC", "OtherSys",
	"PauseTotalNs", "StackInuse", "StackSys", "Sys", "TotalAlloc"}

type Poll struct {
	PollCount      int
	pollInterval   int    `default0:"2"`
	reportInterval int    `default0:"10"`
	RandomValue    uint64 // aka gauge
}

func unpkgValue(value reflect.Value) (string, string, error) {
	switch value.Kind().String() {
	case "float64":
		return fmt.Sprintf("%f", value.Interface().(float64)), "gauge", nil
	case "uint64":
		return strconv.FormatUint(value.Interface().(uint64), 10), "counter", nil // https://go.dev/blog/laws-of-reflection
	case "uint32":
		return strconv.FormatUint(uint64(value.Interface().(uint32)), 10), "counter", nil // https://go.dev/blog/laws-of-reflection
	}
	return "", "", errors.New("unexpected value type") // что возвращать?
}

func readStats() *runtime.MemStats {
	memstats := new(runtime.MemStats)
	runtime.ReadMemStats(memstats)
	return memstats
}

func SendTribute(t string, n string, v string) {
	srv := "http://localhost:8080"
	req, _ := http.NewRequest("POST", fmt.Sprintf("%s/update/%s/%s/%s", srv, t, n, v), nil)
	req.Header.Set("Content-Type", "text/plain")
	http.DefaultClient.Do(req)
}

func main() {
	allocate()
	stats := readStats()
	for _, atTribute := range tributes {
		attrValue := getAttr(stats, atTribute)
		v, t, _ := unpkgValue(attrValue)
		SendTribute(t, atTribute, v)
	}
}
