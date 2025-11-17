package repository

import (
	"encoding/csv"
	"errors"
	"hw3/internal/domain/account"
	"os"
	"strconv"
)

type CSVRepository struct {
	path string
}

func NewCSVRepository(path string) *CSVRepository {
	return &CSVRepository{path: path}
}

func (r *CSVRepository) ReadAll() ([]account.Account, error) {
	file, err := os.Open(r.path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	rows, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	records := make([]account.Account, 0, len(rows))
	for _, row := range rows {
		id, err := strconv.ParseUint(row[0], 10, 64)
		if err != nil {
			return nil, err
		}
		bal, err := strconv.ParseInt(row[1], 10, 64)
		if err != nil {
			return nil, err
		}

		records = append(records, account.Account{Id: id, Balance: bal})
	}
	return records, nil
}

func (r *CSVRepository) WriteAll(records []account.Account) error {
	file, err := os.Create(r.path)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	for _, rec := range records {
		row := []string{strconv.FormatUint(rec.Id, 10), strconv.FormatInt(rec.Balance, 10)}
		if err := writer.Write(row); err != nil {
			return err
		}
	}

	return nil
}

func (r *CSVRepository) GetById(id uint64) (*account.Account, error) {
	records, err := r.ReadAll()
	if err != nil {
		return nil, err
	}

	for _, rec := range records {
		if rec.Id == id {
			return &rec, nil
		}
	}

	return nil, errors.New("record not found")
}

func (r *CSVRepository) Update(acc *account.Account) error {
	records, err := r.ReadAll()
	if err != nil {
		return err
	}

	updated := false
	for i := range records {
		if records[i].Id == acc.Id {
			records[i].Balance = acc.Balance
			updated = true
			break
		}
	}

	if !updated {
		return errors.New("record not found")
	}

	return r.WriteAll(records)
}

func (r *CSVRepository) UpdateTwo(acc1 *account.Account, acc2 *account.Account) error {
	records, err := r.ReadAll()
	if err != nil {
		return err
	}

	found1, found2 := false, false
	for i := range records {
		if records[i].Id == acc1.Id {
			records[i].Balance = acc1.Balance
			found1 = true
		}
		if records[i].Id == acc2.Id {
			records[i].Balance = acc2.Balance
			found2 = true
		}
	}

	if !found1 || !found2 {
		return errors.New("one or both records not found")
	}

	return r.WriteAll(records)
}
