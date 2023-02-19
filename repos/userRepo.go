package repos

import (
	"errors"
	"go-tinder-app/db"
	"go-tinder-app/models"
	"go-tinder-app/utils"
	"time"

	"github.com/go-playground/validator/v10"
	"golang.org/x/crypto/bcrypt"
)

func AddUser(u *models.User) error {
	defer utils.Timer(time.Now(), "AddUser")
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	u.Password = string(hashedPassword)
	dob, err := time.Parse("2006-01-02", u.DOB)
	if err != nil {
		return err
	}
	v := validator.New()
	err = v.Struct(u)
	if err != nil {
		return err
	}

	res, err := db.DB.Exec("INSERT INTO users (username, fullname, email, password, phone, dob, gender, location, profile_picture, bio, interest, created_at, is_active) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11,$12, 1)", u.Username, u.FullName, u.Email, hashedPassword, u.Phone, dob, u.Gender, u.Location, u.ProfilePicture, u.Bio, u.Interest, time.Now())
	if err != nil {
		return err
	}
	row, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if row != 1 {
		return errors.New("data terinsert = " + string(rune(row)))
	}
	return nil
}
