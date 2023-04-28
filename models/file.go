package models

import "gorm.io/gorm"

type File struct {
	gorm.Model
	OriginalName string
	Recipes      []*Recipe `json:"files" gorm:"many2many:recipe_files;"`
}
