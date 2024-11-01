package types

import (
	validate "github.com/go-ozzo/ozzo-validation"
	"time"
)

// SignUp UserRequest
type SignUpRequest struct {
	Email       string     `json:"email"`
	Password    string     `json:"password"`
	Gender      string     `json:"gender,omitempty"`
	DateOfBirth *time.Time `json:"date_of_birth,omitempty"`
	Phone       string     `json:"phone,omitempty"`
	Country     string     `json:"country,omitempty" default:"Bangladesh"`
}

// Validate is a function that validates the request body for the user
func (user SignUpRequest) Validate() error {
	return validate.ValidateStruct(&user,
		validate.Field(&user.Email, validate.Required, validate.Length(10, 100)),
		validate.Field(&user.Password, validate.Required, validate.Length(6, 100)),
	)
}

type UserUpdateRequest struct {
	Gender         string     `json:"gender,omitempty"`
	DateOfBirth    *time.Time `json:"date_of_birth,omitempty"`
	Job            string     `json:"job,omitempty"`
	City           string     `json:"city,omitempty"`
	ZipCode        string     `json:"zipcode,omitempty"`
	ProfilePicture string     `json:"profile_picture,omitempty"`
	FirstName      string     `json:"first_name,omitempty"`
	LastName       string     `json:"last_name,omitempty"`
	Phone          string     `json:"phone,omitempty"`
	Street         string     `json:"street,omitempty"`
	State          string     `json:"state,omitempty"`
	Country        string     `json:"country,omitempty" default:"Bangladesh"`
	Latitude       float64    `json:"latitude,omitempty"`
	Longitude      float64    `json:"longitude,omitempty"`
}

func (user UserUpdateRequest) Validate() error {
	return validate.ValidateStruct(&user,
		validate.Field(user.FirstName, validate.Length(3, 100)),
		validate.Field(user.LastName, validate.Length(2, 100)),
		validate.Field(user.Phone, validate.Length(1, 100)),
	)
}

// Login UserRequest
type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// Validate is a function that validates the request body for the user
func (user LoginRequest) Validate() error {
	return validate.ValidateStruct(&user,
		validate.Field(&user.Email, validate.Required, validate.Length(10, 100)),
		validate.Field(&user.Password, validate.Required, validate.Length(6, 100)),
	)
}

// UserResponse
type UserResp struct {
	ID          string     `json:"id"`
	Email       string     `json:"email"`
	Gender      string     `json:"gender,omitempty"`
	DateOfBirth *time.Time `json:"date_of_birth,omitempty"`
	Phone       string     `json:"phone,omitempty"`
	Country     string     `json:"country,omitempty"`
}
