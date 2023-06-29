package provider

import (
	"atados/challenger/internal/model"
	"go.uber.org/zap"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type databaseProvider struct {
	config *model.Config
	logger *zap.Logger
}

func NewDatabaseProvider(config *model.Config, logger *zap.Logger) *databaseProvider {
	return &databaseProvider{
		config: config,
		logger: logger,
	}
}

func (t *databaseProvider) GetConnection() (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(t.config.DbConnString), &gorm.Config{})
	if err != nil {
		t.logger.Error("Error on get connection", zap.Error(err))
		return nil, err
	}
	return db, err
}
