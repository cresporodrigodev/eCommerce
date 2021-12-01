package valueObject

import (
	"errors"
	"fmt"
	"net/mail"
)

type UserEmail struct {
	email string
}

func NewUserEmail(email string) (UserEmail, error) {

	emailError := errors.New("email is not valid")

	if _, err := mail.ParseAddress(email); err != nil {
		return UserEmail{}, fmt.Errorf("%w: %s", emailError, email)
	}
	return UserEmail{email: email}, nil
}

func (u UserEmail) String() string {
	return u.email
}
