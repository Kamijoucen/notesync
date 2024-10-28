package ctx

import (
	"context"

	"github.com/labstack/echo/v4"
)

type RequestContext struct {
	Ctx     context.Context
	EchoCtx echo.Context
	ServerContext
}

func NewRequestContext(serverCtx *ServerContext, echoCtx echo.Context) *RequestContext {
	return &RequestContext{
		Ctx:           context.Background(),
		EchoCtx:       echoCtx,
		ServerContext: *serverCtx,
	}
}
