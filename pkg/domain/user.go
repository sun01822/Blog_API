package domain

import (
	"Blog_API/pkg/models"
	"Blog_API/pkg/utils"
	"github.com/labstack/echo/v4"
)

// For database Repository opearation (call from service)
type UserRepository interface {
	LoginRepo(email string, password string) error
	CreateUserRepo(user *models.User) error
	GetUserRepo(id uint) (models.User, error)
	GetUsersRepo(pagination *utils.Page) ([]models.User, error)
	UpdateUserRepo(user *models.User) error
	DeleteUserRepo(id uint) error
}

// For service operation (call from controller)
type UserService interface {
	Login(email string, password string) error
	CreateUser(user *models.User) error
	GetUser(id uint) (models.User, error)
	GetUsers(pagination *utils.Page) ([]models.User, error)
	UpdateUser(user *models.User) error
	DeleteUser(id uint) error
}

// For controller operation (call from main)
type UserController interface {
	Login(c echo.Context) error
	CreateUser(c echo.Context) error
	GetUser(c echo.Context) error
	GetUsers(c echo.Context) error
	UpdateUser(c echo.Context) error
	DeleteUser(c echo.Context) error
}
