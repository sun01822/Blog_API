package domain

import (
	"Blog_API/pkg/models"
	"Blog_API/pkg/types"
	"Blog_API/pkg/utils"
	"github.com/labstack/echo/v4"
)

// For database Repository opearation (call from service)
type UserRepository interface {
	Login(email string, password string) (string, error)
	CreateUser(user models.User) error
	GetUser(userID string) (models.User, error)
	GetUsers(pagination utils.Page) ([]models.User, error)
	UpdateUser(user models.User) error
	DeleteUser(userID string) error
}

// For service operation (call from controller)
type UserService interface {
	Login(email string, password string) (string, error)
	CreateUser(user types.SignUpRequest) (types.UserResp, error)
	GetUser(userID string) (types.UserResp, error)
	GetUsers(pagination utils.Page) ([]types.UserResp, error)
	UpdateUser(userID string, user types.UserUpdateRequest) (types.UserResp, error)
	DeleteUser(userID string) (string, error)
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
