package symbolValidation

import (
	"github.com/cresporodrigodev/eCommerce/src/commons/domain/symbolValidation"
	"testing"
)

func TestSymbolValidation_HasSymbol(t *testing.T) {
	type fields struct {
		value string
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{
			name:   "with symbols",
			fields: fields{value: "rodrigo $crespo#"},
			want:   true,
		},
		{
			name:   "whitout symbols",
			fields: fields{value: "rodrigo crespo"},
			want:   false,
		},
	}
	{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := symbolValidation.NewSymbolValidation(tt.fields.value)
			if got := s.HasSymbol(); got != tt.want {
				t.Errorf("HasSymbol() = %v, want %v", got, tt.want)
			}
		})
	}
}
