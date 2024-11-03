package services

import (
	"Blog_API/pkg/domain"
	"Blog_API/pkg/models"
	"Blog_API/pkg/types"
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

//// DeleteUser implements domain.UserService.
//func (svc *userService) DeleteUser(id uint) error {
//	if err := svc.repo.DeleteUserRepo(id); err != nil {
//		return err
//	}
//	return nil
//}

// GetUser implements domain.UserService.
func (svc *userService) GetUser(userID string) (types.UserResp, error) {

	user, err := svc.repo.GetUser(userID)
	if err != nil {
		return types.UserResp{}, err
	}

	return convertUserToUserResp(user), nil
}

//
//// GetUsers implements domain.UserService.
//func (svc *userService) GetUsers(pagination *utils.Page) ([]models.User, error) {
//	users, err := svc.repo.GetUsersRepo(pagination)
//	if err != nil {
//		return users, err
//	}
//	return users, nil
//}
//
//// UpdateUser implements domain.UserService.
//func (svc *userService) UpdateUser(user *models.User) error {
//	if err := svc.repo.UpdateUserRepo(user); err != nil {
//		return err
//	}
//	return nil
//}

func convertUserToUserResp(user models.User) types.UserResp {
	return types.UserResp{
		ID:          user.ID,
		Email:       user.Email,
		Gender:      user.Gender,
		DateOfBirth: user.DateOfBirth,
		Phone:       user.Phone,
		Country:     user.Country,
	}
}
