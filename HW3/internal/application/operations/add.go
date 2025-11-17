package operations

import (
	"errors"
	"hw3/internal/application/interfaces"
)

func Add(id uint64, value int64, repo interfaces.Repository) error {
	acc, err := repo.GetById(id)
	if err != nil {
		return err
	}
	if value < 0 && acc.Balance+value < 0 {
		return errors.New("invalid operation")
	}
	acc.Balance += value
	err = repo.Update(acc)
	if err != nil {
		return err
	}
	return nil
}
