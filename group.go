package zenrpc

// Group allows share middleware between same
type Group struct {
	middleware []MiddlewareFunc
}

func (g *Group) Use(m ...MiddlewareFunc) {
	g.middleware = append(g.middleware, m...)
}
