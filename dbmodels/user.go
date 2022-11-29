package dbmodels

import (
	"golang.org/x/crypto/bcrypt"
)

func (user *User) HashPassword(password string) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return err
	}
	user.Pass = string(bytes)
	return nil
}
func (user *User) CheckPassword(providedPassword string) error {
	err := bcrypt.CompareHashAndPassword([]byte(user.Pass), []byte(providedPassword))
	if err != nil {
		return err
	}
	return nil
}
