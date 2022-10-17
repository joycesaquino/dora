package service

import (
	"dora/internal/entity"
	"dora/internal/repository"
	"dora/internal/types"
)

type GithubService struct {
	repository *repository.DoraRepository
}

func NewGithubService() *GithubService {

	doraRepository, err := repository.NewDoraRepository()
	if err != nil {
		return nil
	}
	return &GithubService{repository: doraRepository}
}

func (s GithubService) Create(event types.PullRequestEvent) (*entity.PR, error) {

	pr, err := s.repository.Create(event.Converter())
	if err != nil {
		return nil, err
	}
	return pr, nil

}
