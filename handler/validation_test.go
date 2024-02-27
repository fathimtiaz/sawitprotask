package handler

import (
	"reflect"
	"testing"

	"github.com/SawitProRecruitment/UserService/generated"
)

func Test_phoneValidation(t *testing.T) {
	type args struct {
		phone string
	}
	tests := []struct {
		name     string
		args     args
		wantErrs []error
	}{
		{
			name:     "empty string",
			args:     args{phone: ""},
			wantErrs: []error{errPhoneTooShort},
		},
		{
			name:     "wrong country code",
			args:     args{phone: "+268111111111"},
			wantErrs: []error{errPhoneArea},
		},
		{
			name:     "too long",
			args:     args{phone: "+26811111111111"},
			wantErrs: []error{errPhoneArea, errPhoneTooLong},
		},
		{
			name:     "valid",
			args:     args{phone: "+628111111111"},
			wantErrs: []error{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotErrs := phoneValidation(tt.args.phone); len(gotErrs) > 0 && !reflect.DeepEqual(gotErrs, tt.wantErrs) ||
				len(gotErrs) > 0 && len(tt.wantErrs) != len(gotErrs) {
				t.Errorf("phoneValidation() = %v, want %v", gotErrs, tt.wantErrs)
			}
		})
	}
}

func Test_nameValidation(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name     string
		args     args
		wantErrs []error
	}{
		{
			name:     "valid",
			args:     args{name: "asep"},
			wantErrs: []error{},
		},
		{
			name:     "too short",
			args:     args{name: ""},
			wantErrs: []error{errNameTooShort},
		},
		{
			name:     "too long",
			args:     args{name: "panjangpanjangpanjangpanjangpanjangpanjangpanjangpanjangpanjangpanjangpanjangpanjangpanjangpanjang"},
			wantErrs: []error{errNameTooLong},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotErrs := nameValidation(tt.args.name); len(gotErrs) > 0 && !reflect.DeepEqual(gotErrs, tt.wantErrs) ||
				len(gotErrs) > 0 && len(tt.wantErrs) != len(gotErrs) {
				t.Errorf("passwordValidation() = %v, want %v", gotErrs, tt.wantErrs)
			}
		})
	}
}

func Test_passwordValidation(t *testing.T) {
	type args struct {
		password string
	}
	tests := []struct {
		name     string
		args     args
		wantErrs []error
	}{
		{
			name:     "valid",
			args:     args{password: "<<SS11"},
			wantErrs: []error{},
		},
		{
			name:     "too short",
			args:     args{password: "<S1"},
			wantErrs: []error{errPwdTooShort},
		},
		{
			name:     "too long",
			args:     args{password: "<S1<S1<S1<S1<S1<S1<S1<S1<S1<S1<S1<S1<S1<S1<S1<S1<S1<S1<S1<S1<S1<S1<S1<S1<S1<S1"},
			wantErrs: []error{errPwdTooLong},
		},
		{
			name:     "no cap",
			args:     args{password: "<<<111"},
			wantErrs: []error{errPwdNoCap},
		},
		{
			name:     "no num",
			args:     args{password: "<<SSSS"},
			wantErrs: []error{errPwdNoNum},
		},
		{
			name:     "no spc char",
			args:     args{password: "11SSSS"},
			wantErrs: []error{errPwdNoSpcChar},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotErrs := passwordValidation(tt.args.password); len(gotErrs) > 0 && !reflect.DeepEqual(gotErrs, tt.wantErrs) ||
				len(gotErrs) > 0 && len(tt.wantErrs) != len(gotErrs) {
				t.Errorf("passwordValidation() = %v, want %v", gotErrs, tt.wantErrs)
			}
		})
	}
}

func Test_updateProfileReqValidation(t *testing.T) {
	type args struct {
		req generated.UpdateProfileRequest
	}
	tests := []struct {
		name     string
		args     args
		wantErrs []error
	}{
		{
			name:     "invalid phone",
			args:     args{generated.UpdateProfileRequest{Phone: new(string)}},
			wantErrs: []error{errPhoneTooShort},
		},
		{
			name:     "invalid name",
			args:     args{generated.UpdateProfileRequest{FullName: new(string)}},
			wantErrs: []error{errNameTooShort},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotErrs := updateProfileReqValidation(tt.args.req); !reflect.DeepEqual(gotErrs, tt.wantErrs) {
				t.Errorf("updateProfileReqValidation() = %v, want %v", gotErrs, tt.wantErrs)
			}
		})
	}
}
