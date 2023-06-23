package service

import "context"

type AuthGenerator interface {
	Ping(ctx context.Context) (response string, err error)
	Generate(ctx context.Context, login string) (jwtToken string, err error)
	Validate(ctx context.Context, jwtToken string) error
}

func (s *Service) Ping(ctx context.Context) (response string, err error) {
	return "", nil
}

func (s *Service) Generate(ctx context.Context, login string) (token string, err error) {
	return "", nil
}

func (s *Service) Validate(ctx context.Context, jwtToken string) error {
	return nil
}
