package controllers

import (
	"Blog_API/pkg/config"
	"Blog_API/pkg/domain"
	"Blog_API/pkg/models"
	"Blog_API/pkg/types"
	"Blog_API/pkg/utils"
	"fmt"
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

// Login godoc
// @Summary User login
// @Description Logs in a user and returns a JWT token
// @Tags user
// @Accept json
// @Produce json
// @Param login body types.LoginRequest true "Login Request"
// @Success 200 {string} string "JWT Token"
// @Failure 400 {string} string "Invalid data request"
// @Failure 401 {string} string "Invalid email or password"
// @Router /user/login [post]
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
	tempUserId := c.Param("userID")
	userId, err := strconv.Atoi(tempUserId)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid user id")
	}

	existingUser, err := ctr.svc.GetUser(uint(userId))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	if existingUser.ID == 0 {
		return c.JSON(http.StatusNotFound, "User not found")
	}
	if err := ctr.svc.DeleteUser(uint(userId)); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, "User deleted successfully")
}

func (ctr *userController) GetUser(c echo.Context) error {
	tempUserId := c.Param("userID")
	userId, err := strconv.Atoi(tempUserId)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid user id")
	}
	user, err := ctr.svc.GetUser(uint(userId))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, user)
}

// GetUsers implements domain.UserController.
func (ctr *userController) GetUsers(c echo.Context) error {
	page := &utils.Page{}
	pageInfo, err := page.GetPageInformation(c)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	fmt.Println(
		"Offset: ", *pageInfo.Offset,
		"Limit: ", *pageInfo.Limit,
	)
	users, err := ctr.svc.GetUsers(pageInfo)
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
	tempUserId := c.Param("userID")
	userId, err := strconv.Atoi(tempUserId)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid user id")
	}
	existingUser, err := ctr.svc.GetUser(uint(userId))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	user := &models.User{
		Model: gorm.Model{ID: uint(userId), CreatedAt: existingUser.CreatedAt, UpdatedAt: time.Now(), DeletedAt: existingUser.DeletedAt},
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
	if reqUser.Password != "" {
		user.Password = utils.HashPassword(reqUser.Password)
	}else if reqUser.Password == "" {
		user.Password = existingUser.Password
	}
	if user.FirstName == "" {
		user.FirstName = existingUser.FirstName
	}
	if user.LastName == "" {
		user.LastName = existingUser.LastName
	}
	if user.City == "" {
		user.City = existingUser.City
	}
	if user.Country == "" {
		user.Country = existingUser.Country
	}
	if user.DateOfBirth == nil {
		user.DateOfBirth = existingUser.DateOfBirth
	}
	if user.Gender == "" {
		user.Gender = existingUser.Gender
	}
	if user.Job == "" {
		user.Job = existingUser.Job
	}
	if user.Latitude == 0 {
		user.Latitude = existingUser.Latitude
	}
	if user.Longitude == 0 {
		user.Longitude = existingUser.Longitude
	}
	if user.Phone == "" {
		user.Phone = existingUser.Phone
	}
	if user.ProfilePicture == "" {
		user.ProfilePicture = existingUser.ProfilePicture
	}
	if user.State == "" {
		user.State = existingUser.State
	}
	if user.Street == "" {
		user.Street = existingUser.Street
	}
	if user.ZipCode == "" {
		user.ZipCode = existingUser.ZipCode
	}
	
	if err := ctr.svc.UpdateUser(user); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, "User updated successfully")
}
