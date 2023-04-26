package models

import "gorm.io/gorm"

type Profile struct {
	gorm.Model
	Rg        string
	Cpf       string
	Birthday  string
	Biography string
	UserID    int
	User      User
	File      File
}
