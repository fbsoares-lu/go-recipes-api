package models

type Favorite struct {
	UserID   int
	User     User
	RecipeID int
	Recipe   Recipe
}
