package controllers

import (
	"Blog_API/pkg/config"
	"Blog_API/pkg/domain"
	"Blog_API/pkg/types"
	"Blog_API/pkg/utils"
	"Blog_API/pkg/utils/consts"
	"Blog_API/pkg/utils/consts/user"
	"Blog_API/pkg/utils/response"
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/nfnt/resize"
	"image"
	"image/jpeg"
	"image/png"
	"mime/multipart"
	"os"
	"time"
)

type userController struct {
	svc domain.Service
}

func SetUserController(svc domain.Service) domain.Controller {
	return &userController{
		svc: svc,
	}
}

// Login godoc
// @Summary User login
// @Description Logs in a user and returns a JWT token
// @Tags User
// @Accept json
// @Produce json
// @Param login body types.LoginRequest true "Login Request"
// @Success 200 {string} string "JWT Token"
// @Failure 400 {string} string "invalid data request"
// @Failure 401 {string} string "invalid email or password"
// @Router /user/login [post]
// Login implements domain.Controller.
func (ctr *userController) Login(ctx echo.Context) error {

	reqUser := types.LoginRequest{}

	if err := ctx.Bind(&reqUser); err != nil {
		return response.ErrorResponse(ctx, err, consts.InvalidDataRequest)
	}

	if validationErr := reqUser.Validate(); validationErr != nil {
		return response.ErrorResponse(ctx, validationErr, consts.ValidationError)
	}

	userID, loginErr := ctr.svc.Login(reqUser.Email, reqUser.Password)
	if loginErr != nil {
		return response.ErrorResponse(ctx, loginErr, userconsts.InvalidEmailOrPassword)
	}

	tokenString, tokenErr := generateToken(userID, reqUser.Email)
	if tokenErr != nil {
		return response.ErrorResponse(ctx, tokenErr, userconsts.ErrorGeneratingToken)
	}

	return response.SuccessResponse(ctx, userconsts.LoginSuccessful, tokenString)
}

// Logout godoc
// @Summary User logout
// @Description Logs out a user
// @Tags User
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param Authorization header string true "Bearer <token>"
// @Success 200 {string} string "user logged out successfully"
// @Failure 400 {string} string "invalid data request"
// @Failure 500 {string} string "error getting user"
// @Router /user/logout [post]
// Logout implements domain.Controller.
func (ctr *userController) Logout(ctx echo.Context) error {

	userID, parseErr := uuid.Parse(ctx.Get(userconsts.UserID).(string))
	if parseErr != nil {
		return response.ErrorResponse(ctx, parseErr, consts.InvalidDataRequest)
	}

	user, err := ctr.svc.GetUser(userID.String())
	if err != nil {
		return response.ErrorResponse(ctx, err, userconsts.ErrorGettingUser)
	}

	if user.ID == "" {
		return response.ErrorResponse(ctx, errors.New(userconsts.LogoutFailed), userconsts.UserNotFound)
	}

	return response.SuccessResponse(ctx, userconsts.LogoutSuccessful, user.Email)
}

// CreateUser implements domain.Controller.
// @Summary Create a new user
// @Description Create a new user
// @Tags User
// @Accept json
// @Produce json
// @Param user body types.SignUpRequest true "User Request"
// @Success 200 {object} types.UserResp "user created successfully"
// @Failure 400 {string} string "invalid data request"
// @Failure 500 {string} string "error creating user"
// @Router /user/create [post]
func (ctr *userController) CreateUser(ctx echo.Context) error {

	reqUser := types.SignUpRequest{}

	if err := ctx.Bind(&reqUser); err != nil {
		return response.ErrorResponse(ctx, err, consts.InvalidDataRequest)
	}

	if validationErr := reqUser.Validate(); validationErr != nil {
		return response.ErrorResponse(ctx, validationErr, consts.ValidationError)
	}

	createdUser, createErr := ctr.svc.CreateUser(reqUser)
	if createErr != nil {
		return response.ErrorResponse(ctx, createErr, userconsts.ErrorCreatingUser)
	}

	return response.SuccessResponse(ctx, userconsts.UserCreatedSuccessfully, createdUser)
}

// GetUser implements domain.Controller.
// @Summary Get a user by ID
// @Description Get a user by ID
// @Tags User
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
		return response.ErrorResponse(c, parseErr, consts.InvalidDataRequest)
	}

	user, err := ctr.svc.GetUser(reqUserID.String())
	if err != nil {
		return response.ErrorResponse(c, err, userconsts.ErrorGettingUser)
	}

	return response.SuccessResponse(c, userconsts.UserFetchSuccessfully, user)
}

// GetUsers implements domain.Controller.
// @Summary Get all users
// @Description Get all users
// @Tags User
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
		return response.ErrorResponse(c, err, consts.InvalidDataRequest)
	}

	users, err := ctr.svc.GetUsers(pageInfo)
	if err != nil {
		return response.ErrorResponse(c, err, userconsts.ErrorGettingUsers)
	}

	return response.SuccessResponse(c, userconsts.UsersFetchSuccessfully, users)
}

// UpdateUser implements domain.Controller.
// @Summary Update a user
// @Description Update a user
// @Tags User
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param Authorization header string true "Bearer <token>"
// @Param user body types.UserUpdateRequest true "User Request"
// @Success 200 {object} types.UserResp "user updated successfully"
// @Failure 400 {string} string "invalid data request"
// @Failure 500 {string} string "error updating user"
// @Router /user/update [put]
func (ctr *userController) UpdateUser(c echo.Context) error {

	userID, parseErr := uuid.Parse(c.Get(userconsts.UserID).(string))
	if parseErr != nil {
		return response.ErrorResponse(c, parseErr, consts.InvalidDataRequest)
	}

	profilePic, err := c.FormFile(userconsts.ProfilePicture)
	if err != nil {
		return response.ErrorResponse(c, err, consts.InvalidDataRequest)
	}

	profilePicURL, err := uploadProfilePicture(profilePic)
	if err != nil {
		return response.ErrorResponse(c, err, consts.InvalidDataRequest)
	}

	reqUser := types.UserUpdateRequest{}
	if err := c.Bind(&reqUser); err != nil {
		return response.ErrorResponse(c, err, consts.InvalidDataRequest)
	}

	reqUser.ProfilePicture = profilePicURL

	if validationErr := reqUser.Validate(); validationErr != nil {
		return response.ErrorResponse(c, validationErr, consts.ValidationError)
	}

	user, err := ctr.svc.UpdateUser(userID.String(), reqUser)
	if err != nil {
		return response.ErrorResponse(c, err, userconsts.ErrorUpdatingUser)
	}

	return response.SuccessResponse(c, userconsts.UserUpdatedSuccessfully, user)
}

// DeleteUser implements domain.Controller.
// @Summary Delete a user
// @Description Delete a user
// @Tags User
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param Authorization header string true "Bearer <token>"
// @Success 200 {string} string "user deleted successfully"
// @Failure 400 {string} string "invalid data request"
// @Failure 500 {string} string "error deleting user"
// @Router /user/delete [delete]
func (ctr *userController) DeleteUser(c echo.Context) error {

	userID, parseErr := uuid.Parse(c.Get(userconsts.UserID).(string))
	if parseErr != nil {
		return response.ErrorResponse(c, parseErr, consts.InvalidDataRequest)
	}

	user, err := ctr.svc.DeleteUser(userID.String())
	if err != nil {
		return response.ErrorResponse(c, err, userconsts.ErrorDeletingUser)
	}

	return response.SuccessResponse(c, userconsts.UserDeletedSuccessfully, user)
}

func generateToken(userID string, userEmail string) (string, error) {

	conf := config.LocalConfig

	now := time.Now().UTC()
	ttl := time.Minute * consts.ExpiredTokenLimit

	jwtClaims := types.JWTClaims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: now.Add(ttl).Unix(),
			IssuedAt:  now.Unix(),
			NotBefore: now.Unix(),
		},
		UserID:    userID,
		UserEmail: userEmail,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwtClaims)

	tokenString, tokenErr := token.SignedString([]byte(conf.JWTSecret))
	if tokenErr != nil {
		return "", tokenErr
	}

	return tokenString, nil
}

func uploadProfilePicture(profilePic *multipart.FileHeader) (string, error) {
	// Open and decode the image
	profilePicture, format, err := openAndDecodeImage(profilePic)
	if err != nil {
		return "", err
	}

	// Resize the image to 600x600
	tempFilePath, err := resizeAndSaveImage(profilePicture, format, 600, 600)
	if err != nil {
		return "", err
	}
	defer os.Remove(tempFilePath) // Clean up the temp file after use

	// Upload the resized image to Kraken
	krakedURL, err := uploadToKraken(tempFilePath)
	if err != nil {
		return "", err
	}

	return krakedURL, nil
}

// openAndDecodeImage opens the uploaded file and decodes it into an image.Image
func openAndDecodeImage(profilePic *multipart.FileHeader) (image.Image, string, error) {
	// Open the uploaded file
	src, err := profilePic.Open()
	if err != nil {
		return nil, "", fmt.Errorf("failed to open uploaded file: %v", err)
	}
	defer src.Close()

	// Decode the image
	profilePicture, format, err := image.Decode(src)
	if err != nil {
		return nil, "", fmt.Errorf("failed to decode image: %v", err)
	}

	return profilePicture, format, nil
}

// resizeAndSaveImage resizes an image to the specified dimensions and saves it to a temporary file
func resizeAndSaveImage(img image.Image, format string, width, height uint) (string, error) {
	// Resize the image
	resizedImg := resize.Resize(width, height, img, resize.Lanczos3)

	// Save the resized image to a temporary file
	tempFile, err := os.CreateTemp("", "resized-*.jpg")
	if err != nil {
		return "", fmt.Errorf("failed to create temporary file: %v", err)
	}

	// Encode the resized image
	switch format {
	case "jpeg", "jpg":
		err = jpeg.Encode(tempFile, resizedImg, nil)
	case "png":
		err = png.Encode(tempFile, resizedImg)
	default:
		return "", fmt.Errorf("unsupported image format: %v", format)
	}

	if err != nil {
		return "", fmt.Errorf("failed to save resized image: %v", err)
	}

	// Close the temporary file
	if err := tempFile.Close(); err != nil {
		return "", fmt.Errorf("failed to close temporary file: %v", err)
	}

	return tempFile.Name(), nil
}

// uploadToKraken uploads the file to Kraken.io and returns the optimized image URL
func uploadToKraken(filePath string) (string, error) {
	kr, err := config.Kraken()
	if err != nil {
		return "", fmt.Errorf("failed to configure Kraken: %v", err)
	}

	// Prepare parameters for Kraken upload
	params := map[string]interface{}{
		"wait": true,
	}

	// Upload to Kraken
	data, err := kr.Upload(params, filePath)
	if err != nil {
		return "", fmt.Errorf("failed to upload to Kraken: %v", err)
	}

	// Check upload success
	if success, ok := data["success"].(bool); !ok || !success {
		return "", fmt.Errorf("failed to upload: %v", data["message"])
	}

	// Return the optimized image URL
	krakedURL, _ := data["kraked_url"].(string)
	return krakedURL, nil
}
