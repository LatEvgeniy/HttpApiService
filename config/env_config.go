package config

import (
	"github.com/caarlos0/env/v10"
	logger "github.com/sirupsen/logrus"
)

type EnvConfig struct {
	RabbitHost     string `env:"RABBITMQ_HOST" envDefault:"rabbitmq"`
	RabbitPort     string `env:"RABBITMQ_PORT" envDefault:"5672"`
	RabbitUser     string `env:"RABBITMQ_USER" envDefault:"guest"`
	RabbitPassword string `env:"RABBITMQ_PASS" envDefault:"guest"`
}

func GetCfg() *EnvConfig {
	envCfg := &EnvConfig{}
	if err := env.Parse(envCfg); err != nil {
		panic(err)
	}
	logger.Infof("Running with evn cfg: %+v", *envCfg)

	return envCfg
}
