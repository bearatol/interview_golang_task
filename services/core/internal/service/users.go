package service

import (
	"context"
	"fmt"

	"github.com/bearatol/interview_golang_task/sevices/core/internal/mapping"
)

type RepoUser interface {
	GetUser(ctx context.Context, login string) (*mapping.User, error)
	CreateUser(ctx context.Context, user *mapping.UserAvailableFileds) error
	UpdateUser(ctx context.Context, login string, user *mapping.UserAvailableFileds) error
	DeleteUser(ctx context.Context, login string) error
}

type AuthGenerator interface {
	Ping(ctx context.Context) error
	Generate(ctx context.Context, login string) (jwtToken []byte, err error)
	Validate(ctx context.Context, jwtToken string) error
}

func (s *Service) UserRegistration(ctx context.Context, user *mapping.UserAvailableFileds) (token []byte, err error) {
	user.Password, err = s.hashPassword(user.Password)
	if err != nil {
		return
	}

	if err = s.repoUser.CreateUser(ctx, user); err != nil {
		return
	}

	return s.authGen.Generate(ctx, user.Login)
}

func (s *Service) UserAuth(ctx context.Context, login, password string) (token []byte, err error) {
	res, err := s.repoUser.GetUser(ctx, login)
	if err != nil {
		return nil, err
	}
	if !s.checkPasswordHash(password, res.Password) {
		return nil, fmt.Errorf("password isn't valid")
	}
	return s.authGen.Generate(ctx, login)
}

func (s *Service) UserCheck(ctx context.Context, token string) error {
	return s.authGen.Validate(ctx, token)
}

func (s *Service) UserGet(ctx context.Context, userToken string) (*mapping.User, error) {
	login, err := s.login(userToken)
	if err != nil {
		return nil, err
	}
	return s.repoUser.GetUser(ctx, login)
}

func (s *Service) UserUpdate(ctx context.Context, userToken string, userUpdate *mapping.UserAvailableFileds) error {
	login, err := s.login(userToken)
	if err != nil {
		return err
	}

	return s.repoUser.UpdateUser(ctx, login, userUpdate)
}

func (s *Service) UserDelete(ctx context.Context, userToken string) error {
	login, err := s.login(userToken)
	if err != nil {
		return err
	}

	return s.repoUser.DeleteUser(ctx, login)
}
