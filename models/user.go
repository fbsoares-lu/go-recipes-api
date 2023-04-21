package models

import (
	"gopkg.in/validator.v2"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     string `json:"name" validate:"nonzero"`
	Email    string `json:"email" validate:"nonzero" gorm:"unique"`
	Password string `json:"password" validate:"nonzero"`
	Role     string `json:"role" validate:"nonzero"`
	Recipes  []Recipe
}

func Validate(user *User) error {
	if err := validator.Validate(user); err != nil {
		return err
	}
	return nil
}
