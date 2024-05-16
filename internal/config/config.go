package config

import (
	"database/sql"
	"github.com/CP-Payne/ecommerce-server/internal/database"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"os"
)

type Config struct {
	Server *ServerConfig
	API    *ApiConfig
}

func NewConfig(port, environment string) *Config {

	logger := NewDefaultLogger(environment)

	return &Config{
		Server: &ServerConfig{
			Port:   port,
			Env:    environment,
			Logger: logger,
		},
		API: &ApiConfig{
			Logger: logger,
			DB:     NewDB(logger),
		},
	}
}

func NewDefaultLogger(env string) *logrus.Logger {
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

func NewDB(logger *logrus.Logger) *database.Queries {
	err := godotenv.Load()
	if err != nil {
		logger.WithFields(
			logrus.Fields{
				"err": err,
			}).Error("failed to load .env file")
	}
	connStr := os.Getenv("CONN_STR")

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		logger.WithFields(
			logrus.Fields{
				"err": err,
			}).Error("failed to open database connection")
	}

	return database.New(db)
}
