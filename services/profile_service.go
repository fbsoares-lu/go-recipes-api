package services

import (
	"errors"
	"fbsoares-lu/go-recipes-api/models"
	"fbsoares-lu/go-recipes-api/repositories"
)

type ProfileService struct {
	ProfileRepository repositories.ProfileRepository
}

func (service *ProfileService) FindOneProfileService(id int) (*models.Profile, error) {
	profile, err := service.ProfileRepository.FindById(id)

	if profile.ID == 0 {
		return profile, errors.New("profile not found")
	}

	return profile, err
}

func (service *ProfileService) ProfileAlreadyExists(id int) (*models.Profile, error) {
	profile, err := service.ProfileRepository.FindById(id)

	if err != nil {
		return nil, err
	}

	if profile.ID == 0 {
		return nil, err
	}

	return profile, nil
}

func (service *ProfileService) CreateProfileService(profile models.Profile) error {
	err := service.ProfileRepository.Create(profile)
	return err
}

func (service *ProfileService) SaveProfileService(id int, profile models.Profile) (*models.Profile, error) {
	currentProfile, err := service.ProfileRepository.FindById(id)

	if err != nil {
		return nil, err
	}

	if currentProfile.ID == 0 {
		return nil, errors.New("profile not found")
	}

	err = service.ProfileRepository.Update(id, profile)

	if err != nil {
		return nil, err
	}

	return &profile, nil
}

func (service *ProfileService) DeleteProfileService(id int) error {
	currentProfile, err := service.ProfileRepository.FindById(id)

	if err != nil {
		return err
	}

	if currentProfile.ID == 0 {
		return errors.New("profile not found")
	}

	err = service.ProfileRepository.Delete(id)

	if err != nil {
		return err
	}

	return nil
}
