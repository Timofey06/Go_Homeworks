package operations

import (
	"errors"
	"hw3/internal/application/interfaces"
)

func Transfer(senderId uint64, targetId uint64, value int64, repo interfaces.Repository) error {
	if value <= 0 {
		return errors.New("incorrected value")
	}
	senderAcc, err1 := repo.GetById(senderId)
	targetAcc, err2 := repo.GetById(targetId)
	if err1 != nil {
		return err1
	}
	if err2 != nil {
		return err2
	}
	if value > senderAcc.Balance {
		return errors.New("not enough funds")
	}
	if senderId == targetId {
		return errors.New("impossible to transfer to oneself")
	}

	senderAcc.Balance -= value
	targetAcc.Balance += value

	err := repo.UpdateTwo(senderAcc, targetAcc)
	if err != nil {
		return err
	}
	return nil
}
