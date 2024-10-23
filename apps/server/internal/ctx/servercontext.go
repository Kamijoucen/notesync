package ctx

import "github.com/kamijoucen/notesync/pkg/config"

type ServerContext struct {
	Config *config.ServerConfig
}

func NewServerContext(cfg *config.ServerConfig) *ServerContext {
	return &ServerContext{
		Config: cfg,
	}
}
