package valueObject

import (
	"errors"
	"fmt"
	"strconv"
)

type UserAdmin struct {
	admin bool
}

func NewUserAdmin(adminString string) (UserAdmin, error) {
	adminError := errors.New("invalid value for admin atribute")
	adminBool, err := strconv.ParseBool(adminString)

	if err != nil {
		return UserAdmin{}, fmt.Errorf("%w: %s", adminError, adminString)
	}

	return UserAdmin{admin: adminBool}, nil
}

func (u UserAdmin) String() string {

	return strconv.FormatBool(u.admin)
}
