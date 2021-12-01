package command

import (
	"github.com/cresporodrigodev/eCommerce/src/users/domain/repository"
	"github.com/cresporodrigodev/eCommerce/src/users/service/register"
)

type RegisterNewUserCommandHandler struct {
	repository repository.UserRepository
}

func NewRegisterNewUserCommandHandler(repository repository.UserRepository) RegisterNewUserCommandHandler {
	return RegisterNewUserCommandHandler{repository: repository}
}

func (r RegisterNewUserCommandHandler) Handle(command RegisterNewUserCommand) error {
	constructorData := register.RegisterNewUserConstructorData{
		Alias:      command.Alias,
		FirstName:  command.FirstName,
		SecondName: command.SecondName,
		Email:      command.Email,
		Admin:      command.Admin,
		Password:   command.Password,
		Repository: r.repository,
	}

	registerNewUserSrv, err := register.NewRegisterNewUser(constructorData)
	if err != nil {
		return err
	}

	err = registerNewUserSrv.RegisterNewUser()
	if err != nil {
		return err
	}

	return nil
}
