package ctx

import "github.com/labstack/echo/v4"

type RequestContext struct {
	EchoCtx   echo.Context
	ServerCtx *ServerContext
}

func NewRequestContext(serverCtx *ServerContext, echoCtx echo.Context) *RequestContext {
	return &RequestContext{
		ServerCtx: serverCtx,
		EchoCtx:   echoCtx,
	}
}
