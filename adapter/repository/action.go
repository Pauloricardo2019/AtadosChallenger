package repository

import (
	"atados/challenger/internal/model"
	"context"
	"errors"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type actionRepository struct {
	baseRepo *BaseRepository
	logger   *zap.Logger
}

func NewActionRepository(db *gorm.DB, logger *zap.Logger) *actionRepository {
	baseRepo := NewBaseRepository(db)
	return &actionRepository{
		baseRepo: baseRepo,
		logger:   logger,
	}

}

func (p *actionRepository) Create(ctx context.Context, action *model.Action) (*model.Action, error) {
	loggerUUID := ctx.Value("logger").(string)
	p.logger.Info("Repository: Creating action", zap.String("correlationID: ", loggerUUID))
	conn, err := p.baseRepo.getConnection(ctx)
	if err != nil {
		return nil, err
	}

	if err = conn.Create(action).Error; err != nil {
		return nil, err
	}

	p.logger.Debug("Action", zap.Any("action", action), zap.String("correlationID: ", loggerUUID))
	p.logger.Info("Repository: Action created", zap.String("correlationID: ", loggerUUID))
	return action, nil
}

func (p *actionRepository) GetCount(ctx context.Context) (int64, error) {
	loggerUUID := ctx.Value("logger").(string)
	p.logger.Info("Repository: Getting count", zap.String("correlationID: ", loggerUUID))
	conn, err := p.baseRepo.getConnection(ctx)
	if err != nil {
		return 0, err
	}

	actions := make([]model.Action, 0)
	var count int64

	if err = conn.Find(&actions).Count(&count).Error; err != nil {
		return 0, err
	}

	p.logger.Debug("Count", zap.Int64("count", count), zap.String("correlationID: ", loggerUUID))
	p.logger.Info("Repository: Count gotten", zap.String("correlationID: ", loggerUUID))

	return count, nil
}

func (p *actionRepository) GetByID(ctx context.Context, id uint64) (bool, *model.Action, error) {
	loggerUUID := ctx.Value("logger").(string)
	p.logger.Info("Repository: Getting action by ID", zap.String("correlationID: ", loggerUUID))
	conn, err := p.baseRepo.getConnection(ctx)
	if err != nil {
		return false, nil, err
	}

	action := &model.Action{}

	if err = conn.Where(&model.Action{
		ID: id,
	}).First(action).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return false, action, nil
		}
		return false, nil, err
	}
	p.logger.Debug("Action", zap.Any("action", action), zap.Uint64("id", id), zap.Bool("found", true), zap.String("correlationID: ", loggerUUID))
	p.logger.Info("Repository: Action gotten by ID", zap.String("correlationID: ", loggerUUID))

	return true, action, nil
}

func (p *actionRepository) GetAll(ctx context.Context, limit, offset int) ([]model.Action, error) {
	loggerUUID := ctx.Value("logger").(string)
	p.logger.Info("Repository: Getting all actions", zap.String("correlationID: ", loggerUUID))
	conn, err := p.baseRepo.getConnection(ctx)
	if err != nil {
		return nil, err
	}

	actions := make([]model.Action, 0)

	if err = conn.
		Limit(limit).
		Offset(offset).
		Find(&actions).Error; err != nil {
		return nil, err
	}
	p.logger.Debug("Actions", zap.Any("actions", actions), zap.Int("limit", limit), zap.Int("offset", offset), zap.String("correlationID: ", loggerUUID))
	p.logger.Info("Repository: Actions gotten", zap.String("correlationID: ", loggerUUID))

	return actions, nil
}

func (p *actionRepository) Update(ctx context.Context, action *model.Action) (*model.Action, error) {
	loggerUUID := ctx.Value("logger").(string)
	p.logger.Info("Repository: Updating action", zap.String("correlationID: ", loggerUUID))
	conn, err := p.baseRepo.getConnection(ctx)
	if err != nil {
		return nil, err
	}

	if err = conn.Debug().Save(action).Error; err != nil {
		return nil, err
	}
	p.logger.Debug("Action", zap.Any("action", action), zap.String("correlationID: ", loggerUUID))
	p.logger.Info("Repository: Action updated", zap.String("correlationID: ", loggerUUID))

	return action, nil
}

func (p *actionRepository) Delete(ctx context.Context, id uint64) error {
	loggerUUID := ctx.Value("logger").(string)
	p.logger.Info("Repository: Deleting action", zap.String("correlationID: ", loggerUUID))
	conn, err := p.baseRepo.getConnection(ctx)
	if err != nil {
		return err
	}

	p.logger.Debug("ID", zap.Uint64("id", id), zap.String("correlationID: ", loggerUUID))
	if err = conn.Delete(&model.Action{
		ID: id,
	}).Error; err != nil {
		return err
	}

	p.logger.Info("Repository: Action deleted", zap.String("correlationID: ", loggerUUID))

	return nil
}
