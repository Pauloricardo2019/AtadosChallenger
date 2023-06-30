package repository

import (
	"atados/challenger/internal/model"
	"context"
	"errors"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type voluntaryRepository struct {
	baseRepo *BaseRepository
	logger   *zap.Logger
}

func NewVoluntaryRepository(db *gorm.DB, logger *zap.Logger) *voluntaryRepository {
	baseRepo := NewBaseRepository(db)
	return &voluntaryRepository{
		baseRepo: baseRepo,
		logger:   logger,
	}
}

func (p *voluntaryRepository) Create(ctx context.Context, voluntary *model.Voluntary) (*model.Voluntary, error) {
	loggerUUID := ctx.Value("logger").(string)
	p.logger.Info("Repository: Creating voluntary", zap.String("correlationID: ", loggerUUID))
	conn, err := p.baseRepo.getConnection(ctx)
	if err != nil {
		return nil, err
	}

	if err = conn.Create(voluntary).Error; err != nil {
		return nil, err
	}

	p.logger.Debug("Voluntary", zap.Any("voluntary", voluntary), zap.String("correlationID: ", loggerUUID))
	p.logger.Info("Repository: Voluntary created", zap.String("correlationID: ", loggerUUID))
	return voluntary, nil
}

func (p *voluntaryRepository) GetCount(ctx context.Context) (int64, error) {
	loggerUUID := ctx.Value("logger").(string)
	p.logger.Info("Repository: Getting count", zap.String("correlationID: ", loggerUUID))
	conn, err := p.baseRepo.getConnection(ctx)
	if err != nil {
		return 0, err
	}

	voluntarys := make([]model.Voluntary, 0)
	var count int64

	if err = conn.Find(&voluntarys).Count(&count).Error; err != nil {
		return 0, err
	}

	p.logger.Debug("Count", zap.Int64("count", count), zap.String("correlationID: ", loggerUUID))
	p.logger.Info("Repository: Count gotten", zap.String("correlationID: ", loggerUUID))

	return count, nil
}

func (p *voluntaryRepository) GetByID(ctx context.Context, id uint64) (bool, *model.Voluntary, error) {
	loggerUUID := ctx.Value("logger").(string)
	p.logger.Info("Repository: Getting voluntary by ID", zap.String("correlationID: ", loggerUUID))
	conn, err := p.baseRepo.getConnection(ctx)
	if err != nil {
		return false, nil, err
	}

	voluntary := &model.Voluntary{}

	if err = conn.Where(&model.Voluntary{
		ID: id,
	}).First(voluntary).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return false, voluntary, nil
		}
		return false, nil, err
	}
	p.logger.Debug("Voluntary", zap.Any("voluntary", voluntary), zap.Uint64("id", id), zap.Bool("found", true), zap.String("correlationID: ", loggerUUID))
	p.logger.Info("Repository: Voluntary gotten by ID", zap.String("correlationID: ", loggerUUID))

	return true, voluntary, nil
}

func (p *voluntaryRepository) GetAll(ctx context.Context, limit, offset int) ([]model.Voluntary, error) {
	loggerUUID := ctx.Value("logger").(string)
	p.logger.Info("Repository: Getting all voluntarys", zap.String("correlationID: ", loggerUUID))
	conn, err := p.baseRepo.getConnection(ctx)
	if err != nil {
		return nil, err
	}

	voluntarys := make([]model.Voluntary, 0)

	if err = conn.
		Limit(limit).
		Offset(offset).
		Find(&voluntarys).Error; err != nil {
		return nil, err
	}
	p.logger.Debug("Voluntaries", zap.Any("voluntarys", voluntarys), zap.Int("limit", limit), zap.Int("offset", offset), zap.String("correlationID: ", loggerUUID))
	p.logger.Info("Repository: Voluntaries gotten", zap.String("correlationID: ", loggerUUID))

	return voluntarys, nil
}

func (p *voluntaryRepository) Update(ctx context.Context, voluntary *model.Voluntary) (*model.Voluntary, error) {
	loggerUUID := ctx.Value("logger").(string)
	p.logger.Info("Repository: Updating voluntary", zap.String("correlationID: ", loggerUUID))
	conn, err := p.baseRepo.getConnection(ctx)
	if err != nil {
		return nil, err
	}

	if err = conn.Debug().Save(voluntary).Error; err != nil {
		return nil, err
	}
	p.logger.Debug("Voluntary", zap.Any("voluntary", voluntary), zap.String("correlationID: ", loggerUUID))
	p.logger.Info("Repository: Voluntary updated", zap.String("correlationID: ", loggerUUID))

	return voluntary, nil
}

func (p *voluntaryRepository) Delete(ctx context.Context, id uint64) error {
	loggerUUID := ctx.Value("logger").(string)
	p.logger.Info("Repository: Deleting voluntary", zap.String("correlationID: ", loggerUUID))
	conn, err := p.baseRepo.getConnection(ctx)
	if err != nil {
		return err
	}

	p.logger.Debug("ID", zap.Uint64("id", id), zap.String("correlationID: ", loggerUUID))
	if err = conn.Delete(&model.Voluntary{
		ID: id,
	}).Error; err != nil {
		return err
	}

	p.logger.Info("Repository: Voluntary deleted", zap.String("correlationID: ", loggerUUID))

	return nil
}
