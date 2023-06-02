package dao

import (
	"go-backend/iternal/database/models"
	"log"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserDAO struct {
	Session *gorm.DB
}

func (u *UserDAO) Create(user *models.User) error {
	result := u.Session.Create(user)

	if result.Error != nil {
		log.Fatal("error in creating new user ", &user)
		return result.Error
	}

	return nil
}

func (u *UserDAO) GetByID(id uuid.UUID) (*models.User, error) {
	var user models.User

	result := u.Session.Model(&models.User{}).Where("ID = ?", id)

	if result.Error != nil {
		log.Fatal("error in finding user with id ", id)
		return nil, result.Error
	}

	return &user, nil
}

func (u *UserDAO) GetByEmail(email string) (*models.User, error) {
	var user models.User

	result := u.Session.Model(&models.User{}).Where("Email = ?", email)

	if result.Error != nil {
		log.Fatal("error in finding user with email ", email)
		return nil, result.Error
	}

	return &user, nil
}

func (u *UserDAO) Update(id uuid.UUID, user *models.User) (*models.User, error) {
	result := u.Session.Model(&models.User{}).Where("ID = ?", id).Updates(user)

	if result.Error != nil {
		log.Fatal("error in updating user ", user)
		return nil, result.Error
	}

	return user, nil
}

func (u *UserDAO) Delete(id uuid.UUID) error {
	result := u.Session.Model(&models.User{}).Delete("ID = ?", id)

	if result.Error != nil {
		log.Fatal("error in deleting user with id", id)
		return result.Error
	}

	return nil
}
