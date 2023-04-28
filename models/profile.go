package models

import (
	"gopkg.in/validator.v2"
	"gorm.io/gorm"
)

type Profile struct {
	gorm.Model
	Rg        string `json:"rg" validate:"min=7,max=7" gorm:"unique"`
	Cpf       string `json:"cpf" validate:"min=11,max=11"  gorm:"unique"`
	Birthday  string `json:"birthday" validate:"nonzero"`
	Biography string `json:"biography" validate:"nonzero"`
	UserID    uint
	FileID    uint
	File      File
}

func (v *Profile) Validate(profile *Profile) error {
	if err := validator.Validate(profile); err != nil {
		return err
	}
	return nil
}
