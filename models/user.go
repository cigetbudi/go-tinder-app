package models

import "github.com/gofrs/uuid"

type User struct {
	ID             uuid.UUID `json:"id"`
	Name           string    `json:"name"`
	Email          string    `json:"email"`
	Age            int       `json:"age"`
	Gender         string    `json:"gender"`
	Location       string    `json:"location"`
	ProfilePicture string    `json:"profile_picture"`
	Bio            string    `json:"bio"`
	Interest       []string  `json:"interest"`
}
