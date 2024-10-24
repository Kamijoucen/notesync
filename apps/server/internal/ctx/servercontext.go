package ctx

import (
	"context"
	"log"

	"entgo.io/ent/dialect"
	"github.com/kamijoucen/notesync/apps/server/internal/ent"
	"github.com/kamijoucen/notesync/apps/server/internal/ent/migrate"
	"github.com/kamijoucen/notesync/pkg/config"

	_ "github.com/mattn/go-sqlite3"
)

type ServerContext struct {
	Config *config.ServerConfig
	Ent    *ent.Client
}

func NewServerContext(cfg *config.ServerConfig) *ServerContext {
	svcCtx := &ServerContext{
		Config: cfg,
	}
	svcCtx.initDB()
	return svcCtx
}

// initDB
func (s *ServerContext) initDB() {
	// client, err := ent.Open(dialect.SQLite, "file:ent?mode=memory&cache=shared&_fk=1")
	client, err := ent.Open(dialect.SQLite, "file:./data.db?_fk=1", ent.Debug())
	if err != nil {
		log.Fatalf("failed opening connection to sqlite: %v", err)
		panic(err)
	}
	// Run the auto migration tool.
	if err := client.Schema.Create(context.Background(),
		migrate.WithDropIndex(true),
	); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
		panic(err)
	}
	s.Ent = client
}
