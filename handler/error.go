package handler

import "github.com/SawitProRecruitment/UserService/generated"

func errorResponseFromErrs(errs []error) (result generated.ErrorResponse) {
	for _, err := range errs {
		result = append(result, generated.ErrorMessage{Error: err.Error()})
	}

	return
}
