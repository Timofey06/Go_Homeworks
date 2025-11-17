package service

import (
	"errors"
	"hw3/internal/application/interfaces"
	"hw3/internal/application/operations"
)

type AccountService struct {
	repo interfaces.Repository
}

func NewAccountService(repo interfaces.Repository) AccountService {
	return AccountService{repo: repo}
}

func (s *AccountService) Balance(id uint64) (int64, error) {
	if s.repo == nil {
		return 0, errors.New("repository not specified")
	}
	return operations.Balance(id, s.repo)
}

func (s *AccountService) Add(id uint64, value int64) error {
	if s.repo == nil {
		return errors.New("repository not specified")
	}
	return operations.Add(id, value, s.repo)
}

func (s *AccountService) Transfer(senderId uint64, targetId uint64, value int64) error {
	if s.repo == nil {
		return errors.New("repository not specified")
	}
	return operations.Transfer(senderId, targetId, value, s.repo)
}
