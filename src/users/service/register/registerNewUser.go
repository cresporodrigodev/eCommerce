package register

import (
	"github.com/cresporodrigodev/ecommerce/src/users/domain/entity"
	"github.com/cresporodrigodev/ecommerce/src/users/domain/repository"
	V0 "github.com/cresporodrigodev/ecommerce/src/users/domain/valueObject"
)

type RegisterNewUser struct {
	alias      string
	firstName  string
	secondName string
	email      string
	admin      string
	password   string
	repository repository.UserRepository
}

type RegisterNewUserConstructorData struct {
	Alias      string
	FirstName  string
	SecondName string
	Email      string
	Admin      string
	Password   string
	Repository repository.UserRepository
}

func NewRegisterNewUser(data RegisterNewUserConstructorData) (*RegisterNewUser, error) {

	return &RegisterNewUser{
		alias:      data.Alias,
		firstName:  data.FirstName,
		secondName: data.SecondName,
		email:      data.Email,
		admin:      data.Admin,
		password:   data.Password,
		repository: data.Repository,
	}, nil
}

func (r *RegisterNewUser) RegisterNewUser() error {
	id, err := V0.NewUserId()
	if err != nil {
		return err
	}

	alias, err := V0.NewUserAlias(r.alias)
	if err != nil {
		return err
	}

	firstName, err := V0.NewUserFirstName(r.firstName)
	if err != nil {
		return err
	}

	secondName, err := V0.NewUserSecondName(r.secondName)
	if err != nil {
		return err
	}

	admin, err := V0.NewUserAdmin(r.admin)

	email, err := V0.NewUserEmail(r.email)
	if err != nil {
		return err
	}

	password, err := V0.NewUserPassword(r.email)
	if err != nil {
		return err
	}

	userData := entity.UserConstructorData{
		Id:         id,
		Alias:      alias,
		FirstName:  firstName,
		SecondName: secondName,
		Email:      email,
		Admin:      admin,
		Password:   password,
		CreateAt:   V0.NewCreateAt(),
	}

	user, err := entity.NewUser(userData)
	if err != nil {
		return err
	}

	err = r.repository.Save(user)
	if err != nil {
		return err
	}

	return nil
}
