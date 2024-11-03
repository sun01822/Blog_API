package controllers

import (
	"Blog_API/pkg/config"
	"Blog_API/pkg/domain"
	"Blog_API/pkg/types"
	"Blog_API/pkg/utils"
	"Blog_API/pkg/utils/consts/user"
	"Blog_API/pkg/utils/response"
	"github.com/google/uuid"
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

	jwtClaims := types.JWTClaims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: now.Add(ttl).Unix(),
			IssuedAt:  now.Unix(),
			NotBefore: now.Unix(),
		},
		UserID:    userID,
		UserEmail: reqUser.Email,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwtClaims)
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
		return response.ErrorResponse(ctx, validationErr, userconsts.ValidationError)
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
// @Summary Get all users
// @Description Get all users
// @Tags user
// @Accept json
// @Produce json
// @Param offset query string false "Offset"
// @Param limit query string false "Limit"
// @Success 200 {object} []types.UserResp "users found successfully"
// @Failure 400 {string} string "invalid data request"
// @Failure 500 {string} string "error getting user"
// @Router /user/getAll [get]
func (ctr *userController) GetUsers(c echo.Context) error {

	page := utils.Page{}

	pageInfo, err := page.GetPageInformation(c)
	if err != nil {
		return response.ErrorResponse(c, err, userconsts.InvalidDataRequest)
	}

	users, err := ctr.svc.GetUsers(pageInfo)
	if err != nil {
		return response.ErrorResponse(c, err, userconsts.ErrorGettingUsers)
	}

	return response.SuccessResponse(c, userconsts.UsersFetchSuccessfully, users)
}

// UpdateUser implements domain.UserController.
// @Summary Update a user
// @Description Update a user
// @Tags user
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param Authorization header string true "Bearer <token>"
// @Param user body types.UserUpdateRequest true "User Request"
// @Success 200 {object} types.UserResp "User updated successfully"
// @Failure 400 {string} string "Invalid data request"
// @Failure 500 {string} string "Error updating user"
// @Router /user/update [put]
func (ctr *userController) UpdateUser(c echo.Context) error {

	userID, parseErr := uuid.Parse(c.Get(userconsts.UserID).(string))
	if parseErr != nil {
		return response.ErrorResponse(c, parseErr, userconsts.InvalidDataRequest)
	}

	reqUser := types.UserUpdateRequest{}
	if err := c.Bind(&reqUser); err != nil {
		return response.ErrorResponse(c, err, userconsts.InvalidDataRequest)
	}

	if validationErr := reqUser.Validate(); validationErr != nil {
		return response.ErrorResponse(c, validationErr, userconsts.ValidationError)
	}

	user, err := ctr.svc.UpdateUser(userID.String(), reqUser)
	if err != nil {
		return response.ErrorResponse(c, err, userconsts.ErrorUpdatingUser)
	}

	return response.SuccessResponse(c, userconsts.UserUpdatedSuccessfully, user)
}
