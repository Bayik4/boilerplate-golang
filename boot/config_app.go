package boot

import (
	"github.com/bayik4/boilerplate-golang/internal/model"
	"github.com/spf13/viper"
)

func LoadAppConfig(vpr *viper.Viper) model.AppConfig {
	return model.AppConfig{
		Name: vpr.GetString("app.name"),
		Host: vpr.GetString("app.host"),
		Port: vpr.GetString("app.port"),
	}
}