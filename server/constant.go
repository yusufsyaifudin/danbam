package server

import "context"

const (
	ContentTypeJSON = "application/json"
	startTimeKey    = "startTime"
)

// Handler represents an api handler
type Handler func(context.Context, Request) Response

// Middleware is abstraction for middleware
type Middleware func(Handler) Handler

// ChainMiddleware Implementing this idea https://hackernoon.com/simple-http-middleware-with-go-79a4ad62889b
func ChainMiddleware(mw ...Middleware) Middleware {
	return func(final Handler) Handler {
		return func(ctx context.Context, req Request) Response {
			last := final
			for i := len(mw) - 1; i >= 0; i-- {
				last = mw[i](last)
			}

			// last middleware
			return last(ctx, req)
		}
	}
}
