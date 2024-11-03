package services

import (
	"Blog_API/pkg/domain"
	"Blog_API/pkg/models"
	"Blog_API/pkg/types"
	"Blog_API/pkg/utils"
	userconsts "Blog_API/pkg/utils/consts/user"
	"errors"
	"github.com/google/uuid"
)

// Parent struct to implement interface binding
type userService struct {
	repo domain.UserRepository
}

// Interface binding
func SetUserService(repo domain.UserRepository) domain.UserService {
	return &userService{
		repo: repo,
	}
}

// Login implements domain.UserService.
func (svc *userService) Login(email string, password string) (string, error) {

	userID, err := svc.repo.Login(email, password)
	if err != nil {
		return "", errors.New(userconsts.LoginFailed)
	}

	return userID, nil
}

// CreateUser implements domain.UserService.
func (svc *userService) CreateUser(reqUser types.SignUpRequest) (types.UserResp, error) {

	user := models.User{
		ID:          uuid.NewString(),
		Email:       reqUser.Email,
		Password:    reqUser.Password,
		Gender:      reqUser.Gender,
		DateOfBirth: reqUser.DateOfBirth,
		Phone:       reqUser.Phone,
		Country:     reqUser.Country,
	}
	if err := svc.repo.CreateUser(user); err != nil {
		return types.UserResp{}, err
	}

	return convertUserToUserResp(user), nil
}

// DeleteUser implements domain.UserService.
func (svc *userService) DeleteUser(userID string) (string, error) {

	existingUser, err := svc.repo.GetUser(userID)
	if err != nil {
		return "", err
	}

	if deleteErr := svc.repo.DeleteUser(existingUser.ID); deleteErr != nil {
		return "", deleteErr
	}

	return existingUser.Email, nil
}

// GetUser implements domain.UserService.
func (svc *userService) GetUser(userID string) (types.UserResp, error) {

	user, err := svc.repo.GetUser(userID)
	if err != nil {
		return types.UserResp{}, err
	}

	return convertUserToUserResp(user), nil
}

// GetUsers implements domain.UserService.
func (svc *userService) GetUsers(pagination utils.Page) ([]types.UserResp, error) {

	var usersResp []types.UserResp

	users, err := svc.repo.GetUsers(pagination)
	if err != nil {
		return usersResp, err
	}

	for _, user := range users {
		usersResp = append(usersResp, convertUserToUserResp(user))
	}

	return usersResp, nil
}

// UpdateUser implements domain.UserService.
func (svc *userService) UpdateUser(userID string, userReq types.UserUpdateRequest) (types.UserResp, error) {

	user, userErr := svc.repo.GetUser(userID)
	if userErr != nil {
		return types.UserResp{}, userErr
	}

	updateUser := models.User{
		ID:             user.ID,
		Email:          user.Email,
		Password:       utils.HashPassword(userReq.Password),
		FirstName:      userReq.FirstName,
		LastName:       userReq.LastName,
		Gender:         userReq.Gender,
		DateOfBirth:    userReq.DateOfBirth,
		Job:            userReq.Job,
		Phone:          userReq.Phone,
		Street:         userReq.Street,
		City:           userReq.City,
		ZipCode:        userReq.ZipCode,
		State:          userReq.State,
		Country:        userReq.Country,
		Latitude:       userReq.Latitude,
		Longitude:      userReq.Longitude,
		ProfilePicture: userReq.ProfilePicture,
	}

	if err := svc.repo.UpdateUser(updateUser); err != nil {
		return types.UserResp{}, err
	}

	return convertUserToUserResp(updateUser), nil
}

func convertUserToUserResp(user models.User) types.UserResp {
	return types.UserResp{
		Email:          user.Email,
		FirstName:      user.FirstName,
		LastName:       user.LastName,
		Gender:         user.Gender,
		DateOfBirth:    user.DateOfBirth,
		Job:            user.Job,
		Phone:          user.Phone,
		Street:         user.Street,
		City:           user.City,
		ZipCode:        user.ZipCode,
		State:          user.State,
		Country:        user.Country,
		Latitude:       user.Latitude,
		Longitude:      user.Longitude,
		ProfilePicture: user.ProfilePicture,
	}
}
