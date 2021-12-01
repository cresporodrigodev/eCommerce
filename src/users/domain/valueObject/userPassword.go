package valueObject

import (
	"errors"
	"fmt"
	"unicode"
)

type UserPassword struct {
	password string
}

func NewUserPassword(password string) (UserPassword, error) {
	const passwordMinLength = 8
	passwordErrUpp := errors.New("password must have a capital letter")
	passwordErrLow := errors.New("password must have a non capital letter")
	passwordErrNum := errors.New("password must have a number")
	passwordErrLen := fmt.Errorf("password must have a %d characters", passwordMinLength)

	var (
		upp, low, num bool
	)

	for _, char := range password {
		switch {
		case unicode.IsUpper(char):
			upp = true
		case unicode.IsLower(char):
			low = true
		case unicode.IsNumber(char):
			num = true
		default:
			continue
		}
	}

	if !upp {
		return UserPassword{}, fmt.Errorf("%w: %s", passwordErrUpp, password)
	}

	if !low {
		return UserPassword{}, fmt.Errorf("%w: %s", passwordErrLow, password)
	}

	if !num {
		return UserPassword{}, fmt.Errorf("%w: %s", passwordErrNum, password)
	}

	if len(password) < passwordMinLength {
		return UserPassword{}, fmt.Errorf("%w: %s", passwordErrLen, password)
	}

	return UserPassword{
		password: password,
	}, nil
}

func (u UserPassword) String() string {
	return u.password
}
