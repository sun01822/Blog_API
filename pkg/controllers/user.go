package controllers

import (
	"Blog_API/pkg/config"
	"Blog_API/pkg/domain"
	"Blog_API/pkg/types"
	"Blog_API/pkg/utils"
	"Blog_API/pkg/utils/consts/user"
	"Blog_API/pkg/utils/response"
	"github.com/google/uuid"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

// Parent struct to implement interface binding
type userController struct {
	svc domain.UserService
}

// Interface binding
func SetUserController(svc domain.UserService) domain.UserController {
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
func (ctr *userController) Login(ctx echo.Context) error {

	conf := config.LocalConfig

	reqUser := types.LoginRequest{}

	if err := ctx.Bind(&reqUser); err != nil {
		return response.ErrorResponse(ctx, err, userconsts.InvalidDataRequest)
	}

	if validationErr := reqUser.Validate(); validationErr != nil {
		return response.ErrorResponse(ctx, validationErr, userconsts.ValidationError)
	}

	userID, loginErr := ctr.svc.Login(reqUser.Email, reqUser.Password)
	if loginErr != nil {
		return response.ErrorResponse(ctx, loginErr, userconsts.InvalidEmailOrPassword)
	}

	now := time.Now().UTC()
	ttl := time.Minute * 15
	claims := struct {
		jwt.StandardClaims
		UserID    string `json:"user_id"`
		UserEmail string `json:"user_email"`
	}{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: now.Add(ttl).Unix(),
			IssuedAt:  now.Unix(),
			NotBefore: now.Unix(),
		},
		UserID:    userID,
		UserEmail: reqUser.Email,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, tokenErr := token.SignedString([]byte(conf.JWTSecret))
	if tokenErr != nil {
		return response.ErrorResponse(ctx, tokenErr, userconsts.ErrorGeneratingToken)
	}

	return response.SuccessResponse(ctx, userconsts.LoginSuccessful, tokenString)
}

// CreateUser implements domain.UserController.
// @Summary Create a new user
// @Description Create a new user
// @Tags user
// @Accept json
// @Produce json
// @Param user body types.SignUpRequest true "User Request"
// @Success 200 {object} types.UserResp "User created successfully"
// @Failure 400 {string} string "Invalid data request"
// @Failure 500 {string} string "Error creating user"
// @Router /user/create [post]
func (ctr *userController) CreateUser(ctx echo.Context) error {

	reqUser := types.SignUpRequest{}

	if err := ctx.Bind(&reqUser); err != nil {
		return response.ErrorResponse(ctx, err, userconsts.InvalidDataRequest)
	}

	if validationErr := reqUser.Validate(); validationErr != nil {
		return ctx.JSON(http.StatusBadRequest, validationErr.Error())
	}

	createdUser, createErr := ctr.svc.CreateUser(reqUser)
	if createErr != nil {
		return response.ErrorResponse(ctx, createErr, userconsts.ErrorCreatingUser)
	}

	return response.SuccessResponse(ctx, userconsts.UserCreatedSuccessfully, createdUser)
}

//// DeleteUser implements domain.UserController.
//func (ctr *userController) DeleteUser(c echo.Context) error {
//	tempUserId := c.Param("userID")
//	userId, err := strconv.Atoi(tempUserId)
//	if err != nil {
//		return c.JSON(http.StatusBadRequest, "Invalid user id")
//	}
//
//	existingUser, err := ctr.svc.GetUser(uint(userId))
//	if err != nil {
//		return c.JSON(http.StatusInternalServerError, err.Error())
//	}
//	if existingUser.ID == "" {
//		return c.JSON(http.StatusNotFound, "User not found")
//	}
//	if err := ctr.svc.DeleteUser(uint(userId)); err != nil {
//		return c.JSON(http.StatusInternalServerError, err.Error())
//	}
//	return c.JSON(http.StatusOK, "User deleted successfully")
//}

// GetUser implements domain.UserController.
// @Summary Get a user by ID
// @Description Get a user by ID
// @Tags user
// @Accept json
// @Produce json
// @Param user_id query string true "User ID"
// @Success 200 {object} types.UserResp "user found successfully"
// @Failure 400 {string} string "invalid data request"
// @Failure 500 {string} string "error getting user"
// @Router /user/get [get]
func (ctr *userController) GetUser(c echo.Context) error {

	reqUserID, parseErr := uuid.Parse(c.QueryParam(userconsts.UserID))
	if parseErr != nil {
		return response.ErrorResponse(c, parseErr, userconsts.InvalidDataRequest)
	}

	user, err := ctr.svc.GetUser(reqUserID.String())
	if err != nil {
		return response.ErrorResponse(c, err, userconsts.ErrorGettingUser)
	}

	return response.SuccessResponse(c, userconsts.UserFetchSuccessfully, user)
}

// GetUsers implements domain.UserController.
func (ctr *userController) GetUsers(c echo.Context) error {

	page := utils.Page{}

	pageInfo, err := page.GetPageInformation(c)
	if err != nil {
		return response.ErrorResponse(c, err, userconsts.InvalidDataRequest)
	}

	users, err := ctr.svc.GetUsers(pageInfo)
	if err != nil {
		return response.ErrorResponse(c, err, userconsts.ErrorGettingUser)
	}

	return response.SuccessResponse(c, userconsts.UsersFetchSuccessfully, users)
}

//
//// UpdateUser implements domain.UserController.
//func (ctr *userController) UpdateUser(c echo.Context) error {
//	reqUser := &types.UserUpdateRequest{}
//	if err := c.Bind(reqUser); err != nil {
//		return c.JSON(http.StatusBadRequest, "Invalid data request")
//	}
//	if err := reqUser.Validate(); err != nil {
//		return c.JSON(http.StatusBadRequest, err.Error())
//	}
//	tempUserId := c.Param("userID")
//	userId, err := strconv.Atoi(tempUserId)
//	if err != nil {
//		return c.JSON(http.StatusBadRequest, "Invalid user id")
//	}
//	existingUser, err := ctr.svc.GetUser(uint(userId))
//	if err != nil {
//		return c.JSON(http.StatusInternalServerError, err.Error())
//	}
//	user := &models.User{
//		ID:             existingUser.ID,
//		CreatedAt:      existingUser.CreatedAt,
//		UpdatedAt:      existingUser.UpdatedAt,
//		DeletedAt:      existingUser.DeletedAt,
//		Email:          existingUser.Email,
//		Gender:         reqUser.Gender,
//		DateOfBirth:    reqUser.DateOfBirth,
//		Job:            reqUser.Job,
//		City:           reqUser.City,
//		ZipCode:        reqUser.ZipCode,
//		ProfilePicture: reqUser.ProfilePicture,
//		FirstName:      reqUser.FirstName,
//		LastName:       reqUser.LastName,
//		Phone:          reqUser.Phone,
//		Street:         reqUser.Street,
//		State:          reqUser.State,
//		Country:        reqUser.Country,
//		Latitude:       reqUser.Latitude,
//		Longitude:      reqUser.Longitude,
//	}
//	user.Email = existingUser.Email
//	if user.FirstName == "" {
//		user.FirstName = existingUser.FirstName
//	}
//	if user.LastName == "" {
//		user.LastName = existingUser.LastName
//	}
//	if user.City == "" {
//		user.City = existingUser.City
//	}
//	if user.Country == "" {
//		user.Country = existingUser.Country
//	}
//	if user.DateOfBirth == nil {
//		user.DateOfBirth = existingUser.DateOfBirth
//	}
//	if user.Gender == "" {
//		user.Gender = existingUser.Gender
//	}
//	if user.Job == "" {
//		user.Job = existingUser.Job
//	}
//	if user.Latitude == 0 {
//		user.Latitude = existingUser.Latitude
//	}
//	if user.Longitude == 0 {
//		user.Longitude = existingUser.Longitude
//	}
//	if user.Phone == "" {
//		user.Phone = existingUser.Phone
//	}
//	if user.ProfilePicture == "" {
//		user.ProfilePicture = existingUser.ProfilePicture
//	}
//	if user.State == "" {
//		user.State = existingUser.State
//	}
//	if user.Street == "" {
//		user.Street = existingUser.Street
//	}
//	if user.ZipCode == "" {
//		user.ZipCode = existingUser.ZipCode
//	}
//
//	if err := ctr.svc.UpdateUser(user); err != nil {
//		return c.JSON(http.StatusInternalServerError, err.Error())
//	}
//	return c.JSON(http.StatusOK, "User updated successfully")
//}
