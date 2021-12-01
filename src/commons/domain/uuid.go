package domain

import (
	"errors"
	"github.com/google/uuid"
)

type Uuid struct {
	uuid uuid.UUID
}

func NewUuid() *Uuid {
	return &Uuid{}
}

func (u *Uuid) GenerateNewUuid() error {

	var err error
	err = nil

	defer func() {
		if r := recover(); r != nil {
			err = errors.New("an error occurred while generating uuid")
		}
	}()

	u.uuid = uuid.New()

	return err
}

func (u Uuid) GetUuid() (string, error) {

	if u.uuid.String() == "" {
		return "", errors.New("uuid has not been created yet")
	}

	return u.uuid.String(), nil
}
