package main

import (
	"flag"
	"fmt"
	"github.com/KvevriGit/go-metrics/cmd/agent/internal"
	"github.com/KvevriGit/go-metrics/cmd/agent/internal/settings"
	"github.com/caarlos0/env"
	"log"
	"net/http"
	"runtime"
	"time"
)

var (
	addrFlag       = flag.String("a", "http://localhost:8080", "address of server with port")
	reportInterval = flag.Int("r", 10, "report interval")
	pollInterval   = flag.Int("p", 2, "poll interval")
)

var metricsNames []string = []string{"Alloc",
	"BuckHashSys", "Frees", "GCCPUFraction",
	"GCSys", "HeapAlloc", "HeapIdle",
	"HeapInuse", "HeapObjects", "HeapReleased",
	"HeapSys", "LastGC", "Lookups", "MCacheInuse",
	"MCacheSys", "MSpanInuse", "MSpanSys", "Mallocs",
	"NextGC", "NumForcedGC", "NumGC", "OtherSys",
	"PauseTotalNs", "StackInuse", "StackSys", "Sys", "TotalAlloc"}

func readStats(s *runtime.MemStats) {
	runtime.ReadMemStats(s)
}

func sendAttribute(attributeType string, name string, value string) {
	srv := addrFlag
	req, err := http.NewRequest("POST", fmt.Sprintf("%s/update/%s/%s/%s", *srv, attributeType, name, value), nil)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("Content-Type", "text/plain")
	http.DefaultClient.Do(req)
}

func main() {
	var Poll settings.Poll
	var envConfig settings.EnvironmentConfig
	err := env.Parse(&envConfig) //прочитать параметры окружения в структуру
	if err != nil {
		log.Fatal(err)
	}

	flag.Parse() // прочитать все флаги

	memoryStats := new(runtime.MemStats)
	readStats(memoryStats) // read once
	go func() {
		for {
			time.Sleep(time.Duration(*pollInterval) * time.Second)
			readStats(memoryStats)
			Poll.PollCount++
			println("poll")
		}
	}()
	for {
		for _, metricName := range metricsNames {
			reflectionValueOfStat := internal.GetMetricByName(memoryStats, metricName)
			value, metricType, _ := internal.GetValueAndTypeFromReflection(reflectionValueOfStat)
			sendAttribute(metricType, metricName, value)
		}
		time.Sleep(time.Duration(*reportInterval) * time.Second)
		println("report")
	}
}
