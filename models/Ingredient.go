package models

import "gorm.io/gorm"

type Ingredient struct {
	gorm.Model
	Name        string
	Description string
	Recipes     []*Recipe `gorm:"many2many:recipe_ingredients;"`
	// FileRefer   uint
	// File        File `gorm:"foreignKey:FileRefer"`
}
