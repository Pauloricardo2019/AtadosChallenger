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

func init() {
	logger, _ = zap.NewDevelopment()
}

func TestActionRepository_Create(t *testing.T) {
	ctx := context.Background()

	cfg := config.NewConfig(logger).GetConfig(envPath)
	db, err := gorm.Open(postgres.Open(cfg.DbConnString), &gorm.Config{})
	assert.NoError(t, err)

	actionToCreate := &model.Action{
		Name:         "Action test",
		Institution:  "Institution fake",
		City:         "São Paulo",
		Neighborhood: "Limoeiro",
		Address:      "Rua Palmeira, 25",
		Description:  "Reuniao as 15 horas",
	}

	actionRepository := NewActionRepository(db, logger)

	action, err := actionRepository.Create(ctx, actionToCreate)
	assert.NoError(t, err)
	assert.True(t, action.ID > 0)

}

func TestActionRepository_GetCount(t *testing.T) {
	ctx := context.Background()

	cfg := config.NewConfig(logger).GetConfig(envPath)
	db, err := gorm.Open(postgres.Open(cfg.DbConnString), &gorm.Config{})
	assert.NoError(t, err)

	actionToCreate := &model.Action{
		Name:         "Action test",
		Institution:  "Institution fake",
		City:         "São Paulo",
		Neighborhood: "Limoeiro",
		Address:      "Rua Palmeira, 25",
		Description:  "Reuniao as 15 horas",
	}

	actionRepository := NewActionRepository(db, logger)

	action, err := actionRepository.Create(ctx, actionToCreate)
	assert.NoError(t, err)
	assert.True(t, action.ID > 0)

	count, err := actionRepository.GetCount(ctx)
	assert.NoError(t, err)
	assert.True(t, count > 0)

}

func TestActionRepository_GetByID(t *testing.T) {
	ctx := context.Background()

	cfg := config.NewConfig(logger).GetConfig(envPath)
	db, err := gorm.Open(postgres.Open(cfg.DbConnString), &gorm.Config{})
	assert.NoError(t, err)

	actionToCreate := &model.Action{
		Name:         "Action test",
		Institution:  "Institution fake",
		City:         "São Paulo",
		Neighborhood: "Limoeiro",
		Address:      "Rua Palmeira, 25",
		Description:  "Reuniao as 15 horas",
	}

	actionRepository := NewActionRepository(db, logger)

	action, err := actionRepository.Create(ctx, actionToCreate)
	assert.NoError(t, err)
	assert.True(t, action.ID > 0)

	found, actionFound, err := actionRepository.GetByID(ctx, action.ID)
	assert.NoError(t, err)
	assert.True(t, found)
	assert.True(t, actionFound.ID == actionToCreate.ID)

}

func TestActionRepository_GetAll(t *testing.T) {
	ctx := context.Background()

	cfg := config.NewConfig(logger).GetConfig(envPath)
	db, err := gorm.Open(postgres.Open(cfg.DbConnString), &gorm.Config{})
	assert.NoError(t, err)

	actionToCreate := &model.Action{
		Name:         "Action test",
		Institution:  "Institution fake",
		City:         "São Paulo",
		Neighborhood: "Limoeiro",
		Address:      "Rua Palmeira, 25",
		Description:  "Reuniao as 15 horas",
	}

	limit := 10
	offset := 0

	actionRepository := NewActionRepository(db, logger)

	action, err := actionRepository.Create(ctx, actionToCreate)
	assert.NoError(t, err)
	assert.True(t, action.ID > 0)

	voluntaries, err := actionRepository.GetAll(ctx, limit, offset)
	assert.NoError(t, err)
	assert.True(t, len(voluntaries) > 0)

}

func TestActionRepository_Update(t *testing.T) {
	ctx := context.Background()

	cfg := config.NewConfig(logger).GetConfig(envPath)
	db, err := gorm.Open(postgres.Open(cfg.DbConnString), &gorm.Config{})
	assert.NoError(t, err)

	actionToCreate := &model.Action{
		Name:         "Action test",
		Institution:  "Institution fake",
		City:         "São Paulo",
		Neighborhood: "Limoeiro",
		Address:      "Rua Palmeira, 25",
		Description:  "Reuniao as 15 horas",
	}

	actionRepository := NewActionRepository(db, logger)

	actionCreated, err := actionRepository.Create(ctx, actionToCreate)
	assert.NoError(t, err)
	assert.True(t, actionToCreate.ID > 0)

	actionCreated.Name = "Action São Miguel"
	actionCreated.Institution = "Instituto São Miguel"
	actionCreated.City = "São Paulo"
	actionCreated.Neighborhood = "Mangueira"
	actionCreated.Address = "Rua das flores, 544"
	actionCreated.Description = "Reuniao as 19 horas"

	actionChanged, err := actionRepository.Update(ctx, actionCreated)
	assert.NoError(t, err)
	assert.Equal(t, "Action São Miguel", actionChanged.Name)
	assert.Equal(t, "Instituto São Miguel", actionChanged.Institution)
	assert.Equal(t, "São Paulo", actionChanged.City)
	assert.Equal(t, "Mangueira", actionChanged.Neighborhood)
	assert.Equal(t, "Rua das flores, 544", actionChanged.Address)
	assert.Equal(t, "Reuniao as 19 horas", actionChanged.Description)

}

func TestActionRepository_Delete(t *testing.T) {
	ctx := context.Background()

	cfg := config.NewConfig(logger).GetConfig(envPath)
	db, err := gorm.Open(postgres.Open(cfg.DbConnString), &gorm.Config{})
	assert.NoError(t, err)

	actionToCreate := &model.Action{
		Name:         "Action test",
		Institution:  "Institution fake",
		City:         "São Paulo",
		Neighborhood: "Limoeiro",
		Address:      "Rua Palmeira, 25",
		Description:  "Reuniao as 15 horas",
	}

	actionRepository := NewActionRepository(db, logger)

	action, err := actionRepository.Create(ctx, actionToCreate)
	assert.NoError(t, err)
	assert.True(t, action.ID > 0)

	err = actionRepository.Delete(ctx, action.ID)
	assert.NoError(t, err)

	found, _, err := actionRepository.GetByID(ctx, action.ID)
	assert.NoError(t, err)
	assert.False(t, found)
}
