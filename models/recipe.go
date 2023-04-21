package models

import "gorm.io/gorm"

type Recipe struct {
	gorm.Model
	Name        string        `json:"name"`
	Description string        `json:"description"`
	Ingredients []*Ingredient `json:"ingredients" gorm:"many2many:recipe_ingredients;"`
	Files       []File        `json:"files"`
	UserID      uint
}
