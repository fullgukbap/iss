// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"letsgo-mini-is/internal/adapter/storage/mysql/ent/picture"
	"letsgo-mini-is/internal/adapter/storage/mysql/ent/predicate"
	"sync"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
)

const (
	// Operation types.
	OpCreate    = ent.OpCreate
	OpDelete    = ent.OpDelete
	OpDeleteOne = ent.OpDeleteOne
	OpUpdate    = ent.OpUpdate
	OpUpdateOne = ent.OpUpdateOne

	// Node types.
	TypePicture = "Picture"
)

// PictureMutation represents an operation that mutates the Picture nodes in the graph.
type PictureMutation struct {
	config
	op            Op
	typ           string
	id            *uuid.UUID
	content       *[]byte
	extension     *string
	clearedFields map[string]struct{}
	done          bool
	oldValue      func(context.Context) (*Picture, error)
	predicates    []predicate.Picture
}

var _ ent.Mutation = (*PictureMutation)(nil)

// pictureOption allows management of the mutation configuration using functional options.
type pictureOption func(*PictureMutation)

// newPictureMutation creates new mutation for the Picture entity.
func newPictureMutation(c config, op Op, opts ...pictureOption) *PictureMutation {
	m := &PictureMutation{
		config:        c,
		op:            op,
		typ:           TypePicture,
		clearedFields: make(map[string]struct{}),
	}
	for _, opt := range opts {
		opt(m)
	}
	return m
}

// withPictureID sets the ID field of the mutation.
func withPictureID(id uuid.UUID) pictureOption {
	return func(m *PictureMutation) {
		var (
			err   error
			once  sync.Once
			value *Picture
		)
		m.oldValue = func(ctx context.Context) (*Picture, error) {
			once.Do(func() {
				if m.done {
					err = errors.New("querying old values post mutation is not allowed")
				} else {
					value, err = m.Client().Picture.Get(ctx, id)
				}
			})
			return value, err
		}
		m.id = &id
	}
}

// withPicture sets the old Picture of the mutation.
func withPicture(node *Picture) pictureOption {
	return func(m *PictureMutation) {
		m.oldValue = func(context.Context) (*Picture, error) {
			return node, nil
		}
		m.id = &node.ID
	}
}

// Client returns a new `ent.Client` from the mutation. If the mutation was
// executed in a transaction (ent.Tx), a transactional client is returned.
func (m PictureMutation) Client() *Client {
	client := &Client{config: m.config}
	client.init()
	return client
}

// Tx returns an `ent.Tx` for mutations that were executed in transactions;
// it returns an error otherwise.
func (m PictureMutation) Tx() (*Tx, error) {
	if _, ok := m.driver.(*txDriver); !ok {
		return nil, errors.New("ent: mutation is not running in a transaction")
	}
	tx := &Tx{config: m.config}
	tx.init()
	return tx, nil
}

// SetID sets the value of the id field. Note that this
// operation is only accepted on creation of Picture entities.
func (m *PictureMutation) SetID(id uuid.UUID) {
	m.id = &id
}

// ID returns the ID value in the mutation. Note that the ID is only available
// if it was provided to the builder or after it was returned from the database.
func (m *PictureMutation) ID() (id uuid.UUID, exists bool) {
	if m.id == nil {
		return
	}
	return *m.id, true
}

// IDs queries the database and returns the entity ids that match the mutation's predicate.
// That means, if the mutation is applied within a transaction with an isolation level such
// as sql.LevelSerializable, the returned ids match the ids of the rows that will be updated
// or updated by the mutation.
func (m *PictureMutation) IDs(ctx context.Context) ([]uuid.UUID, error) {
	switch {
	case m.op.Is(OpUpdateOne | OpDeleteOne):
		id, exists := m.ID()
		if exists {
			return []uuid.UUID{id}, nil
		}
		fallthrough
	case m.op.Is(OpUpdate | OpDelete):
		return m.Client().Picture.Query().Where(m.predicates...).IDs(ctx)
	default:
		return nil, fmt.Errorf("IDs is not allowed on %s operations", m.op)
	}
}

// SetContent sets the "content" field.
func (m *PictureMutation) SetContent(b []byte) {
	m.content = &b
}

// Content returns the value of the "content" field in the mutation.
func (m *PictureMutation) Content() (r []byte, exists bool) {
	v := m.content
	if v == nil {
		return
	}
	return *v, true
}

// OldContent returns the old "content" field's value of the Picture entity.
// If the Picture object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *PictureMutation) OldContent(ctx context.Context) (v []byte, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldContent is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldContent requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldContent: %w", err)
	}
	return oldValue.Content, nil
}

// ResetContent resets all changes to the "content" field.
func (m *PictureMutation) ResetContent() {
	m.content = nil
}

// SetExtension sets the "extension" field.
func (m *PictureMutation) SetExtension(s string) {
	m.extension = &s
}

// Extension returns the value of the "extension" field in the mutation.
func (m *PictureMutation) Extension() (r string, exists bool) {
	v := m.extension
	if v == nil {
		return
	}
	return *v, true
}

// OldExtension returns the old "extension" field's value of the Picture entity.
// If the Picture object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *PictureMutation) OldExtension(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldExtension is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldExtension requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldExtension: %w", err)
	}
	return oldValue.Extension, nil
}

// ResetExtension resets all changes to the "extension" field.
func (m *PictureMutation) ResetExtension() {
	m.extension = nil
}

// Where appends a list predicates to the PictureMutation builder.
func (m *PictureMutation) Where(ps ...predicate.Picture) {
	m.predicates = append(m.predicates, ps...)
}

// WhereP appends storage-level predicates to the PictureMutation builder. Using this method,
// users can use type-assertion to append predicates that do not depend on any generated package.
func (m *PictureMutation) WhereP(ps ...func(*sql.Selector)) {
	p := make([]predicate.Picture, len(ps))
	for i := range ps {
		p[i] = ps[i]
	}
	m.Where(p...)
}

// Op returns the operation name.
func (m *PictureMutation) Op() Op {
	return m.op
}

// SetOp allows setting the mutation operation.
func (m *PictureMutation) SetOp(op Op) {
	m.op = op
}

// Type returns the node type of this mutation (Picture).
func (m *PictureMutation) Type() string {
	return m.typ
}

// Fields returns all fields that were changed during this mutation. Note that in
// order to get all numeric fields that were incremented/decremented, call
// AddedFields().
func (m *PictureMutation) Fields() []string {
	fields := make([]string, 0, 2)
	if m.content != nil {
		fields = append(fields, picture.FieldContent)
	}
	if m.extension != nil {
		fields = append(fields, picture.FieldExtension)
	}
	return fields
}

// Field returns the value of a field with the given name. The second boolean
// return value indicates that this field was not set, or was not defined in the
// schema.
func (m *PictureMutation) Field(name string) (ent.Value, bool) {
	switch name {
	case picture.FieldContent:
		return m.Content()
	case picture.FieldExtension:
		return m.Extension()
	}
	return nil, false
}

// OldField returns the old value of the field from the database. An error is
// returned if the mutation operation is not UpdateOne, or the query to the
// database failed.
func (m *PictureMutation) OldField(ctx context.Context, name string) (ent.Value, error) {
	switch name {
	case picture.FieldContent:
		return m.OldContent(ctx)
	case picture.FieldExtension:
		return m.OldExtension(ctx)
	}
	return nil, fmt.Errorf("unknown Picture field %s", name)
}

// SetField sets the value of a field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *PictureMutation) SetField(name string, value ent.Value) error {
	switch name {
	case picture.FieldContent:
		v, ok := value.([]byte)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetContent(v)
		return nil
	case picture.FieldExtension:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetExtension(v)
		return nil
	}
	return fmt.Errorf("unknown Picture field %s", name)
}

// AddedFields returns all numeric fields that were incremented/decremented during
// this mutation.
func (m *PictureMutation) AddedFields() []string {
	return nil
}

// AddedField returns the numeric value that was incremented/decremented on a field
// with the given name. The second boolean return value indicates that this field
// was not set, or was not defined in the schema.
func (m *PictureMutation) AddedField(name string) (ent.Value, bool) {
	return nil, false
}

// AddField adds the value to the field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *PictureMutation) AddField(name string, value ent.Value) error {
	switch name {
	}
	return fmt.Errorf("unknown Picture numeric field %s", name)
}

// ClearedFields returns all nullable fields that were cleared during this
// mutation.
func (m *PictureMutation) ClearedFields() []string {
	return nil
}

// FieldCleared returns a boolean indicating if a field with the given name was
// cleared in this mutation.
func (m *PictureMutation) FieldCleared(name string) bool {
	_, ok := m.clearedFields[name]
	return ok
}

// ClearField clears the value of the field with the given name. It returns an
// error if the field is not defined in the schema.
func (m *PictureMutation) ClearField(name string) error {
	return fmt.Errorf("unknown Picture nullable field %s", name)
}

// ResetField resets all changes in the mutation for the field with the given name.
// It returns an error if the field is not defined in the schema.
func (m *PictureMutation) ResetField(name string) error {
	switch name {
	case picture.FieldContent:
		m.ResetContent()
		return nil
	case picture.FieldExtension:
		m.ResetExtension()
		return nil
	}
	return fmt.Errorf("unknown Picture field %s", name)
}

// AddedEdges returns all edge names that were set/added in this mutation.
func (m *PictureMutation) AddedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// AddedIDs returns all IDs (to other nodes) that were added for the given edge
// name in this mutation.
func (m *PictureMutation) AddedIDs(name string) []ent.Value {
	return nil
}

// RemovedEdges returns all edge names that were removed in this mutation.
func (m *PictureMutation) RemovedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// RemovedIDs returns all IDs (to other nodes) that were removed for the edge with
// the given name in this mutation.
func (m *PictureMutation) RemovedIDs(name string) []ent.Value {
	return nil
}

// ClearedEdges returns all edge names that were cleared in this mutation.
func (m *PictureMutation) ClearedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// EdgeCleared returns a boolean which indicates if the edge with the given name
// was cleared in this mutation.
func (m *PictureMutation) EdgeCleared(name string) bool {
	return false
}

// ClearEdge clears the value of the edge with the given name. It returns an error
// if that edge is not defined in the schema.
func (m *PictureMutation) ClearEdge(name string) error {
	return fmt.Errorf("unknown Picture unique edge %s", name)
}

// ResetEdge resets all changes to the edge with the given name in this mutation.
// It returns an error if the edge is not defined in the schema.
func (m *PictureMutation) ResetEdge(name string) error {
	return fmt.Errorf("unknown Picture edge %s", name)
}
