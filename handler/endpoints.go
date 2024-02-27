package handler

import (
	"errors"
	"net/http"

	"github.com/SawitProRecruitment/UserService/domain"
	"github.com/SawitProRecruitment/UserService/generated"
	"github.com/SawitProRecruitment/UserService/repository"
	"github.com/labstack/echo/v4"
)

func registrationReqToUser(req generated.RegistrationRequest) (user domain.User) {
	user.Name = req.FullName
	user.Phone = req.Phone
	user.Password = req.Password

	return
}

// This is just a test endpoint to get you started. Please delete this endpoint.
// (GET /hello)
func (s *Server) Registration(ctx echo.Context) (err error) {
	var ctx_ = ctx.Request().Context()
	var req generated.RegistrationRequest
	var resp generated.RegistrationResponse
	var user domain.User

	if err = ctx.Bind(&req); err != nil {
		ctx.Logger().Error(err)
		return ctx.JSON(http.StatusBadRequest, nil)
	}

	if errs := registrationReqValidation(req); len(errs) > 0 {
		return ctx.JSON(http.StatusBadRequest, errorResponseFromErrs(errs))
	}

	user = registrationReqToUser(req)

	if err = user.HashPassword(); err != nil {
		ctx.Logger().Error(err)
		return ctx.JSON(http.StatusInternalServerError, nil)
	}

	if err = s.Repository.SaveUser(ctx_, &user); err != nil {
		ctx.Logger().Error(err)
		return ctx.JSON(http.StatusInternalServerError, nil)
	}

	resp.Id = user.Id

	return ctx.JSON(http.StatusOK, resp)
}

func (s *Server) Login(ctx echo.Context) (err error) {

	var ctx_ = ctx.Request().Context()
	var req generated.LoginRequest
	var resp generated.LoginResponse
	var user domain.User

	if err = ctx.Bind(&req); err != nil {
		ctx.Logger().Error(err)
		return ctx.JSON(http.StatusBadRequest, nil)
	}

	if user, err = s.Repository.GetUserByPhone(ctx_, req.Phone); err != nil {
		ctx.Logger().Error(err)
		return ctx.JSON(http.StatusInternalServerError, nil)
	}

	resp.Id = user.Id

	if err = user.ComparePassword(req.Password); err != nil {
		ctx.Logger().Error(err)
		return ctx.JSON(http.StatusInternalServerError, nil)
	}

	if resp.Token, err = user.GenerateToken(); err != nil {
		ctx.Logger().Error(err)
		return ctx.JSON(http.StatusInternalServerError, nil)
	}

	return ctx.JSON(http.StatusOK, resp)
}

func (s *Server) GetProfile(ctx echo.Context) (err error) {
	var ctx_ = ctx.Request().Context()
	var resp generated.GetProfileResponse
	var userId int64
	var user domain.User

	_, ok := ctx.Get(domain.UserIdCtxKey).(float64)
	if !ok {
		ctx.Logger().Error(errors.New("invalid user_id ctx"))
		return ctx.JSON(http.StatusInternalServerError, nil)
	}

	userId = int64(ctx.Get(domain.UserIdCtxKey).(float64))

	if user, err = s.Repository.GetUserById(ctx_, userId); err != nil {
		ctx.Logger().Error(err)
		return ctx.JSON(http.StatusInternalServerError, nil)
	}

	resp.FullName = user.Name
	resp.Phone = user.Phone

	return ctx.JSON(http.StatusOK, resp)
}

func updateProfileRequestToInput(userId int64, req generated.UpdateProfileRequest) repository.UpdateUserInput {
	return repository.UpdateUserInput{
		WhereId:  userId,
		Phone:    *req.Phone,
		FullName: *req.FullName,
	}
}

func (s *Server) UpdateProfile(ctx echo.Context) (err error) {

	var ctx_ = ctx.Request().Context()
	var req generated.UpdateProfileRequest
	var resp generated.UpdateProfileResponse
	var userId int64
	var updateUserInput repository.UpdateUserInput
	var user domain.User

	_, ok := ctx.Get(domain.UserIdCtxKey).(float64)
	if !ok {
		ctx.Logger().Error(errors.New("invalid user_id ctx"))
		return ctx.JSON(http.StatusInternalServerError, nil)
	}

	userId = int64(ctx.Get(domain.UserIdCtxKey).(float64))

	if err = ctx.Bind(&req); err != nil {
		ctx.Logger().Error(err)
		return ctx.JSON(http.StatusBadRequest, nil)
	}

	if errs := updateProfileReqValidation(req); len(errs) > 0 {
		return ctx.JSON(http.StatusBadRequest, errorResponseFromErrs(errs))
	}

	updateUserInput = updateProfileRequestToInput(userId, req)

	if err = s.Repository.UpdateUser(ctx_, updateUserInput); err != nil {
		ctx.Logger().Error(err)
		return ctx.JSON(http.StatusForbidden, nil)
	}

	if user, err = s.Repository.GetUserById(ctx_, userId); err != nil {
		ctx.Logger().Error(err)
		return ctx.JSON(http.StatusForbidden, nil)
	}

	resp.FullName = user.Name
	resp.Phone = user.Phone

	return ctx.JSON(http.StatusOK, resp)
}
