package config

import "github.com/sirupsen/logrus"

type Config struct {
	Port   string
	Env    string
	Logger *logrus.Logger
}

func NewConfig(port, environment string) *Config {
	return &Config{
		Port:   port,
		Env:    environment,
		Logger: setDefaultLogger(environment),
	}
}

func setDefaultLogger(env string) *logrus.Logger {
	log := logrus.New()

	if env == "development" || env == "dev" {
		log.SetLevel(logrus.TraceLevel)
	} else if env == "debug" {
		log.SetLevel(logrus.DebugLevel)
	} else {
		log.SetLevel(logrus.InfoLevel)
	}

	log.SetFormatter(&logrus.TextFormatter{
		DisableColors: false,
		FullTimestamp: true,
	})
	return log
}
