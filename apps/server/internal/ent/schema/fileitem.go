package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
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
		field.Bool("is_dir").Default(false),
	}
}

// Edges of the FileItem.
func (FileItem) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("repository", Repository.Type).
			Ref("files").
			Unique(),
		edge.To("children", FileItem.Type).
			From("parent").
			Unique(),
	}
}

// Mixin of the FileItem.
func (FileItem) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.Time{},
	}
}
