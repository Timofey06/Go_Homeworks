package operations

import (
	"hw3/internal/application/interfaces"
)

func Balance(id uint64, repo interfaces.Repository) (int64, error) {
	acc, err := repo.GetById(id)
	if err != nil {
		return 0, err
	}
	return acc.Balance, nil
}
