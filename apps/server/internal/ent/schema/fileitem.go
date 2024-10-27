package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// FileItem holds the schema definition for the FileItem entity.
type FileItem struct {
	ent.Schema
}

// Fields of the FileItem.
func (FileItem) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").NotEmpty(),
		field.String("hash").NotEmpty(),
		field.Strings("path"),
	}
}

// Edges of the FileItem.
func (FileItem) Edges() []ent.Edge {
	return nil
}
