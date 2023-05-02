package models

import "gorm.io/gorm"

type Recipe struct {
	gorm.Model
	Name            string        `json:"name"`
	Description     string        `json:"description"`
	PreparationTime string        `json:"preparationTime"`
	PeopleServed    int           `json:"peopleServed"`
	Ingredients     []*Ingredient `json:"ingredients" gorm:"many2many:recipe_ingredients;"`
	Files           []*File       `json:"files" gorm:"many2many:recipe_files;"`
	Steps           []Step        `json:"steps"`
	UserID          uint
}
