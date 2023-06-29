package config

import (
	"atados/challenger/internal/model"
	"github.com/joho/godotenv"
	"go.uber.org/zap"
	"os"
	"strconv"
	"time"
)

type config struct {
	cfg    model.Config
	logger *zap.Logger
}

func NewConfig(logger *zap.Logger) *config {
	return &config{
		logger: logger,
	}
}

func (c *config) GetConfig(envFile string) *model.Config {

	err := godotenv.Load(envFile)
	if err != nil {
		zap.Error(err)
	}

	c.logger.Info("Init config",
		zap.Time("StartedAt", time.Now()),
	)
	port := os.Getenv("REST_PORT")

	if port != "" {
		portInt, err := strconv.Atoi(port)
		if err != nil {
			c.cfg.RestPort = 9090
		}
		c.cfg.RestPort = portInt

	} else {
		c.cfg.RestPort = 9090
	}

	c.logger.Debug("REST_PORT", zap.Int("REST_PORT", c.cfg.RestPort), zap.Time("StartedAt", time.Now()))

	c.cfg.DbConnString = os.Getenv("DB_CONNSTRING")

	c.logger.Info("Config loaded",
		zap.Time("StartedAt", time.Now()))

	return &c.cfg
}
