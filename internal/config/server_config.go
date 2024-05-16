package config

import "github.com/sirupsen/logrus"

type ServerConfig struct {
	Port   string
	Env    string
	Logger *logrus.Logger
}
