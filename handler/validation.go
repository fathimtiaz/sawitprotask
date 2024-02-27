package handler

import (
	"errors"

	"github.com/SawitProRecruitment/UserService/generated"
	"github.com/SawitProRecruitment/UserService/helper"
)

var errPhoneTooShort = errors.New("phone number must be at least 10 characters")
var errPhoneTooLong = errors.New("phone number must be at most 13 characters")
var errPhoneArea = errors.New("phone number must start with +62")
var errNameTooShort = errors.New("name must be at least 3 characters")
var errNameTooLong = errors.New("name must be at most 60 characters")
var errPwdTooShort = errors.New("password must be at least 6 characters")
var errPwdTooLong = errors.New("password must be at most 64 characters")
var errPwdNoCap = errors.New("password must have at least 1 capital letter")
var errPwdNoNum = errors.New("password must have at least 1 number")
var errPwdNoSpcChar = errors.New("password must have at least 1 special character")

func phoneValidation(phone string) (errs []error) {
	if len(phone) < 10 {
		errs = append(errs, errPhoneTooShort)
	} else if len(phone) > 3 && phone[:3] != "+62" {
		errs = append(errs, errPhoneArea)
	}

	if len(phone) > 13 {
		errs = append(errs, errPhoneTooLong)
	}

	return
}

func nameValidation(name string) (errs []error) {
	if len(name) < 3 {
		errs = append(errs, errNameTooShort)
	}

	if len(name) > 60 {
		errs = append(errs, errNameTooLong)
	}

	return
}

func passwordValidation(password string) (errs []error) {
	if len(password) < 6 {
		errs = append(errs, errPwdTooShort)
	}

	if len(password) > 64 {
		errs = append(errs, errPwdTooLong)
	}

	if !helper.StringHasCapital(password) {
		errs = append(errs, errPwdNoCap)
	}

	if !helper.StringHasNumber(password) {
		errs = append(errs, errPwdNoNum)
	}

	if !helper.StringHasSpecialChar(password) {
		errs = append(errs, errPwdNoSpcChar)
	}

	return
}

func registrationReqValidation(req generated.RegistrationRequest) (errs []error) {
	errs = append(errs, phoneValidation(req.Phone)...)
	errs = append(errs, nameValidation(req.FullName)...)
	errs = append(errs, passwordValidation(req.Password)...)
	return
}

func updateProfileReqValidation(req generated.UpdateProfileRequest) (errs []error) {
	if req.Phone != nil {
		errs = append(errs, phoneValidation(*req.Phone)...)
	}

	if req.FullName != nil {
		errs = append(errs, nameValidation(*req.FullName)...)
	}

	return
}
