package boot

import (
	"github.com/bayik4/boilerplate-golang/internal/model"
	"github.com/spf13/viper"
)

func LoadTelemetryConfig(vpr *viper.Viper) model.ConfigTelemetry {
	return model.ConfigTelemetry{
		Enabled:       vpr.GetBool("otel.enabled"),
		ServiceName:   vpr.GetString("otel.service_name"),
		Environment:   vpr.GetString("app.env"),
		Endpoint:      vpr.GetString("otel.exporter_otlp_endpoint"),
		Insecure:      viper.GetBool("otel.exporter_otlp_insecure"),
		SamplingRatio: vpr.GetFloat64("otel.sampling_ratio"),
	}
}
