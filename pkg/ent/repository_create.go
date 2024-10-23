// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/kamijoucen/notesync/pkg/ent/repository"
)

// RepositoryCreate is the builder for creating a Repository entity.
type RepositoryCreate struct {
	config
	mutation *RepositoryMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetName sets the "name" field.
func (rc *RepositoryCreate) SetName(s string) *RepositoryCreate {
	rc.mutation.SetName(s)
	return rc
}

// SetDescription sets the "description" field.
func (rc *RepositoryCreate) SetDescription(s string) *RepositoryCreate {
	rc.mutation.SetDescription(s)
	return rc
}

// SetPath sets the "path" field.
func (rc *RepositoryCreate) SetPath(s string) *RepositoryCreate {
	rc.mutation.SetPath(s)
	return rc
}

// SetSize sets the "size" field.
func (rc *RepositoryCreate) SetSize(i int) *RepositoryCreate {
	rc.mutation.SetSize(i)
	return rc
}

// SetCreatedAt sets the "created_at" field.
func (rc *RepositoryCreate) SetCreatedAt(t time.Time) *RepositoryCreate {
	rc.mutation.SetCreatedAt(t)
	return rc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (rc *RepositoryCreate) SetNillableCreatedAt(t *time.Time) *RepositoryCreate {
	if t != nil {
		rc.SetCreatedAt(*t)
	}
	return rc
}

// SetUpdatedAt sets the "updated_at" field.
func (rc *RepositoryCreate) SetUpdatedAt(t time.Time) *RepositoryCreate {
	rc.mutation.SetUpdatedAt(t)
	return rc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (rc *RepositoryCreate) SetNillableUpdatedAt(t *time.Time) *RepositoryCreate {
	if t != nil {
		rc.SetUpdatedAt(*t)
	}
	return rc
}

// Mutation returns the RepositoryMutation object of the builder.
func (rc *RepositoryCreate) Mutation() *RepositoryMutation {
	return rc.mutation
}

// Save creates the Repository in the database.
func (rc *RepositoryCreate) Save(ctx context.Context) (*Repository, error) {
	rc.defaults()
	return withHooks(ctx, rc.sqlSave, rc.mutation, rc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (rc *RepositoryCreate) SaveX(ctx context.Context) *Repository {
	v, err := rc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (rc *RepositoryCreate) Exec(ctx context.Context) error {
	_, err := rc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (rc *RepositoryCreate) ExecX(ctx context.Context) {
	if err := rc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (rc *RepositoryCreate) defaults() {
	if _, ok := rc.mutation.CreatedAt(); !ok {
		v := repository.DefaultCreatedAt()
		rc.mutation.SetCreatedAt(v)
	}
	if _, ok := rc.mutation.UpdatedAt(); !ok {
		v := repository.DefaultUpdatedAt()
		rc.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (rc *RepositoryCreate) check() error {
	if _, ok := rc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "Repository.name"`)}
	}
	if _, ok := rc.mutation.Description(); !ok {
		return &ValidationError{Name: "description", err: errors.New(`ent: missing required field "Repository.description"`)}
	}
	if _, ok := rc.mutation.Path(); !ok {
		return &ValidationError{Name: "path", err: errors.New(`ent: missing required field "Repository.path"`)}
	}
	if _, ok := rc.mutation.Size(); !ok {
		return &ValidationError{Name: "size", err: errors.New(`ent: missing required field "Repository.size"`)}
	}
	if _, ok := rc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "Repository.created_at"`)}
	}
	if _, ok := rc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "Repository.updated_at"`)}
	}
	return nil
}

func (rc *RepositoryCreate) sqlSave(ctx context.Context) (*Repository, error) {
	if err := rc.check(); err != nil {
		return nil, err
	}
	_node, _spec := rc.createSpec()
	if err := sqlgraph.CreateNode(ctx, rc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	rc.mutation.id = &_node.ID
	rc.mutation.done = true
	return _node, nil
}

func (rc *RepositoryCreate) createSpec() (*Repository, *sqlgraph.CreateSpec) {
	var (
		_node = &Repository{config: rc.config}
		_spec = sqlgraph.NewCreateSpec(repository.Table, sqlgraph.NewFieldSpec(repository.FieldID, field.TypeInt))
	)
	_spec.OnConflict = rc.conflict
	if value, ok := rc.mutation.Name(); ok {
		_spec.SetField(repository.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if value, ok := rc.mutation.Description(); ok {
		_spec.SetField(repository.FieldDescription, field.TypeString, value)
		_node.Description = value
	}
	if value, ok := rc.mutation.Path(); ok {
		_spec.SetField(repository.FieldPath, field.TypeString, value)
		_node.Path = value
	}
	if value, ok := rc.mutation.Size(); ok {
		_spec.SetField(repository.FieldSize, field.TypeInt, value)
		_node.Size = value
	}
	if value, ok := rc.mutation.CreatedAt(); ok {
		_spec.SetField(repository.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := rc.mutation.UpdatedAt(); ok {
		_spec.SetField(repository.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Repository.Create().
//		SetName(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.RepositoryUpsert) {
//			SetName(v+v).
//		}).
//		Exec(ctx)
func (rc *RepositoryCreate) OnConflict(opts ...sql.ConflictOption) *RepositoryUpsertOne {
	rc.conflict = opts
	return &RepositoryUpsertOne{
		create: rc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Repository.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (rc *RepositoryCreate) OnConflictColumns(columns ...string) *RepositoryUpsertOne {
	rc.conflict = append(rc.conflict, sql.ConflictColumns(columns...))
	return &RepositoryUpsertOne{
		create: rc,
	}
}

type (
	// RepositoryUpsertOne is the builder for "upsert"-ing
	//  one Repository node.
	RepositoryUpsertOne struct {
		create *RepositoryCreate
	}

	// RepositoryUpsert is the "OnConflict" setter.
	RepositoryUpsert struct {
		*sql.UpdateSet
	}
)

// SetName sets the "name" field.
func (u *RepositoryUpsert) SetName(v string) *RepositoryUpsert {
	u.Set(repository.FieldName, v)
	return u
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *RepositoryUpsert) UpdateName() *RepositoryUpsert {
	u.SetExcluded(repository.FieldName)
	return u
}

// SetDescription sets the "description" field.
func (u *RepositoryUpsert) SetDescription(v string) *RepositoryUpsert {
	u.Set(repository.FieldDescription, v)
	return u
}

// UpdateDescription sets the "description" field to the value that was provided on create.
func (u *RepositoryUpsert) UpdateDescription() *RepositoryUpsert {
	u.SetExcluded(repository.FieldDescription)
	return u
}

// SetPath sets the "path" field.
func (u *RepositoryUpsert) SetPath(v string) *RepositoryUpsert {
	u.Set(repository.FieldPath, v)
	return u
}

// UpdatePath sets the "path" field to the value that was provided on create.
func (u *RepositoryUpsert) UpdatePath() *RepositoryUpsert {
	u.SetExcluded(repository.FieldPath)
	return u
}

// SetSize sets the "size" field.
func (u *RepositoryUpsert) SetSize(v int) *RepositoryUpsert {
	u.Set(repository.FieldSize, v)
	return u
}

// UpdateSize sets the "size" field to the value that was provided on create.
func (u *RepositoryUpsert) UpdateSize() *RepositoryUpsert {
	u.SetExcluded(repository.FieldSize)
	return u
}

// AddSize adds v to the "size" field.
func (u *RepositoryUpsert) AddSize(v int) *RepositoryUpsert {
	u.Add(repository.FieldSize, v)
	return u
}

// SetUpdatedAt sets the "updated_at" field.
func (u *RepositoryUpsert) SetUpdatedAt(v time.Time) *RepositoryUpsert {
	u.Set(repository.FieldUpdatedAt, v)
	return u
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *RepositoryUpsert) UpdateUpdatedAt() *RepositoryUpsert {
	u.SetExcluded(repository.FieldUpdatedAt)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create.
// Using this option is equivalent to using:
//
//	client.Repository.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//		).
//		Exec(ctx)
func (u *RepositoryUpsertOne) UpdateNewValues() *RepositoryUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.CreatedAt(); exists {
			s.SetIgnore(repository.FieldCreatedAt)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Repository.Create().
//	    OnConflict(sql.ResolveWithIgnore()).
//	    Exec(ctx)
func (u *RepositoryUpsertOne) Ignore() *RepositoryUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *RepositoryUpsertOne) DoNothing() *RepositoryUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the RepositoryCreate.OnConflict
// documentation for more info.
func (u *RepositoryUpsertOne) Update(set func(*RepositoryUpsert)) *RepositoryUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&RepositoryUpsert{UpdateSet: update})
	}))
	return u
}

// SetName sets the "name" field.
func (u *RepositoryUpsertOne) SetName(v string) *RepositoryUpsertOne {
	return u.Update(func(s *RepositoryUpsert) {
		s.SetName(v)
	})
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *RepositoryUpsertOne) UpdateName() *RepositoryUpsertOne {
	return u.Update(func(s *RepositoryUpsert) {
		s.UpdateName()
	})
}

// SetDescription sets the "description" field.
func (u *RepositoryUpsertOne) SetDescription(v string) *RepositoryUpsertOne {
	return u.Update(func(s *RepositoryUpsert) {
		s.SetDescription(v)
	})
}

// UpdateDescription sets the "description" field to the value that was provided on create.
func (u *RepositoryUpsertOne) UpdateDescription() *RepositoryUpsertOne {
	return u.Update(func(s *RepositoryUpsert) {
		s.UpdateDescription()
	})
}

// SetPath sets the "path" field.
func (u *RepositoryUpsertOne) SetPath(v string) *RepositoryUpsertOne {
	return u.Update(func(s *RepositoryUpsert) {
		s.SetPath(v)
	})
}

// UpdatePath sets the "path" field to the value that was provided on create.
func (u *RepositoryUpsertOne) UpdatePath() *RepositoryUpsertOne {
	return u.Update(func(s *RepositoryUpsert) {
		s.UpdatePath()
	})
}

// SetSize sets the "size" field.
func (u *RepositoryUpsertOne) SetSize(v int) *RepositoryUpsertOne {
	return u.Update(func(s *RepositoryUpsert) {
		s.SetSize(v)
	})
}

// AddSize adds v to the "size" field.
func (u *RepositoryUpsertOne) AddSize(v int) *RepositoryUpsertOne {
	return u.Update(func(s *RepositoryUpsert) {
		s.AddSize(v)
	})
}

// UpdateSize sets the "size" field to the value that was provided on create.
func (u *RepositoryUpsertOne) UpdateSize() *RepositoryUpsertOne {
	return u.Update(func(s *RepositoryUpsert) {
		s.UpdateSize()
	})
}

// SetUpdatedAt sets the "updated_at" field.
func (u *RepositoryUpsertOne) SetUpdatedAt(v time.Time) *RepositoryUpsertOne {
	return u.Update(func(s *RepositoryUpsert) {
		s.SetUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *RepositoryUpsertOne) UpdateUpdatedAt() *RepositoryUpsertOne {
	return u.Update(func(s *RepositoryUpsert) {
		s.UpdateUpdatedAt()
	})
}

// Exec executes the query.
func (u *RepositoryUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for RepositoryCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *RepositoryUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *RepositoryUpsertOne) ID(ctx context.Context) (id int, err error) {
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *RepositoryUpsertOne) IDX(ctx context.Context) int {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// RepositoryCreateBulk is the builder for creating many Repository entities in bulk.
type RepositoryCreateBulk struct {
	config
	err      error
	builders []*RepositoryCreate
	conflict []sql.ConflictOption
}

// Save creates the Repository entities in the database.
func (rcb *RepositoryCreateBulk) Save(ctx context.Context) ([]*Repository, error) {
	if rcb.err != nil {
		return nil, rcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(rcb.builders))
	nodes := make([]*Repository, len(rcb.builders))
	mutators := make([]Mutator, len(rcb.builders))
	for i := range rcb.builders {
		func(i int, root context.Context) {
			builder := rcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*RepositoryMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, rcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = rcb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, rcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, rcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (rcb *RepositoryCreateBulk) SaveX(ctx context.Context) []*Repository {
	v, err := rcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (rcb *RepositoryCreateBulk) Exec(ctx context.Context) error {
	_, err := rcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (rcb *RepositoryCreateBulk) ExecX(ctx context.Context) {
	if err := rcb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Repository.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.RepositoryUpsert) {
//			SetName(v+v).
//		}).
//		Exec(ctx)
func (rcb *RepositoryCreateBulk) OnConflict(opts ...sql.ConflictOption) *RepositoryUpsertBulk {
	rcb.conflict = opts
	return &RepositoryUpsertBulk{
		create: rcb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Repository.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (rcb *RepositoryCreateBulk) OnConflictColumns(columns ...string) *RepositoryUpsertBulk {
	rcb.conflict = append(rcb.conflict, sql.ConflictColumns(columns...))
	return &RepositoryUpsertBulk{
		create: rcb,
	}
}

// RepositoryUpsertBulk is the builder for "upsert"-ing
// a bulk of Repository nodes.
type RepositoryUpsertBulk struct {
	create *RepositoryCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.Repository.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//		).
//		Exec(ctx)
func (u *RepositoryUpsertBulk) UpdateNewValues() *RepositoryUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.CreatedAt(); exists {
				s.SetIgnore(repository.FieldCreatedAt)
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Repository.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
func (u *RepositoryUpsertBulk) Ignore() *RepositoryUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *RepositoryUpsertBulk) DoNothing() *RepositoryUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the RepositoryCreateBulk.OnConflict
// documentation for more info.
func (u *RepositoryUpsertBulk) Update(set func(*RepositoryUpsert)) *RepositoryUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&RepositoryUpsert{UpdateSet: update})
	}))
	return u
}

// SetName sets the "name" field.
func (u *RepositoryUpsertBulk) SetName(v string) *RepositoryUpsertBulk {
	return u.Update(func(s *RepositoryUpsert) {
		s.SetName(v)
	})
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *RepositoryUpsertBulk) UpdateName() *RepositoryUpsertBulk {
	return u.Update(func(s *RepositoryUpsert) {
		s.UpdateName()
	})
}

// SetDescription sets the "description" field.
func (u *RepositoryUpsertBulk) SetDescription(v string) *RepositoryUpsertBulk {
	return u.Update(func(s *RepositoryUpsert) {
		s.SetDescription(v)
	})
}

// UpdateDescription sets the "description" field to the value that was provided on create.
func (u *RepositoryUpsertBulk) UpdateDescription() *RepositoryUpsertBulk {
	return u.Update(func(s *RepositoryUpsert) {
		s.UpdateDescription()
	})
}

// SetPath sets the "path" field.
func (u *RepositoryUpsertBulk) SetPath(v string) *RepositoryUpsertBulk {
	return u.Update(func(s *RepositoryUpsert) {
		s.SetPath(v)
	})
}

// UpdatePath sets the "path" field to the value that was provided on create.
func (u *RepositoryUpsertBulk) UpdatePath() *RepositoryUpsertBulk {
	return u.Update(func(s *RepositoryUpsert) {
		s.UpdatePath()
	})
}

// SetSize sets the "size" field.
func (u *RepositoryUpsertBulk) SetSize(v int) *RepositoryUpsertBulk {
	return u.Update(func(s *RepositoryUpsert) {
		s.SetSize(v)
	})
}

// AddSize adds v to the "size" field.
func (u *RepositoryUpsertBulk) AddSize(v int) *RepositoryUpsertBulk {
	return u.Update(func(s *RepositoryUpsert) {
		s.AddSize(v)
	})
}

// UpdateSize sets the "size" field to the value that was provided on create.
func (u *RepositoryUpsertBulk) UpdateSize() *RepositoryUpsertBulk {
	return u.Update(func(s *RepositoryUpsert) {
		s.UpdateSize()
	})
}

// SetUpdatedAt sets the "updated_at" field.
func (u *RepositoryUpsertBulk) SetUpdatedAt(v time.Time) *RepositoryUpsertBulk {
	return u.Update(func(s *RepositoryUpsert) {
		s.SetUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *RepositoryUpsertBulk) UpdateUpdatedAt() *RepositoryUpsertBulk {
	return u.Update(func(s *RepositoryUpsert) {
		s.UpdateUpdatedAt()
	})
}

// Exec executes the query.
func (u *RepositoryUpsertBulk) Exec(ctx context.Context) error {
	if u.create.err != nil {
		return u.create.err
	}
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the RepositoryCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for RepositoryCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *RepositoryUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
