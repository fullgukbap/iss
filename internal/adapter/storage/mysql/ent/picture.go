// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"letsgo-mini-is/internal/adapter/storage/mysql/ent/picture"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
)

// Picture is the model entity for the Picture schema.
type Picture struct {
	config `json:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// Content holds the value of the "content" field.
	Content []byte `json:"content,omitempty"`
	// Extension holds the value of the "extension" field.
	Extension    string `json:"extension,omitempty"`
	selectValues sql.SelectValues
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Picture) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case picture.FieldContent:
			values[i] = new([]byte)
		case picture.FieldExtension:
			values[i] = new(sql.NullString)
		case picture.FieldID:
			values[i] = new(uuid.UUID)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Picture fields.
func (pi *Picture) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case picture.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				pi.ID = *value
			}
		case picture.FieldContent:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field content", values[i])
			} else if value != nil {
				pi.Content = *value
			}
		case picture.FieldExtension:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field extension", values[i])
			} else if value.Valid {
				pi.Extension = value.String
			}
		default:
			pi.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Picture.
// This includes values selected through modifiers, order, etc.
func (pi *Picture) Value(name string) (ent.Value, error) {
	return pi.selectValues.Get(name)
}

// Update returns a builder for updating this Picture.
// Note that you need to call Picture.Unwrap() before calling this method if this Picture
// was returned from a transaction, and the transaction was committed or rolled back.
func (pi *Picture) Update() *PictureUpdateOne {
	return NewPictureClient(pi.config).UpdateOne(pi)
}

// Unwrap unwraps the Picture entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (pi *Picture) Unwrap() *Picture {
	_tx, ok := pi.config.driver.(*txDriver)
	if !ok {
		panic("ent: Picture is not a transactional entity")
	}
	pi.config.driver = _tx.drv
	return pi
}

// String implements the fmt.Stringer.
func (pi *Picture) String() string {
	var builder strings.Builder
	builder.WriteString("Picture(")
	builder.WriteString(fmt.Sprintf("id=%v, ", pi.ID))
	builder.WriteString("content=")
	builder.WriteString(fmt.Sprintf("%v", pi.Content))
	builder.WriteString(", ")
	builder.WriteString("extension=")
	builder.WriteString(pi.Extension)
	builder.WriteByte(')')
	return builder.String()
}

// Pictures is a parsable slice of Picture.
type Pictures []*Picture
