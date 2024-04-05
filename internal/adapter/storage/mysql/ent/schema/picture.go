package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

type Picture struct {
	ent.Schema
}

func (Picture) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.New()),
		field.Bytes("content").
			MaxLen(1000000),
	}
}
