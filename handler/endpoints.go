package handler

import (
	"fmt"
	"net/http"

	"github.com/SawitProRecruitment/UserService/domain"
	"github.com/SawitProRecruitment/UserService/generated"
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
		fmt.Println(err)
		return ctx.JSON(http.StatusBadRequest, nil)
	}

	user = registrationReqToUser(req)

	if err = user.HashPassword(); err != nil {
		fmt.Println(err)
		return ctx.JSON(http.StatusInternalServerError, nil)
	}

	if err = s.Repository.SaveUser(ctx_, &user); err != nil {
		fmt.Println(err)
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
		fmt.Println(err)
		return ctx.JSON(http.StatusBadRequest, nil)
	}

	if user, err = s.Repository.GetUserByPhone(ctx_, req.Phone); err != nil {
		fmt.Println(err)
		return ctx.JSON(http.StatusInternalServerError, nil)
	}

	resp.Id = user.Id

	if err = user.ComparePassword(req.Password); err != nil {
		fmt.Println(err)
		return ctx.JSON(http.StatusInternalServerError, nil)
	}

	if resp.Token, err = user.GenerateToken(); err != nil {
		fmt.Println(err)
		return ctx.JSON(http.StatusInternalServerError, nil)
	}

	return ctx.JSON(http.StatusOK, resp)
}

func (s *Server) GetProfile(ctx echo.Context) (err error) {
	var ctx_ = ctx.Request().Context()
	var resp generated.GetProfileResponse
	var user domain.User

	if user, err = domain.NewAuthenticatedUser(ctx.Request().Context()); err != nil {
		fmt.Println(err)
		return ctx.JSON(http.StatusInternalServerError, nil)
	}

	if user, err = s.Repository.GetUserByPhone(ctx_, user.Phone); err != nil {
		fmt.Println(err)
		return ctx.JSON(http.StatusInternalServerError, nil)
	}

	resp.FullName = user.Name
	resp.Phone = user.Phone

	return ctx.JSON(http.StatusOK, resp)
}
