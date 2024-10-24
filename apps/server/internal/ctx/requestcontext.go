package ctx

import (
	"context"

	"github.com/labstack/echo/v4"
)

type RequestContext struct {
	Ctx       context.Context
	EchoCtx   echo.Context
	ServerCtx *ServerContext
}

func NewRequestContext(serverCtx *ServerContext, echoCtx echo.Context) *RequestContext {
	return &RequestContext{
		ServerCtx: serverCtx,
		EchoCtx:   echoCtx,
	}
}
