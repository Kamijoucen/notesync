package schema

import "entgo.io/ent"

// Repository holds the schema definition for the Repository entity.
type Repository struct {
	ent.Schema
}

// Fields of the Repository.
func (Repository) Fields() []ent.Field {
	return nil
}

// Edges of the Repository.
func (Repository) Edges() []ent.Edge {
	return nil
}
