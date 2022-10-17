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

func (repository DoraRepository) Create(pr *types.PR) (*types.PR, error) {
	return pr, nil
}

func NewDoraRepository() (*DoraRepository, error) {
	db, err := dao.NewDatabase()
	if err != nil {
		return nil, err
	}

	return &DoraRepository{
		dao: db,
	}, nil
}
