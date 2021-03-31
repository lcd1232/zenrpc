package content2

import (
	"github.com/lcd1232/zenrpc/v3"
	"github.com/lcd1232/zenrpc/v3/testdata/auth"
)

type Service struct {
	zenrpc.Service

	AuthService *auth.Service
}

func NewAuthService() *Service {
	return &Service{
		AuthService: auth.NewAuthService(),
	}
}
