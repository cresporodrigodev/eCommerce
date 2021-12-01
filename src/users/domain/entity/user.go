package entity

import VO "github.com/cresporodrigodev/eCommerce/src/users/domain/valueObject"

type User struct {
	id         VO.UserId
	alias      VO.UserAlias
	firstName  VO.UserFirstName
	secondName VO.UserSecondName
	email      VO.UserEmail
	admin      VO.UserAdmin
	password   VO.UserPassword
	createAt   VO.CreateAt
	updateAt   VO.UpdateAt
}

type UserConstructorData struct {
	Id         VO.UserId
	Alias      VO.UserAlias
	FirstName  VO.UserFirstName
	SecondName VO.UserSecondName
	Email      VO.UserEmail
	Admin      VO.UserAdmin
	Password   VO.UserPassword
	CreateAt   VO.CreateAt
	UpdateAt   VO.UpdateAt
}

func NewUser(userData UserConstructorData) (User, error) {

	return User{
		id:         userData.Id,
		alias:      userData.Alias,
		firstName:  userData.FirstName,
		secondName: userData.SecondName,
		email:      userData.Email,
		admin:      userData.Admin,
		password:   userData.Password,
		createAt:   userData.CreateAt,
		updateAt:   userData.UpdateAt,
	}, nil
}

func (u User) Id() string {
	return u.id.String()
}

func (u User) Alias() string {
	return u.alias.String()
}

func (u User) FirstName() string {
	return u.firstName.String()
}

func (u User) SecondName() string {
	return u.secondName.String()
}

func (u User) Email() string {
	return u.email.String()
}

func (u User) Admin() string {
	return u.admin.String()
}

func (u User) Password() string {
	return u.password.String()
}

func (u User) CreateAt() string {
	return u.createAt.String()
}

func (u User) UpdateAt() string {
	return u.updateAt.String()
}
