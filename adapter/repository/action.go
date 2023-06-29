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
	p.logger.Info("Repository: Creating action")
	conn, err := p.baseRepo.getConnection(ctx)
	if err != nil {
		return nil, err
	}

	if err = conn.Create(action).Error; err != nil {
		return nil, err
	}

	p.logger.Debug("Action", zap.Any("action", action))
	p.logger.Info("Repository: Action created")
	return action, nil
}

func (p *actionRepository) GetCount(ctx context.Context) (int64, error) {
	p.logger.Info("Repository: Getting count")
	conn, err := p.baseRepo.getConnection(ctx)
	if err != nil {
		return 0, err
	}

	actions := make([]model.Action, 0)
	var count int64

	if err = conn.Find(&actions).Count(&count).Error; err != nil {
		return 0, err
	}

	p.logger.Debug("Count", zap.Int64("count", count))
	p.logger.Info("Repository: Count gotten")

	return count, nil
}

func (p *actionRepository) GetByID(ctx context.Context, id uint64) (bool, *model.Action, error) {
	p.logger.Info("Repository: Getting action by ID")
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
	p.logger.Debug("Action", zap.Any("action", action), zap.Uint64("id", id), zap.Bool("found", true))
	p.logger.Info("Repository: Action gotten by ID")

	return true, action, nil
}

func (p *actionRepository) GetAll(ctx context.Context, limit, offset int) ([]model.Action, error) {
	p.logger.Info("Repository: Getting all actions")
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
	p.logger.Debug("Actions", zap.Any("actions", actions), zap.Int("limit", limit), zap.Int("offset", offset))
	p.logger.Info("Repository: Actions gotten")

	return actions, nil
}

func (p *actionRepository) Update(ctx context.Context, action *model.Action) (*model.Action, error) {
	p.logger.Info("Repository: Updating action")
	conn, err := p.baseRepo.getConnection(ctx)
	if err != nil {
		return nil, err
	}

	if err = conn.Debug().Save(action).Error; err != nil {
		return nil, err
	}
	p.logger.Debug("Action", zap.Any("action", action))
	p.logger.Info("Repository: Action updated")

	return action, nil
}

func (p *actionRepository) Delete(ctx context.Context, id uint64) error {
	p.logger.Info("Repository: Deleting action")
	conn, err := p.baseRepo.getConnection(ctx)
	if err != nil {
		return err
	}

	p.logger.Debug("ID", zap.Uint64("id", id))
	if err = conn.Delete(&model.Action{
		ID: id,
	}).Error; err != nil {
		return err
	}

	p.logger.Info("Repository: Action deleted")

	return nil
}
