package repository

import (
	"fmt"

	"github.com/jinzhu/gorm"
	"github.com/renatospaka/imersao/codepix-go/domain/model"
)

type TransactionDbRepository struct {
	Db *gorm.DB
}

func (r TransactionDbRepository) Register(transaction *model.Transaction) error {
	err := r.Db.Create(transaction).Error
	if err != nil {
		return err
	}
	return nil
}

func (r TransactionDbRepository) Save(transaction *model.Transaction) error {
	err := r.Db.Save(transaction).Error
	if err != nil {
		return err
	}
	return nil
}

func (r TransactionDbRepository) Find(id string) (*model.Transaction, error) {
	var transaction model.Transaction

	r.Db.Preload("AccountFrom.Bank").First(&transaction, "id = ?", id)
	if transaction.ID == "" {
		return nil, fmt.Errorf("No transaction found.")
	}

	return &transaction, nil
}
