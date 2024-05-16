package config

import (
	"github.com/CP-Payne/ecommerce-server/internal/database"
	"github.com/sirupsen/logrus"
)

type ApiConfig struct {
	Logger *logrus.Logger
	DB     *database.Queries
}
