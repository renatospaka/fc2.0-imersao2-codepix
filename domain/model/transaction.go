package model

import (
	"errors"
	"time"

	"github.com/asaskevich/govalidator"
	uuid "github.com/satori/go.uuid"
)

const (
	TransactionPendind   string = "pending"
	TransactionComplete  string = "complete"
	TransactionError     string = "error"
	TransactionConfirmed string = "confirmed"
)

type TransactionRepositoryInterface interface {
	Register(transaction *Transaction) error
	Save(transaction *Transaction) error
	Find(id string) (*Transaction, error)
}

type Transactions struct {
	Transaction []Transaction
}

type Transaction struct {
	Base              `valid:"required"`
	AccountFrom       *Account `valid:"-"`
	AccountFromID     string   `gorm:"column:account_from_id;type:uuid" valid:"notnull"`
	PixKeyTo          *PixKey  `valid:"-"`
	PixIdTo           string   `gorm:"column:pix_key_id_to;type:uuid" valid:"notnull"`
	Amount            float64  `json:"amount" gorm:"type:float" valid:"notnull"`
	Status            string   `json:"status" gorm:"type:varchar(20)" valid:"notnull"`
	Description       string   `json:"description" gorm:"type:varchar(255)" valid:"notnull"`
	CancelDescription string   `json:"cancel_description" gorm:"type:varchar(255)" valid:"-"`
}

func (t *Transaction) isValid() error {
	_, err := govalidator.ValidateStruct(t)

	if t.Amount <= 0 {
		return errors.New("Amount must be greater than zero.")
	}

	if t.Status != TransactionComplete &&
		t.Status != TransactionPendind &&
		t.Status != TransactionError {
		return errors.New("Invalid status for transaction.")
	}

	if t.PixKeyTo.AccountID == t.AccountFrom.ID {
		return errors.New("Source and destination accounts cannot be the same.")
	}

	if err != nil {
		return err
	}
	return nil
}

func NewTransaction(accountFrom *Account, amount float64, pixKey *PixKey, description string) (*Transaction, error) {
	t := Transaction{
		AccountFrom: accountFrom,
		PixKeyTo:    pixKey,
		Amount:      amount,
		Status:      TransactionPendind,
		Description: description,
	}
	t.ID = uuid.NewV4().String()
	t.CreatedAt = time.Now()

	err := t.isValid()
	if err != nil {
		return nil, err
	}

	return &t, nil
}

func (t *Transaction) Complete() error {
	t.Status = TransactionComplete
	t.UpdatedAt = time.Now()

	err := t.isValid()
	return err
}

func (t *Transaction) Cancel(description string) error {
	t.Status = TransactionError
	t.UpdatedAt = time.Now()
	t.CancelDescription = description

	err := t.isValid()
	return err
}

func (t *Transaction) Confirm() error {
	t.Status = TransactionConfirmed
	t.UpdatedAt = time.Now()

	err := t.isValid()
	return err
}
