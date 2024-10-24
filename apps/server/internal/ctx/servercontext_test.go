package ctx

import (
	"fmt"
	"testing"

	"github.com/kamijoucen/notesync/pkg/config"
)

// TestCreateServerContext
func TestCreateServerContext(t *testing.T) {
	serverCtx := NewServerContext(&config.ServerConfig{})

	fmt.Println(serverCtx)
}
