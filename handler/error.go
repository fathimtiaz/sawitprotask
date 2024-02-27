package handler

import "github.com/SawitProRecruitment/UserService/generated"

const internalServerErrorMsg = "internal server error"
const badRequestMsg = "bad request"

func errorResponseFromErr(err error) generated.ErrorResponse {
	return []generated.ErrorMessage{generated.ErrorMessage{Error: err.Error()}}
}

func errorResponseFromErrs(errs []error) (result generated.ErrorResponse) {
	for _, err := range errs {
		result = append(result, generated.ErrorMessage{Error: err.Error()})
	}

	return
}

func errorResponseFromMsg(msg string) generated.ErrorResponse {
	return []generated.ErrorMessage{generated.ErrorMessage{Error: msg}}
}
