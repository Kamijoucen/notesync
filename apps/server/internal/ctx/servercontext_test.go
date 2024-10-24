package ctx

import (
	"context"
	"testing"

	"github.com/kamijoucen/notesync/pkg/config"
)

// TestCreateServerContext
func TestCreateServerContext(t *testing.T) {
	serverCtx := NewServerContext(&config.ServerConfig{})
	serverCtx.Ent.Repository.Create().SetName("test").SetDescription("test").
		SaveX(context.Background())
	rlist := serverCtx.Ent.Repository.Query().AllX(context.Background())

	if len(rlist) == 0 {
		t.Errorf("expect` rlist is not empty, but got empty")
	}
	for _, r := range rlist {
		if r.Name != "test" {
			t.Errorf("expect name is test, but got %s", r.Name)
		}
		if r.Description != "test" {
			t.Errorf("expect description is test, but got %s", r.Description)
		}
	}
	serverCtx.Ent.Repository.Delete().Exec(context.Background())
}
