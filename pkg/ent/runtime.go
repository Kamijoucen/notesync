// Code generated by ent, DO NOT EDIT.

package ent

import (
	"time"

	"github.com/kamijoucen/notesync/pkg/ent/repository"
	"github.com/kamijoucen/notesync/pkg/ent/schema"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	repositoryFields := schema.Repository{}.Fields()
	_ = repositoryFields
	// repositoryDescCreatedAt is the schema descriptor for created_at field.
	repositoryDescCreatedAt := repositoryFields[4].Descriptor()
	// repository.DefaultCreatedAt holds the default value on creation for the created_at field.
	repository.DefaultCreatedAt = repositoryDescCreatedAt.Default.(func() time.Time)
	// repositoryDescUpdatedAt is the schema descriptor for updated_at field.
	repositoryDescUpdatedAt := repositoryFields[5].Descriptor()
	// repository.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	repository.DefaultUpdatedAt = repositoryDescUpdatedAt.Default.(func() time.Time)
	// repository.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	repository.UpdateDefaultUpdatedAt = repositoryDescUpdatedAt.UpdateDefault.(func() time.Time)
}
