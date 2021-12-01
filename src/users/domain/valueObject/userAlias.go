package valueObject

import (
	"errors"
	"fmt"
	"github.com/cresporodrigodev/eCommerce/src/commons/domain"
)

type UserAlias struct {
	alias string
}

func NewUserAlias(alias string) (UserAlias, error) {
	aliasError := errors.New("invalid user alias")
	lenError := errors.New("user alias must be less than 50 characters")
	symbolError := errors.New("user alias cannot contain special characters")

	str := domain.NewStringValidation(alias, aliasError)
	str.Validation()
	strValue, err := str.GetErrAndValue()

	if err != nil {
		return UserAlias{}, err
	}

	if len(strValue) > 50 {
		return UserAlias{}, fmt.Errorf("%w: %s", lenError, alias)
	}

	symbolValidation := domain.NewSymbolValidation(strValue, symbolError)
	if hasSymbol := symbolValidation.HasSymbol(); hasSymbol != false {
		return UserAlias{}, fmt.Errorf("%w: %s", symbolError, alias)
	}

	return UserAlias{alias: strValue}, nil
}

func (u UserAlias) String() string {
	return u.alias
}
