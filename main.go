package main

import (
	"fbsoares-lu/go-recipes-api/database"
	"fbsoares-lu/go-recipes-api/helpers"
	"fbsoares-lu/go-recipes-api/routes"
)

func init() {
	helpers.LoadVariables()
	database.Connection()
	database.SyncMigrations()
}

func main() {
	routes.HandleRequests()
}
