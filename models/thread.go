package models

type Thread struct {
	Title       string
	Description string
	RecipeID    int
	Recipe      Recipe
	UserID      int
	User        User
}
