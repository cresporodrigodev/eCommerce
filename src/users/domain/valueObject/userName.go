package valueObject

import (
	"errors"
	"fmt"
	strValidation "github.com/cresporodrigodev/ecommerce/src/commons/domain/stringValidation"
	symValidation "github.com/cresporodrigodev/ecommerce/src/commons/domain/symbolValidation"
)

type UserFirstName struct {
	firstName string
}

func NewUserFirstName(firstName string) (UserFirstName, error) {
	firstNameError := errors.New("invalid user alias")
	symbolError := errors.New("user first name cannot contain special characters")

	strVal := strValidation.NewStringValidation(firstName)
	strValue, err := strVal.Validation()

	if err != nil {
		return UserFirstName{}, fmt.Errorf("%w: %s", firstNameError, firstName)
	}

	symbolValidation := symValidation.NewSymbolValidation(strValue)
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

	strVal := strValidation.NewStringValidation(secondName)
	strValue, err := strVal.Validation()

	if err != nil {
		return UserSecondName{}, fmt.Errorf("%w: %s", secondNameError, secondName)
	}

	symbolValidation := symValidation.NewSymbolValidation(strValue)
	if hasSymbol := symbolValidation.HasSymbol(); hasSymbol != false {
		return UserSecondName{}, fmt.Errorf("%w: %s", symbolError, secondName)
	}

	return UserSecondName{secondName: strValue}, nil
}

func (u UserSecondName) String() string {
	return u.secondName
}
