package settings

type EnvironmentConfig struct {
	address         string `env:"ADDRESS"`
	report_interval int    `env:"REPORT_INTERVAL"`
	poll_interval   int    `env:"POLL_INTERVAL"`
}

type Poll struct {
	PollCount      int
	pollInterval   int    `default0:"2"`
	reportInterval int    `default0:"10"`
	RandomValue    uint64 // aka gauge
}
