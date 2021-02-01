package model

import (
	"time"

	"github.com/asaskevich/govalidator"
)

func init() {
	govalidator.SetFieldsRequiredByDefault(true)
}

type Base struct {
	ID        string    `json:"id" valid:"uuid"`
	UpdatedAt time.Time `json:"created_at" valid:"-"`
	CreatedAt time.Time `json:"updated_at" valid:"-"`
}
