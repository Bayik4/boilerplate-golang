package model

type ConfigTelemetry struct {
	Enabled       bool
	ServiceName   string
	Environment   string
	Endpoint      string
	Insecure      bool
	SamplingRatio float64
}
