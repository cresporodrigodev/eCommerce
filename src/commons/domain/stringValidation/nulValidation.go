package stringValidation

import (
	"errors"
	"strings"
)

type StringValidation struct {
	value string
	err   error
}

func NewStringValidation(value string) StringValidation {
	return StringValidation{
		value: strings.TrimSpace(value),
	}
}

func (s *StringValidation) Validation() (string, error) {
	s.err = nil

	if s.value == "" {
		s.err = errors.New("the value is a null string")
	}

	return s.value, s.err
}
