package controllers

import (
	"Blog_API/pkg/config"
	"Blog_API/pkg/domain"
	"Blog_API/pkg/models"
	"Blog_API/pkg/types"
	"Blog_API/pkg/utils"
	"net/http"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

// Parent struct to implement interface binding
type userController struct {
	svc domain.UserService
}

// Interface binding
func NewUserController(svc domain.UserService) domain.UserController {
	return &userController{
		svc: svc,
	}
}


// Login implements domain.UserController.
func (ctr *userController) Login(c echo.Context) error {
	config := config.LocalConfig
	reqUser := &types.LoginRequest{}
	if err := c.Bind(reqUser); err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid data request")
	}
	if err := reqUser.Validate(); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	if err := ctr.svc.Login(reqUser.Email, reqUser.Password); err != nil {
		return c.JSON(http.StatusUnauthorized, "Invalid email or password")
	}
	now := time.Now().UTC()
	ttl := time.Minute * 15
	claims := jwt.StandardClaims{
		ExpiresAt: now.Add(ttl).Unix(),
		IssuedAt:  now.Unix(),
		NotBefore: now.Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(config.JWTSecret))
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, tokenString)
}


// CreateUser implements domain.UserController.
func (ctr *userController) CreateUser(c echo.Context) error {
	reqUser := &types.SignUpRequest{}
	if err := c.Bind(reqUser); err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid data request")
	}
	if err := reqUser.Validate(); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	user := &models.User{
		Email 			: reqUser.Email,
		Password 		: reqUser.Password,
		Gender  		: reqUser.Gender,  
		DateOfBirth 	: reqUser.DateOfBirth,    
		Job				: reqUser.Job,            
		City        	: reqUser.City,  
		ZipCode     	: reqUser.ZipCode,   
		ProfilePicture 	: reqUser.ProfilePicture,
		FirstName       : reqUser.FirstName,      
		LastName        : reqUser.LastName,    
		Phone           : reqUser.Phone, 
		Street       	: reqUser.Street, 
		State           : reqUser.State,
		Country         : reqUser.Country,
		Latitude      	: reqUser.Latitude,
		Longitude     	: reqUser.Longitude,
	}
	if err := ctr.svc.CreateUser(user); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusCreated, "User created successfully")
}


// DeleteUser implements domain.UserController.
func (ctr *userController) DeleteUser(c echo.Context) error {
	tempUserId := c.Param("id")
	userId, err := strconv.Atoi(tempUserId)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid user id")
	}

	existingUser, err := ctr.svc.GetUser(userId)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	if existingUser.ID == 0 {
		return c.JSON(http.StatusNotFound, "User not found")
	}
	if err := ctr.svc.DeleteUser(userId); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, "User deleted successfully")
}

// GetUser implements domain.UserController.
func (ctr *userController) GetUser(c echo.Context) error {
	tempUserId := c.Param("id")
	userId, err := strconv.Atoi(tempUserId)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid user id")
	}
	user, err := ctr.svc.GetUser(userId)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, user)
}

// GetUsers implements domain.UserController.
func (ctr *userController) GetUsers(c echo.Context) error {
	users, err := ctr.svc.GetUsers()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, users)
}

// UpdateUser implements domain.UserController.
func (ctr *userController) UpdateUser(c echo.Context) error {
	reqUser := &types.UserUpdateRequest{}
	if err := c.Bind(reqUser); err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid data request")
	}
	if err := reqUser.Validate(); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	if reqUser.Email != "" {
		return c.JSON(http.StatusBadRequest, "Email cannot be updated")
	}
	tempUserId := c.Param("id")
	userId, err := strconv.Atoi(tempUserId)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid user id")
	}
	existingUser, err := ctr.svc.GetUser(userId)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	user := &models.User{
		Model: gorm.Model{ID: uint(userId), CreatedAt: existingUser.CreatedAt, UpdatedAt: existingUser.UpdatedAt, DeletedAt: existingUser.DeletedAt},
		Password 		: reqUser.Password,
		Gender  		: reqUser.Gender,  
		DateOfBirth 	: reqUser.DateOfBirth,    
		Job				: reqUser.Job,            
		City        	: reqUser.City,  
		ZipCode     	: reqUser.ZipCode,   
		ProfilePicture 	: reqUser.ProfilePicture,
		FirstName       : reqUser.FirstName,      
		LastName        : reqUser.LastName,    
		Phone           : reqUser.Phone, 
		Street       	: reqUser.Street, 
		State           : reqUser.State,
		Country         : reqUser.Country,
		Latitude      	: reqUser.Latitude,
		Longitude     	: reqUser.Longitude,
	}
	user.Email = existingUser.Email
	if existingUser.Password != "" {
		user.Password = utils.HashPassword(reqUser.Password)
	}
	if existingUser.Password == "" {
		user.Password = existingUser.Password
	}
	if existingUser.Gender == "" {
		user.Gender = existingUser.Gender
	}
	if existingUser.DateOfBirth.IsZero(){
		user.DateOfBirth = existingUser.DateOfBirth
	}
	if existingUser.Job == "" {
		user.Job = existingUser.Job
	}
	if existingUser.City == "" {
		user.City = existingUser.City
	}
	if existingUser.ZipCode == "" {
		user.ZipCode = existingUser.ZipCode
	}
	if existingUser.ProfilePicture == "" {
		user.ProfilePicture = existingUser.ProfilePicture
	}
	if existingUser.FirstName == "" {
		user.FirstName = existingUser.FirstName
	}
	if existingUser.LastName == "" {
		user.LastName = existingUser.LastName
	}
	if existingUser.Phone == "" {
		user.Phone = existingUser.Phone
	}
	if existingUser.Street == "" {
		user.Street = existingUser.Street
	}
	if existingUser.State == "" {
		user.State = existingUser.State
	}
	if existingUser.Country == "" {
		user.Country = existingUser.Country
	}
	if existingUser.Latitude == 0 {
		user.Latitude = existingUser.Latitude
	}
	if existingUser.Longitude == 0 {
		user.Longitude = existingUser.Longitude
	}
	

	if err := ctr.svc.UpdateUser(user); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, "User updated successfully")
}
