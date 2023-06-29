package repository

import (
	"context"
	"gorm.io/gorm"
)

const DbTransactionKey string = "transaction_key"

type BaseRepository struct {
	db *gorm.DB
}

func NewBaseRepository(db *gorm.DB) *BaseRepository {
	return &BaseRepository{
		db: db,
	}
}

func (b *BaseRepository) getTransaction(ctx context.Context) (*gorm.DB, error) {
	connWithTransaction := ctx.Value(DbTransactionKey)
	if connWithTransaction == nil {
		return nil, nil
	}
	return connWithTransaction.(*gorm.DB), nil
}

func (b *BaseRepository) getConnection(ctx context.Context) (*gorm.DB, error) {
	var result *gorm.DB
	connWithTransaction, err := b.getTransaction(ctx)

	if err != nil {
		return nil, err
	}

	if connWithTransaction != nil {
		result = connWithTransaction
	} else {
		result = b.db
	}

	return result.WithContext(ctx), nil
}
