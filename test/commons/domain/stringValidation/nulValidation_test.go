package stringValidation

import (
	"github.com/cresporodrigodev/eCommerce/src/commons/domain/stringValidation"
	"testing"
)

func TestStringValidation_Validation(t *testing.T) {
	type fields struct {
		value string
	}
	tests := []struct {
		name    string
		fields  fields
		want    string
		wantErr bool
	}{
		{
			name:    "nil string",
			fields:  fields{value: ""},
			wantErr: true,
		},
		{
			name:    "trim spaces",
			fields:  fields{value: " "},
			wantErr: true,
		},
		{
			name:    "text with trim spaces",
			fields:  fields{value: " hello "},
			want:    "hello",
			wantErr: false,
		},
	}
	{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := stringValidation.NewStringValidation(tt.fields.value)
			got, err := s.Validation()
			if (err != nil) != tt.wantErr {
				t.Errorf("Validation() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if err != nil {
				return
			}
			if got != tt.want {
				t.Errorf("Validation() got = %v, want %v", got, tt.want)
			}
		})
	}
}
