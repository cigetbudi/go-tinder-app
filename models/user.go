package models

import (
	"github.com/gofrs/uuid"
)

type User struct {
	ID             uuid.UUID `json:"id"`
	Username       string    `json:"username" validate:"required"`
	FullName       string    `json:"fullname" validate:"required"`
	Email          string    `json:"email" validate:"required,email"`
	Password       string    `json:"password" validate:"required,min=6"`
	Phone          string    `json:"phone" validate:"required,numeric,min=11,max=15"`
	DOB            string    `json:"dob" validate:"required"`
	Gender         string    `json:"gender" validate:"required,len=1"`
	Location       string    `json:"location" validate:"required"`
	ProfilePicture string    `json:"profile_picture"`
	Bio            string    `json:"bio"`
	Interest       []string  `json:"interest"`
}
