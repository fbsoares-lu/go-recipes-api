package services

import (
	"errors"
	"fbsoares-lu/go-recipes-api/models"
	"fbsoares-lu/go-recipes-api/repositories"
)

func FindOneProfileService(id int) (models.Profile, error) {
	profile, err := repositories.FindByIdProfile(id)

	if profile.ID == 0 {
		return profile, errors.New("profile not found")
	}

	return profile, err
}
