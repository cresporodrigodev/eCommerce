package valueObject

import (
	"errors"
	"fmt"
	strValidation "github.com/cresporodrigodev/ecommerce/src/commons/domain/stringValidation"
	symValidation "github.com/cresporodrigodev/ecommerce/src/commons/domain/symbolValidation"
)

type UserAlias struct {
	alias string
}

func NewUserAlias(alias string) (UserAlias, error) {
	aliasError := errors.New("invalid user alias")
	lenError := errors.New("user alias must be less than 50 characters")
	symbolError := errors.New("user alias cannot contain special characters")

	strVal := strValidation.NewStringValidation(alias)
	strValue, err := strVal.Validation()

	if err != nil {
		return UserAlias{}, fmt.Errorf("%w; %s", aliasError, alias)
	}

	if len(strValue) > 50 {
		return UserAlias{}, fmt.Errorf("%w: %s", lenError, alias)
	}

	symbolValidation := symValidation.NewSymbolValidation(strValue)
	if hasSymbol := symbolValidation.HasSymbol(); hasSymbol != false {
		return UserAlias{}, fmt.Errorf("%w: %s", symbolError, alias)
	}

	return UserAlias{alias: strValue}, nil
}

func (u UserAlias) String() string {
	return u.alias
}
