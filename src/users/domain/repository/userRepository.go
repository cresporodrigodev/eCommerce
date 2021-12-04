package repository

import "github.com/cresporodrigodev/ecommerce/src/users/domain/entity"

type UserRepository interface {
	Save(user entity.User) error
}
