package database

import (
	"go-backend/iternal/database/dao"
	"go-backend/iternal/database/models"

	"github.com/google/uuid"
)

type DBFacadeInterface interface {
	CreateUser(user *models.User) error
	GeUsertByID(id uuid.UUID) (*models.User, error)
	GetUserByEmail(email string) (*models.User, error)
	UpdateUser(id uuid.UUID, user *models.User) (*models.User, error)
	DeleteUser(id uuid.UUID) error
}

type DBFacade struct {
	UserDao *dao.UserDAO
}

func (f *DBFacade) CreateUser(user *models.User) error {
	return f.UserDao.Create(user)
}

func (f *DBFacade) GeUsertByID(id uuid.UUID) (*models.User, error) {
	return f.UserDao.GetByID(id)
}

func (f *DBFacade) GetUserByEmail(email string) (*models.User, error) {
	return f.UserDao.GetByEmail(email)
}

func (f *DBFacade) UpdateUser(id uuid.UUID, user *models.User) (*models.User, error) {
	return f.UserDao.Update(id, user)
}

func (f *DBFacade) DeleteUser(id uuid.UUID) error {
	return f.UserDao.Delete(id)
}
