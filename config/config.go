package config

import (
	"fmt"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
	gormpg "go-template-microservice-v2/pkg/gorm_pg"
	echoserver "go-template-microservice-v2/pkg/http/server"
	"os"
)

type Config struct {
	ServiceName string                 `mapstructure:"serviceName"`
	Echo        *echoserver.EchoConfig `mapstructure:"echo"`
	PgConfig    *gormpg.PgConfig       `mapstructure:"pgConfig"`
}

func NewConfig() (*Config, *echoserver.EchoConfig, *gormpg.PgConfig, error) {
	env := os.Getenv("APP_ENV")
	if env == "" {
		env = "development"
	}

	cfg := &Config{}

	viper.SetConfigName(fmt.Sprintf("config.%s", env))
	viper.AddConfigPath("./config/")
	viper.SetConfigType("json")

	if err := viper.ReadInConfig(); err != nil {
		return nil, nil, nil, errors.Wrap(err, "viper.ReadInConfig")
	}

	if err := viper.Unmarshal(cfg); err != nil {
		return nil, nil, nil, errors.Wrap(err, "viper.Unmarshal")
	}

	return cfg, cfg.Echo, cfg.PgConfig, nil
}
