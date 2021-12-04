package db

import (
	"github.com/cresporodrigodev/ecommerce/src/commons/infrastructure/db"
	"github.com/cresporodrigodev/ecommerce/src/users/domain/entity"
)

type userModel struct {
	Id         string `bson:"id"`
	Alias      string `bson:"alias"`
	FirstName  string `bson:"first_name"`
	SecondName string `bson:"second_name"`
	Email      string `bson:"email"`
	Admin      string `bson:"admin"`
	Password   string `bson:"password"`
	CreateAt   string `bson:"create_at"`
	UpdateAt   string `bson:"update_at"`
}

type UserRepositoryMongoDb struct {
	mongoRepository *db.MongoDbRepository
}

func NewMongoDbRepository(mongoRepository *db.MongoDbRepository) UserRepositoryMongoDb {
	return UserRepositoryMongoDb{mongoRepository: mongoRepository}
}

func (m UserRepositoryMongoDb) Save(user entity.User) error {
	document := userModel{
		Id:         user.Id(),
		Alias:      user.Alias(),
		FirstName:  user.FirstName(),
		SecondName: user.SecondName(),
		Email:      user.Email(),
		Admin:      user.Admin(),
		Password:   user.Password(),
		CreateAt:   user.CreateAt(),
		UpdateAt:   user.UpdateAt(),
	}
	collections := []interface{}{document}
	err := m.mongoRepository.Insert("database", "coll", collections)
	return err
}
