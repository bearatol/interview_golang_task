package auth_generator

import "context"

type AuthGenerator struct {
	authGenAddr string
}

func NewAuthGenerator(authGenAddr string) *AuthGenerator {
	return &AuthGenerator{authGenAddr: authGenAddr}
}

func (s *AuthGenerator) Ping(ctx context.Context) (response string, err error) {
	return "", nil
}

func (s *AuthGenerator) Generate(ctx context.Context, login string) (token string, err error) {
	return "", nil
}

func (s *AuthGenerator) Validate(ctx context.Context, jwtToken string) error {
	return nil
}
