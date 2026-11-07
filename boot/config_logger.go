package boot

import (
	"log"

	"github.com/bayik4/boilerplate-golang/internal/model"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

func LoadLoggerConfig(vpr *viper.Viper) model.ConfigLogger {
	vpr.SetConfigName("config")
	vpr.SetConfigType("yaml")
	vpr.AddConfigPath(".")
	viper.AutomaticEnv()

	if err := vpr.ReadInConfig(); err != nil {
		log.Fatal(
			"failed to read config file",
			zap.Error(err),
		)
	}

	return model.ConfigLogger{
		ServiceName: vpr.GetString("log.service_name"),
		Environment: vpr.GetString("app.env"),
		Level:       vpr.GetString("log.level"),
		Format:      vpr.GetString("log.format"),
	}
}
