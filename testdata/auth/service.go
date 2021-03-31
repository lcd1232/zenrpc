package auth

import (
	"context"
	"net/http"
)

type Service struct{}

func NewAuthService() *Service {
	return &Service{}
}

type Session string

// CreateNewSessionFromRequest creates new session object from http request.
func (s *Service) CreateNewSessionFromRequest(
	ctx context.Context,
	r *http.Request,
) (
	session Session,
	err error,
) {
	return "session-1", err
}
