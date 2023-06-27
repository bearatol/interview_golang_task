package auth_generator

import (
	"bytes"
	"context"
	"io"
	"net/http"

	"github.com/bearatol/interview_golang_task/sevices/core/internal/mapping"
)

const HTTPProtocol = "http"

type AuthGenerator struct {
	authGenAddr string
}

func NewAuthGenerator(ctx context.Context, authGenAddr string) (*AuthGenerator, error) {
	authGen := &AuthGenerator{authGenAddr: authGenAddr}
	return authGen, authGen.Ping(ctx)
}

func (a *AuthGenerator) getRemote(ctx context.Context, url, token string) ([]byte, error) {
	client := http.Client{
		Timeout: mapping.TimeoutConnect,
	}
	req, err := http.NewRequest(
		"GET", HTTPProtocol+"://"+a.authGenAddr+url, nil,
	)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	req.Header.Add("Content-Type", "text/html; charset=utf-8")
	if token != "" {
		req.Header.Add("Authorization", "bearer "+token)
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var b bytes.Buffer
	if _, err := io.Copy(&b, resp.Body); err != nil {
		return nil, err
	}

	return b.Bytes(), nil
}

func (a *AuthGenerator) Ping(ctx context.Context) error {
	_, err := a.getRemote(ctx, "/ping", "")
	return err
}

func (a *AuthGenerator) Generate(ctx context.Context, login string) (token []byte, err error) {
	return a.getRemote(ctx, "/generate?login="+login, "")
}

func (a *AuthGenerator) Validate(ctx context.Context, jwtToken string) error {
	_, err := a.getRemote(ctx, "/validate", jwtToken)
	return err
}
