package valueObject

import (
	"encoding/xml"
	"github.com/cresporodrigodev/eCommerce/src/users/domain/valueObject"
	"reflect"
	"testing"
	"time"
)

//func TestCreateAt_String(t *testing.T) {
//	type fields struct {
//		time time.Time
//	}
//	tests := []struct {
//		name   string
//		fields fields
//		want   string
//	}{
//		// TODO: Add test cases.
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			c := valueObject.CreateAt{
//				time: tt.fields.time,
//			}
//			if got := c.String(); got != tt.want {
//				t.Errorf("String() = %v, want %v", got, tt.want)
//			}
//		})
//	}
//}
//
//func TestNewCreateAt(t *testing.T) {
//	tests := []struct {
//		name string
//		want valueObject.CreateAt
//	}{
//		// TODO: Add test cases.
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			if got := valueObject.NewCreateAt(); !reflect.DeepEqual(got, tt.want) {
//				t.Errorf("NewCreateAt() = %v, want %v", got, tt.want)
//			}
//		})
//	}
//}
//
//func TestNewUpdateAt(t *testing.T) {
//	tests := []struct {
//		name string
//		want valueObject.UpdateAt
//	}{
//		// TODO: Add test cases.
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			if got := valueObject.NewUpdateAt(); !reflect.DeepEqual(got, tt.want) {
//				t.Errorf("NewUpdateAt() = %v, want %v", got, tt.want)
//			}
//		})
//	}
//}

func TestNewUserAdmin(t *testing.T) {
	type args struct {
		adminString string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name:    "True",
			args:    args{adminString: "True"},
			want:    "True",
			wantErr: false,
		},
		{
			name:    "Trux",
			args:    args{adminString: "True"},
			wantErr: true,
		},
		{
			name:    "Nil",
			args:    args{adminString: " "},
			wantErr: true,
		},
	}
	{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := valueObject.NewUserAdmin(tt.args.adminString)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewUserAdmin() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got.String() != tt.want {
				t.Errorf("NewUserAdmin() got = %v, want %v", got.String(), tt.want)
			}
		})
	}
}

func TestNewUserAlias(t *testing.T) {
	type args struct {
		alias string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name:    "correct alias",
			args:    args{alias: "RodrigoDev"},
			want:    "RodrigoDev",
			wantErr: false,
		},
		{
			name:    "alias with space",
			args:    args{alias: " RodrigoDev "},
			want:    "RodrigoDev",
			wantErr: false,
		},
		{
			name:    "alias nil",
			args:    args{alias: ""},
			wantErr: true,
		},
		{
			name:    "alias space",
			args:    args{alias: " "},
			wantErr: true,
		},
	}
	{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := valueObject.NewUserAlias(tt.args.alias)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewUserAlias() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got.String() != tt.want {
				t.Errorf("NewUserAlias() got = %v, want %v", got.String(), tt.want)
			}
		})
	}
}

func TestNewUserEmail(t *testing.T) {
	type args struct {
		email string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name:    "email ok",
			args:    args{email: "cresporodrigodev@gmail.com"},
			want:    "cresporodrigodev@gmail.com",
			wantErr: false,
		},
		{
			name:    "email without @",
			args:    args{email: "cresporodrigodevgmail.com"},
			wantErr: true,
		},
		{
			name:    "email without .com",
			args:    args{email: "cresporodrigodev@gmailcom"},
			wantErr: true,
		},
		{
			name:    "email nil",
			args:    args{email: ""},
			wantErr: true,
		},
		{
			name:    "email space",
			args:    args{email: " "},
			wantErr: true,
		},
	}
	{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := valueObject.NewUserEmail(tt.args.email)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewUserEmail() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got.String() != tt.want {
				t.Errorf("NewUserEmail() got = %v, want %v", got.String(), tt.want)
			}
		})
	}
}

func TestNewUserFirstName(t *testing.T) {
	type args struct {
		firstName string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name:    "correct name",
			args:    args{firstName: "Cresporodrigodev"},
			want:    "Cresporodrigodev",
			wantErr: false,
		},
		{
			name:    "space",
			args:    args{firstName: "Cresporodrigodev "},
			wantErr: true,
		},
		{
			name:    "nil",
			args:    args{firstName: ""},
			wantErr: true,
		},
	}
	{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := valueObject.NewUserFirstName(tt.args.firstName)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewUserFirstName() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got.String() != tt.want {
				t.Errorf("NewUserFirstName() got = %v, want %v", got.String(), tt.want)
			}
		})
	}
}

func TestNewUserId(t *testing.T) {
	tests := []struct {
		name    string
		want    valueObject.UserId
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := valueObject.NewUserId()
			if (err != nil) != tt.wantErr {
				t.Errorf("NewUserId() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewUserId() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewUserPassword(t *testing.T) {
	type args struct {
		password string
	}
	tests := []struct {
		name    string
		args    args
		want    valueObject.UserPassword
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := valueObject.NewUserPassword(tt.args.password)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewUserPassword() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewUserPassword() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewUserSecondName(t *testing.T) {
	type args struct {
		secondName string
	}
	tests := []struct {
		name    string
		args    args
		want    valueObject.UserSecondName
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := valueObject.NewUserSecondName(tt.args.secondName)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewUserSecondName() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewUserSecondName() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUpdateAt_String(t *testing.T) {
	type fields struct {
		time time.Time
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := valueObject.UpdateAt{
				time: tt.fields.time,
			}
			if got := c.String(); got != tt.want {
				t.Errorf("String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUserAdmin_String(t *testing.T) {
	type fields struct {
		admin bool
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := valueObject.UserAdmin{
				admin: tt.fields.admin,
			}
			if got := u.String(); got != tt.want {
				t.Errorf("String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUserAlias_String(t *testing.T) {
	type fields struct {
		alias string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := valueObject.UserAlias{
				alias: tt.fields.alias,
			}
			if got := u.String(); got != tt.want {
				t.Errorf("String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUserEmail_String(t *testing.T) {
	type fields struct {
		email string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := valueObject.UserEmail{
				email: tt.fields.email,
			}
			if got := u.String(); got != tt.want {
				t.Errorf("String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUserFirstName_String(t *testing.T) {
	type fields struct {
		firstName string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := valueObject.UserFirstName{
				firstName: tt.fields.firstName,
			}
			if got := u.String(); got != tt.want {
				t.Errorf("String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUserId_String(t *testing.T) {
	type fields struct {
		uuid string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := valueObject.UserId{
				uuid: tt.fields.uuid,
			}
			if got := u.String(); got != tt.want {
				t.Errorf("String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUserPassword_String(t *testing.T) {
	type fields struct {
		password string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := valueObject.UserPassword{
				password: tt.fields.password,
			}
			if got := u.String(); got != tt.want {
				t.Errorf("String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUserSecondName_String(t *testing.T) {
	type fields struct {
		secondName string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := valueObject.UserSecondName{
				secondName: tt.fields.secondName,
			}
			if got := u.String(); got != tt.want {
				t.Errorf("String() = %v, want %v", got, tt.want)
			}
		})
	}
}
