package content1

import "github.com/lcd1232/zenrpc/v3"

type Foo struct{}

func (f *Foo) Bar() error {
	return nil
}

type Service struct {
	zenrpc.Service
	f *Foo
}

//zenrpc:sum2
func (s *Service) Sum(c zenrpc.Context, a int, b int) (sum int) {
	return a + b
}
