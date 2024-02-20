package models

import (
	"gorm.io/gorm"
	"time"
)

// User struct
type User struct {
	gorm.Model // Embedding the gorm.Model for ID, CreatedAt, UpdatedAt, and DeletedAt fields
	Gender      string    `json:"gender"`
	DateOfBirth *time.Time `json:"date_of_birth"`
	Job         string    `json:"job"`
	City        string    `json:"city"`
	ZipCode     string    `json:"zipcode"`
	ProfilePicture string `json:"profile_picture"`
	FirstName   string    `json:"first_name"`
	LastName    string    `json:"last_name"`
	Email       string    `json:"email"`
	Password    string    `json:"password"`
	Phone       string    `json:"phone"`
	Street      string    `json:"street"`
	State       string    `json:"state"`
	Country     string    `json:"country"`
	Latitude    float64   `json:"latitude"`
	Longitude   float64   `json:"longitude"`
}