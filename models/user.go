package models

import (
	"time"

	"github.com/gofrs/uuid"
)

type User struct {
	ID             uuid.UUID `json:"id"`
	Username       string    `json:"username"`
	FullName       string    `json:"fullname"`
	Email          string    `json:"email"`
	Password       string    `json:"password"`
	Phone          string    `json:"phone"`
	DOB            time.Time `json:"dob"`
	Gender         string    `json:"gender"`
	Location       string    `json:"location"`
	ProfilePicture string    `json:"profile_picture"`
	Bio            string    `json:"bio"`
	Interest       []string  `json:"interest"`
}
