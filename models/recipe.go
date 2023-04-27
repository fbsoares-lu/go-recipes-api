package models

import "gorm.io/gorm"

type Recipe struct {
	gorm.Model
	Name            string        `json:"name"`
	Description     string        `json:"description"`
	PreparationTime string        `json:"preparation_time"`
	PeopleServed    int           `json:"people_served"`
	Ingredients     []*Ingredient `json:"ingredients" gorm:"many2many:recipe_ingredients;"`
	Files           []File        `json:"files"`
	Steps           []Step        `json:"steps"`
	UserID          uint
}
