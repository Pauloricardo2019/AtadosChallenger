package repository

import (
	"atados/challenger/internal/config"
	"atados/challenger/internal/model"
	"context"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"testing"
)

var logger *zap.Logger
var envPath = "../../dev.env"

func init() {
	logger, _ = zap.NewDevelopment()
}

func TestVoluntaryRepository_Create(t *testing.T) {
	ctx := context.Background()

	cfg := config.NewConfig(logger).GetConfig(envPath)
	db, err := gorm.Open(postgres.Open(cfg.DbConnString), &gorm.Config{})
	assert.NoError(t, err)

	voluntaryToCreate := &model.Voluntary{
		FirstName:    "Miguel",
		LastName:     "Ferreira",
		Neighborhood: "Centro",
		City:         "Rio de Janeiro",
	}

	voluntaryRepository := NewVoluntaryRepository(db, logger)

	voluntary, err := voluntaryRepository.Create(ctx, voluntaryToCreate)
	assert.NoError(t, err)
	assert.True(t, voluntary.ID > 0)

}

func TestVoluntaryRepository_GetCount(t *testing.T) {
	ctx := context.Background()

	cfg := config.NewConfig(logger).GetConfig(envPath)
	db, err := gorm.Open(postgres.Open(cfg.DbConnString), &gorm.Config{})
	assert.NoError(t, err)

	voluntaryToCreate := &model.Voluntary{
		FirstName:    "Miguel",
		LastName:     "Ferreira",
		Neighborhood: "Centro",
		City:         "Rio de Janeiro",
	}

	voluntaryRepository := NewVoluntaryRepository(db, logger)

	voluntary, err := voluntaryRepository.Create(ctx, voluntaryToCreate)
	assert.NoError(t, err)
	assert.True(t, voluntary.ID > 0)

	count, err := voluntaryRepository.GetCount(ctx)
	assert.NoError(t, err)
	assert.True(t, count > 0)

}

func TestVoluntaryRepository_GetByID(t *testing.T) {
	ctx := context.Background()

	cfg := config.NewConfig(logger).GetConfig(envPath)
	db, err := gorm.Open(postgres.Open(cfg.DbConnString), &gorm.Config{})
	assert.NoError(t, err)

	voluntaryToCreate := &model.Voluntary{
		FirstName:    "Miguel",
		LastName:     "Ferreira",
		Neighborhood: "Centro",
		City:         "Rio de Janeiro",
	}

	voluntaryRepository := NewVoluntaryRepository(db, logger)

	voluntary, err := voluntaryRepository.Create(ctx, voluntaryToCreate)
	assert.NoError(t, err)
	assert.True(t, voluntary.ID > 0)

	found, voluntaryFound, err := voluntaryRepository.GetByID(ctx, voluntary.ID)
	assert.NoError(t, err)
	assert.True(t, found)
	assert.True(t, voluntaryFound.ID == voluntaryToCreate.ID)

}

func TestVoluntaryRepository_GetAll(t *testing.T) {
	ctx := context.Background()

	cfg := config.NewConfig(logger).GetConfig(envPath)
	db, err := gorm.Open(postgres.Open(cfg.DbConnString), &gorm.Config{})
	assert.NoError(t, err)

	voluntaryToCreate := &model.Voluntary{
		FirstName:    "Miguel",
		LastName:     "Ferreira",
		Neighborhood: "Centro",
		City:         "Rio de Janeiro",
	}

	limit := 10
	offset := 0

	voluntaryRepository := NewVoluntaryRepository(db, logger)

	voluntary, err := voluntaryRepository.Create(ctx, voluntaryToCreate)
	assert.NoError(t, err)
	assert.True(t, voluntary.ID > 0)

	voluntaries, err := voluntaryRepository.GetAll(ctx, limit, offset)
	assert.NoError(t, err)
	assert.True(t, len(voluntaries) > 0)

}

func TestVoluntaryRepository_Update(t *testing.T) {
	ctx := context.Background()

	cfg := config.NewConfig(logger).GetConfig(envPath)
	db, err := gorm.Open(postgres.Open(cfg.DbConnString), &gorm.Config{})
	assert.NoError(t, err)

	voluntaryToCreate := &model.Voluntary{
		FirstName:    "Miguel",
		LastName:     "Ferreira",
		Neighborhood: "Centro",
		City:         "Rio de Janeiro",
	}

	voluntaryRepository := NewVoluntaryRepository(db, logger)

	voluntaryCreated, err := voluntaryRepository.Create(ctx, voluntaryToCreate)
	assert.NoError(t, err)
	assert.True(t, voluntaryToCreate.ID > 0)

	voluntaryCreated.FirstName = "Caio"
	voluntaryCreated.LastName = "Matos"
	voluntaryCreated.Neighborhood = "Limoeiro"
	voluntaryCreated.City = "São Paulo"

	voluntaryChanged, err := voluntaryRepository.Update(ctx, voluntaryCreated)
	assert.NoError(t, err)
	assert.Equal(t, "Caio", voluntaryChanged.FirstName)
	assert.Equal(t, "Matos", voluntaryChanged.LastName)
	assert.Equal(t, "Limoeiro", voluntaryChanged.Neighborhood)
	assert.Equal(t, "São Paulo", voluntaryChanged.City)

}

func TestVoluntaryRepository_Delete(t *testing.T) {
	ctx := context.Background()

	cfg := config.NewConfig(logger).GetConfig(envPath)
	db, err := gorm.Open(postgres.Open(cfg.DbConnString), &gorm.Config{})
	assert.NoError(t, err)

	voluntaryToCreate := &model.Voluntary{
		FirstName:    "Miguel",
		LastName:     "Ferreira",
		Neighborhood: "Centro",
		City:         "Rio de Janeiro",
	}

	voluntaryRepository := NewVoluntaryRepository(db, logger)

	voluntary, err := voluntaryRepository.Create(ctx, voluntaryToCreate)
	assert.NoError(t, err)
	assert.True(t, voluntary.ID > 0)

	err = voluntaryRepository.Delete(ctx, voluntary.ID)
	assert.NoError(t, err)

	found, _, err := voluntaryRepository.GetByID(ctx, voluntary.ID)
	assert.NoError(t, err)
	assert.False(t, found)
}
