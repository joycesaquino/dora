package repository

import (
	"context"
	"dora/internal/dao"
	types "dora/internal/types/dora"
)

type DoraRepository struct {
	dao *dao.Database
}

type DoraRepositoryInterface interface {
	Create(ctx context.Context, customer *types.PR) (*types.PR, error)
}

func (repository DoraRepository) Create() {
}

func NewDoraRepository(ctx context.Context) (*DoraRepository, error) {
	db, err := dao.NewDatabase(ctx)
	if err != nil {
		return nil, err
	}

	return &DoraRepository{
		dao: db,
	}, nil
}
