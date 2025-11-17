package interfaces

import (
	"hw3/internal/domain/account"
)

type Repository interface {
	GetById(id uint64) (*account.Account, error)
	Update(acc *account.Account) error
	UpdateTwo(acc1 *account.Account, acc2 *account.Account) error
}
