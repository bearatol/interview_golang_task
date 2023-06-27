package handler

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/mail"

	"github.com/bearatol/interview_golang_task/sevices/core/internal/mapping"
	"github.com/gin-gonic/gin"
)

type User interface {
	UserRegistration(ctx context.Context, user *mapping.UserAvailableFileds) (token []byte, err error)
	UserAuth(ctx context.Context, login, password string) (token []byte, err error)
	UserCheck(ctx context.Context, token string) error
	UserGet(ctx context.Context, userToken string) (*mapping.User, error)
	UserUpdate(ctx context.Context, userToken string, userUpdate *mapping.UserAvailableFileds) error
	UserDelete(ctx context.Context, userToken string) error
}

// @Summary      Registration
// @Description  registration user
// @Tags         users
// @Param        registration  body      mapping.UserAvailableFileds true "user data"
// @Success      200 {object} mapping.UserToken
// @Failure 	 400 {object} errorResponse
// @Router       /users/regis [post]
func (h *Handler) UserRegistration(c *gin.Context) {
	jsonBody, err := io.ReadAll(c.Request.Body)
	if err != nil {
		newErrorResponse(c, h.log, http.StatusBadRequest, err, "invalid data")
		return
	}
	user := &mapping.UserAvailableFileds{}
	if err := json.Unmarshal(jsonBody, user); err != nil {
		newErrorResponse(c, h.log, http.StatusBadRequest, err, "invalid data")
		return
	}

	if len(user.Login) == 0 {
		newErrorResponse(c, h.log, http.StatusBadRequest, err, "invalid login")
		return
	}
	if len(user.Name) == 0 {
		newErrorResponse(c, h.log, http.StatusBadRequest, err, "invalid name")
		return
	}
	if len(user.Password) == 0 {
		newErrorResponse(c, h.log, http.StatusBadRequest, err, "invalid password")
		return
	}
	if _, err := mail.ParseAddress(user.Email); err != nil {
		newErrorResponse(c, h.log, http.StatusBadRequest, err, "invalid email")
		return
	}

	jwtToken, err := h.user.UserRegistration(h.ctx, user)
	if err != nil {
		newErrorResponse(c, h.log, http.StatusBadRequest, err, "cannot registeration a new user")
		return
	}

	c.JSON(http.StatusAccepted, &mapping.UserToken{Token: string(jwtToken)})
}

// @Summary      Auth
// @Description  auth user
// @Tags         users
// @Param        login    query     string  true  "login"
// @Param        password    query     string  true  "password"
// @Success      200 {object} mapping.UserToken
// @Failure 	 400 {object} errorResponse
// @Router       /users/auth [get]
func (h *Handler) UserAuth(c *gin.Context) {
	login := c.Query("login")
	password := c.Query("password")

	if login == "" || password == "" {
		newErrorResponse(c, h.log, http.StatusBadRequest, nil, "invalid login or password")
		return
	}

	token, err := h.user.UserAuth(h.ctx, login, password)
	if err != nil {
		newErrorResponse(c, h.log, http.StatusBadRequest, err, "problem with auth")
		return
	}

	c.JSON(http.StatusOK, &mapping.UserToken{Token: string(token)})
}

// @Summary      Get user information
// @Description  get user information
// @Tags         users
// @Success      200 {object} mapping.User
// @Failure 	 400 {object} errorResponse
// @Security     BearerAuth
// @Router       /users [get]
func (h *Handler) UserGet(c *gin.Context) {
	// token set in middleware
	token := c.Param("token")

	res, err := h.user.UserGet(h.ctx, token)
	if err != nil {
		newErrorResponse(c, h.log, http.StatusBadRequest, err, "cannot get user")
		return
	}

	c.JSON(http.StatusOK, res)
}

// @Summary      Update user
// @Description  update user
// @Tags         users
// @Param        update  body      mapping.UserAvailableFileds true "user data"
// @Success      200 {object} successResponse
// @Failure 	 400 {object} errorResponse
// @Security     BearerAuth
// @Router       /users [put]
func (h *Handler) UserUpdate(c *gin.Context) {
	jsonBody, err := io.ReadAll(c.Request.Body)
	if err != nil {
		newErrorResponse(c, h.log, http.StatusBadRequest, err, "invalid data")
		return
	}
	user := &mapping.UserAvailableFileds{}
	if err := json.Unmarshal(jsonBody, user); err != nil {
		newErrorResponse(c, h.log, http.StatusBadRequest, err, "invalid data")
		return
	}

	if len(user.Login) == 0 {
		newErrorResponse(c, h.log, http.StatusBadRequest, err, "invalid login")
		return
	}
	if len(user.Name) == 0 {
		newErrorResponse(c, h.log, http.StatusBadRequest, err, "invalid name")
		return
	}
	if len(user.Password) == 0 {
		newErrorResponse(c, h.log, http.StatusBadRequest, err, "invalid password")
		return
	}
	if _, err := mail.ParseAddress(user.Email); err != nil {
		newErrorResponse(c, h.log, http.StatusBadRequest, err, "invalid email")
		return
	}

	// token set in middleware
	token := c.Param("token")
	err = h.user.UserUpdate(h.ctx, token, user)
	if err != nil {
		newErrorResponse(c, h.log, http.StatusBadRequest, err, "cannot update user")
		return
	}

	newSuccessResponse(c)
}

// @Summary      Delete user
// @Description  delete user
// @Tags         users
// @Success      200 {object} successResponse
// @Failure 	 400 {object} errorResponse
// @Security     BearerAuth
// @Router       /users [delete]
func (h *Handler) UserDelete(c *gin.Context) {
	// token set in middleware
	token := c.Param("token")

	err := h.user.UserDelete(h.ctx, token)
	if err != nil {
		newErrorResponse(c, h.log, http.StatusBadRequest, err, "cannot delete user")
		return
	}
	newSuccessResponse(c)
}
