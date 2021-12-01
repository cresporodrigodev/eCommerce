package valueObject

import (
	"errors"
	"fmt"
	"github.com/cresporodrigodev/eCommerce/src/commons/domain"
)

type UserFirstName struct {
	firstName string
}

func NewUserFirstName(firstName string) (UserFirstName, error) {
	firstNameError := errors.New("invalid user alias")
	symbolError := errors.New("user first name cannot contain special characters")

	str := domain.NewStringValidation(firstName, firstNameError)
	str.Validation()

	strValue, err := str.GetErrAndValue()
	if err != nil {
		return UserFirstName{}, err
	}

	symbolValidation := domain.NewSymbolValidation(strValue, symbolError)
	if hasSymbol := symbolValidation.HasSymbol(); hasSymbol != false {
		return UserFirstName{}, fmt.Errorf("%w: %s", symbolError, firstName)
	}

	return UserFirstName{firstName: strValue}, nil
}

func (u UserFirstName) String() string {
	return u.firstName
}

type UserSecondName struct {
	secondName string
}

func NewUserSecondName(secondName string) (UserSecondName, error) {
	secondNameError := errors.New("invalid user alias")
	symbolError := errors.New("user second name cannot contain special characters")

	str := domain.NewStringValidation(secondName, secondNameError)
	str.Validation()

	strValue, err := str.GetErrAndValue()
	if err != nil {
		return UserSecondName{}, err
	}

	symbolValidation := domain.NewSymbolValidation(strValue, symbolError)
	if hasSymbol := symbolValidation.HasSymbol(); hasSymbol != false {
		return UserSecondName{}, fmt.Errorf("%w: %s", symbolError, secondName)
	}

	return UserSecondName{secondName: strValue}, nil
}

func (u UserSecondName) String() string {
	return u.secondName
}
