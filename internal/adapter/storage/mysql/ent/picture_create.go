// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"letsgo-mini-is/internal/adapter/storage/mysql/ent/picture"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// PictureCreate is the builder for creating a Picture entity.
type PictureCreate struct {
	config
	mutation *PictureMutation
	hooks    []Hook
}

// SetContent sets the "content" field.
func (pc *PictureCreate) SetContent(b []byte) *PictureCreate {
	pc.mutation.SetContent(b)
	return pc
}

// SetExtension sets the "extension" field.
func (pc *PictureCreate) SetExtension(s string) *PictureCreate {
	pc.mutation.SetExtension(s)
	return pc
}

// SetID sets the "id" field.
func (pc *PictureCreate) SetID(u uuid.UUID) *PictureCreate {
	pc.mutation.SetID(u)
	return pc
}

// Mutation returns the PictureMutation object of the builder.
func (pc *PictureCreate) Mutation() *PictureMutation {
	return pc.mutation
}

// Save creates the Picture in the database.
func (pc *PictureCreate) Save(ctx context.Context) (*Picture, error) {
	return withHooks(ctx, pc.sqlSave, pc.mutation, pc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (pc *PictureCreate) SaveX(ctx context.Context) *Picture {
	v, err := pc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (pc *PictureCreate) Exec(ctx context.Context) error {
	_, err := pc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pc *PictureCreate) ExecX(ctx context.Context) {
	if err := pc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (pc *PictureCreate) check() error {
	if _, ok := pc.mutation.Content(); !ok {
		return &ValidationError{Name: "content", err: errors.New(`ent: missing required field "Picture.content"`)}
	}
	if v, ok := pc.mutation.Content(); ok {
		if err := picture.ContentValidator(v); err != nil {
			return &ValidationError{Name: "content", err: fmt.Errorf(`ent: validator failed for field "Picture.content": %w`, err)}
		}
	}
	if _, ok := pc.mutation.Extension(); !ok {
		return &ValidationError{Name: "extension", err: errors.New(`ent: missing required field "Picture.extension"`)}
	}
	return nil
}

func (pc *PictureCreate) sqlSave(ctx context.Context) (*Picture, error) {
	if err := pc.check(); err != nil {
		return nil, err
	}
	_node, _spec := pc.createSpec()
	if err := sqlgraph.CreateNode(ctx, pc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(*uuid.UUID); ok {
			_node.ID = *id
		} else if err := _node.ID.Scan(_spec.ID.Value); err != nil {
			return nil, err
		}
	}
	pc.mutation.id = &_node.ID
	pc.mutation.done = true
	return _node, nil
}

func (pc *PictureCreate) createSpec() (*Picture, *sqlgraph.CreateSpec) {
	var (
		_node = &Picture{config: pc.config}
		_spec = sqlgraph.NewCreateSpec(picture.Table, sqlgraph.NewFieldSpec(picture.FieldID, field.TypeUUID))
	)
	if id, ok := pc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := pc.mutation.Content(); ok {
		_spec.SetField(picture.FieldContent, field.TypeBytes, value)
		_node.Content = value
	}
	if value, ok := pc.mutation.Extension(); ok {
		_spec.SetField(picture.FieldExtension, field.TypeString, value)
		_node.Extension = value
	}
	return _node, _spec
}

// PictureCreateBulk is the builder for creating many Picture entities in bulk.
type PictureCreateBulk struct {
	config
	err      error
	builders []*PictureCreate
}

// Save creates the Picture entities in the database.
func (pcb *PictureCreateBulk) Save(ctx context.Context) ([]*Picture, error) {
	if pcb.err != nil {
		return nil, pcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(pcb.builders))
	nodes := make([]*Picture, len(pcb.builders))
	mutators := make([]Mutator, len(pcb.builders))
	for i := range pcb.builders {
		func(i int, root context.Context) {
			builder := pcb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*PictureMutation)
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
					_, err = mutators[i+1].Mutate(root, pcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, pcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
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
		if _, err := mutators[0].Mutate(ctx, pcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (pcb *PictureCreateBulk) SaveX(ctx context.Context) []*Picture {
	v, err := pcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (pcb *PictureCreateBulk) Exec(ctx context.Context) error {
	_, err := pcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pcb *PictureCreateBulk) ExecX(ctx context.Context) {
	if err := pcb.Exec(ctx); err != nil {
		panic(err)
	}
}
