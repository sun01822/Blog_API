package domain

import (
	"Blog_API/pkg/models"
	"github.com/labstack/echo/v4"
)

// For database Repository opearation (call from service)
type UserRepository interface {
	CreateUser(user *models.User) error
	GetUser(id int) (models.User, error)
	GetUsers() ([]models.User, error)
	UpdateUser(user *models.User) error
	DeleteUser(id int) error
}

// For service operation (call from controller)	
type UserService interface {
	CreateUser(user *models.User) error
	GetUser(id int) (models.User, error)
	GetUsers() ([]models.User, error)
	UpdateUser(user *models.User) error
	DeleteUser(id int) error
}

// For controller operation (call from main)
type UserController interface {
	CreateUser(c echo.Context) error
	GetUser(c echo.Context) error
	GetUsers(c echo.Context) error
	UpdateUser(c echo.Context) error
	DeleteUser(c echo.Context) error
}

