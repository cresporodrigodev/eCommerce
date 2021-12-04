package valueObject

import (
	"github.com/cresporodrigodev/ecommerce/src/users/domain/valueObject"
	"testing"
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
