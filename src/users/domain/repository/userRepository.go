package repository

import "github.com/cresporodrigodev/eCommerce/src/users/domain/entity"

type UserRepository interface {
	Save(user entity.User) error
}
