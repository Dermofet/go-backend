package dao

import (
	"go-backend/iternal/database"
	"go-backend/iternal/database/models"
	"go-backend/iternal/schemas"
	"log"

	"github.com/google/uuid"
)

type UserDAO struct {
	DB database.Database
}

func (u *UserDAO) Create(user *models.User) error {
	result := u.DB.Session.Create(user)
	if result.Error != nil {
		log.Fatal("error in creating new user ", &user)
		return result.Error
	}
	log.Print("ggg")

	return nil
}

func (u *UserDAO) GetByID(id uuid.UUID) (*models.User, error) {
	var user models.User

	result := u.DB.Session.Model(&models.User{}).Where("ID = ?", id)

	if result.Error != nil {
		log.Fatal("error in finding user with id ", id)
		return nil, result.Error
	}

	result.First(&user)

	return &user, nil
}

func (u *UserDAO) GetByEmail(email string) (*models.User, error) {
	var user models.User

	result := u.DB.Session.Model(&models.User{}).Where("Email = ?", email)

	if result.Error != nil {
		log.Fatal("error in finding user with email ", email)
		return nil, result.Error
	}

	result.First(&user)

	return &user, nil
}

func (u *UserDAO) Update(id uuid.UUID, user *schemas.UserUpdate) (*models.User, error) {
	var userDB models.User

	result := u.DB.Session.Model(&models.User{}).Where("ID = ?", id)

	if result.Error != nil {
		log.Fatal("error in updating user ", user)
		return nil, result.Error
	}

	result.First(&userDB)

	userDB.Email = user.Email
	userDB.Username = user.Username

	result.Save(&userDB)

	return &userDB, nil
}

func (u *UserDAO) Delete(id uuid.UUID) error {
	result := u.DB.Session.Model(&models.User{}).Delete("ID = ?", id)

	if result.Error != nil {
		log.Fatal("error in deleting user with id", id)
		return result.Error
	}

	return nil
}
