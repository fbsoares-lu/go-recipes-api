package services

import (
	"errors"
	"fbsoares-lu/go-recipes-api/models"
	"fbsoares-lu/go-recipes-api/repositories"
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	UserRepository repositories.UserRepository
}

func (us *UserService) FindByEmail(email string) (*models.User, error) {
	user, err := us.UserRepository.FindByEmail(email)

	if err != nil {
		return nil, err
	}

	if user.ID == 0 {
		return nil, errors.New("user not found")
	}

	return user, nil
}

func (us *UserService) CreateUserService(user models.User) error {
	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), 10)

	if err != nil {
		return err
	}

	userError := us.UserRepository.Create(user.Name, user.Email, string(hash), user.Role)

	if userError != nil {
		return userError
	}

	return nil
}

func (us *UserService) DecodePassword(hasedPassword string, password string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hasedPassword), []byte(password))

	if err != nil {
		return err
	}

	return nil
}

func (us *UserService) Authentication(email string, password string) (string, error) {
	user, _ := us.FindByEmail(email)
	us.DecodePassword(user.Password, password)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.ID,
		"exp": time.Now().Add(time.Hour * 24 * 30).Unix(),
	})

	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))

	if err != nil {
		return "", fmt.Errorf("failed to get token: %v", err)
	}

	return tokenString, nil
}
