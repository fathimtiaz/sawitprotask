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

func registrationReqValidation(req generated.RegistrationRequest) (errs []error) {
	if len(req.Phone) < 10 {
		errs = append(errs, errPhoneTooShort)
	}

	if len(req.Phone) > 13 {
		errs = append(errs, errPhoneTooLong)
	}

	if req.Phone[:3] != "+62" {
		errs = append(errs, errPhoneArea)
	}

	if len(req.FullName) < 3 {
		errs = append(errs, errNameTooShort)
	}

	if len(req.FullName) > 60 {
		errs = append(errs, errNameTooLong)
	}

	if len(req.Password) < 6 {
		errs = append(errs, errPwdTooShort)
	}

	if len(req.Password) > 64 {
		errs = append(errs, errPwdTooLong)
	}

	if !helper.StringHasCapital(req.Password) {
		errs = append(errs, errPwdNoCap)
	}

	if !helper.StringHasNumber(req.Password) {
		errs = append(errs, errPwdNoNum)
	}

	if !helper.StringHasSpecialChar(req.Password) {
		errs = append(errs, errPwdNoSpcChar)
	}

	return
}
