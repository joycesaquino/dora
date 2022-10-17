package service

import (
	"dora/internal/repository"
	"dora/internal/types/dora"
	"dora/internal/types/github"
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

func (s GithubService) Create(event github.PullRequestEvent) (*dora.PR, error) {

	pr, err := s.repository.Create(event.Converter())
	if err != nil {
		return nil, err
	}
	return pr, nil

}
