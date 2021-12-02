package db

import (
	"errors"
	"fmt"
	"github.com/cresporodrigodev/ecommerce/src/users/domain/entity"
)

type userData struct {
	id         string
	alias      string
	firstName  string
	secondName string
	email      string
	admin      string
	password   string
	createAt   string
	updateAt   string
}
type MockRepository struct {
	memoryDb []userData
}

func NewMockRepository() MockRepository {
	documents := []userData{
		{
			alias: "belubrin",
		},
		{
			email: "cresporodrigo42@gmail.com",
		},
	}
	return MockRepository{memoryDb: documents}
}

func (m MockRepository) Save(user entity.User) error {
	aliasErr := errors.New("alias already exists")
	emailErr := errors.New("email already exists")

	for _, value := range m.memoryDb {

		if value.alias == user.Alias() {
			return fmt.Errorf("%w: %s", aliasErr, user.Alias())
		}

		if value.email == user.Email() {
			return fmt.Errorf("%w: %s", emailErr, user.Email())
		}
	}

	return nil
}

func (m MockRepository) Connect() error {
	fmt.Println("connected to db")
	return nil
}

func (m MockRepository) Close() error {
	fmt.Println("Connection closed")
	return nil
}
