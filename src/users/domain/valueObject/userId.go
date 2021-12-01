package valueObject

import (
	"errors"
	"fmt"
	"github.com/cresporodrigodev/eCommerce/src/commons/domain"
)

type UserId struct {
	uuid string
}

func NewUserId() (UserId, error) {
	idError := errors.New("invalid user Id")

	newUuid := domain.NewUuid()
	err := newUuid.GenerateNewUuid()
	uuidValue, err := newUuid.GetUuid()

	if err != nil {
		return UserId{}, fmt.Errorf("%w: %s", idError, uuidValue)
	}

	return UserId{uuid: uuidValue}, nil

}

func (u UserId) String() string {
	return u.uuid
}
