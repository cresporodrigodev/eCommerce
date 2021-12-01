package register

import (
	"github.com/cresporodrigodev/eCommerce/src/users/app/register"
	"github.com/cresporodrigodev/eCommerce/src/users/domain/repository"
	"github.com/cresporodrigodev/eCommerce/src/users/infrastructure/mock"
	"testing"
)

func TestRegisterNewUser_RegisterNewUser(t *testing.T) {
	type fields struct {
		alias      string
		firstName  string
		secondName string
		email      string
		admin      string
		password   string
		repository repository.UserRepository
	}

	mockRepository := mock.NewMockRepository()

	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name: "all fields with correct value",
			fields: fields{
				alias:      "RodrigoDev",
				firstName:  "Rodrigo",
				secondName: "Crespo",
				email:      "cresporodrigodev@gmail.com",
				admin:      "true",
				password:   "Qwerty1989",
				repository: mockRepository,
			},
			wantErr: false,
		},
		{
			name: "nil alias",
			fields: fields{
				alias:      "",
				firstName:  "Rodrigo",
				secondName: "Crespo",
				email:      "cresporodrigodev@gmail.com",
				admin:      "true",
				password:   "Qwerty1989",
				repository: mockRepository,
			},
			wantErr: true,
		},
		{
			name: "space alias",
			fields: fields{
				alias:      " ",
				firstName:  "Rodrigo",
				secondName: "Crespo",
				email:      "cresporodrigodev@gmail.com",
				admin:      "true",
				password:   "Qwerty1989",
				repository: mockRepository,
			},
			wantErr: true,
		},
		{
			name: "nil first name",
			fields: fields{
				alias:      "RodrigoDev",
				firstName:  "",
				secondName: "Crespo",
				email:      "cresporodrigodev@gmail.com",
				admin:      "true",
				password:   "Qwerty1989",
				repository: mockRepository,
			},
			wantErr: true,
		},
		{
			name: "nil second name",
			fields: fields{
				alias:      "RodrigoDev",
				firstName:  "Rodrigo",
				secondName: "",
				email:      "cresporodrigodev@gmail.com",
				admin:      "true",
				password:   "Qwerty1989",
				repository: mockRepository,
			},
			wantErr: true,
		},
		{
			name: "nil email",
			fields: fields{
				alias:      "RodrigoDev",
				firstName:  "Rodrigo",
				secondName: "Crespo",
				email:      "",
				admin:      "true",
				password:   "Qwerty1989",
				repository: mockRepository,
			},
			wantErr: true,
		},
		{
			name: "email without @",
			fields: fields{
				alias:      "RodrigoDev",
				firstName:  "Rodrigo",
				secondName: "Crespo",
				email:      "cresporodrigodevgmail.com",
				admin:      "true",
				password:   "Qwerty1989",
				repository: mockRepository,
			},
			wantErr: true,
		},
		{
			name: "email without host",
			fields: fields{
				alias:      "RodrigoDev",
				firstName:  "Rodrigo",
				secondName: "Crespo",
				email:      "cresporodrigodev@.com",
				admin:      "true",
				password:   "Qwerty1989",
				repository: mockRepository,
			},
			wantErr: true,
		},
		{
			name: "incorrect 'true' value",
			fields: fields{
				alias:      "RodrigoDev",
				firstName:  "Rodrigo",
				secondName: "Crespo",
				email:      "cresporodrigodev@gmail.com",
				admin:      "trux",
				password:   "Qwerty1989",
				repository: mockRepository,
			},
			wantErr: true,
		},
		{
			name: "incorrect 'false' value",
			fields: fields{
				alias:      "RodrigoDev",
				firstName:  "Rodrigo",
				secondName: "Crespo",
				email:      "cresporodrigodev@gmail.com",
				admin:      "fas",
				password:   "Qwerty1989",
				repository: mockRepository,
			},
			wantErr: true,
		},
		{
			name: "email already exists",
			fields: fields{
				alias:      "RodrigoDev",
				firstName:  "Rodrigo",
				secondName: "Crespo",
				email:      "cresporodrigo42@gmail.com",
				admin:      "true",
				password:   "Qwerty1989",
				repository: mockRepository,
			},
			wantErr: true,
		},
		{
			name: "alias already exists",
			fields: fields{
				alias:      "belubrin",
				firstName:  "Rodrigo",
				secondName: "Crespo",
				email:      "cresporodrigo42@gmail.com",
				admin:      "true",
				password:   "Qwerty1989",
				repository: mockRepository,
			},
			wantErr: true,
		},
		{
			name: "Non capital letter in password",
			fields: fields{
				alias:      "belubrin",
				firstName:  "Rodrigo",
				secondName: "Crespo",
				email:      "cresporodrigo42@gmail.com",
				admin:      "true",
				password:   "qwerty1989",
				repository: mockRepository,
			},
			wantErr: true,
		},
		{
			name: "Non lower letter in password",
			fields: fields{
				alias:      "belubrin",
				firstName:  "Rodrigo",
				secondName: "Crespo",
				email:      "cresporodrigo42@gmail.com",
				admin:      "true",
				password:   "QWERTY163",
				repository: mockRepository,
			},
			wantErr: true,
		},
		{
			name: "Non number in password",
			fields: fields{
				alias:      "belubrin",
				firstName:  "Rodrigo",
				secondName: "Crespo",
				email:      "cresporodrigo42@gmail.com",
				admin:      "true",
				password:   "QwertyQwerty",
				repository: mockRepository,
			},
			wantErr: true,
		},
		{
			name: "No length valid in password",
			fields: fields{
				alias:      "belubrin",
				firstName:  "Rodrigo",
				secondName: "Crespo",
				email:      "cresporodrigo42@gmail.com",
				admin:      "true",
				password:   "Qw1989",
				repository: mockRepository,
			},
			wantErr: true,
		},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r, _ := app.NewRegisterNewUser(
				tt.fields.alias,
				tt.fields.firstName,
				tt.fields.secondName,
				tt.fields.email,
				tt.fields.admin,
				tt.fields.password,
			)
			err := r.RegisterNewUser(tt.fields.repository)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewRegisterNewUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if len(err) > 1 {
				t.Errorf("NewRegisterNewUser() errorLen = %d, wantErr %d", len(err), 1)
			}
		})
	}
}
