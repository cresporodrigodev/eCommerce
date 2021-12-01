package symbolValidation

import "unicode"

type SymbolValidation struct {
	value string
}

func NewSymbolValidation(value string) SymbolValidation {
	return SymbolValidation{
		value: value,
	}
}

func (s SymbolValidation) HasSymbol() bool {
	for _, character := range s.value {
		if unicode.IsSymbol(character) {
			return true
		}
	}
	return false
}
