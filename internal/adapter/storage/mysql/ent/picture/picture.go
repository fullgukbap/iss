// Code generated by ent, DO NOT EDIT.

package picture

import (
	"entgo.io/ent/dialect/sql"
)

const (
	// Label holds the string label denoting the picture type in the database.
	Label = "picture"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldContent holds the string denoting the content field in the database.
	FieldContent = "content"
	// FieldExtension holds the string denoting the extension field in the database.
	FieldExtension = "extension"
	// Table holds the table name of the picture in the database.
	Table = "pictures"
)

// Columns holds all SQL columns for picture fields.
var Columns = []string{
	FieldID,
	FieldContent,
	FieldExtension,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// ContentValidator is a validator for the "content" field. It is called by the builders before save.
	ContentValidator func([]byte) error
)

// OrderOption defines the ordering options for the Picture queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByExtension orders the results by the extension field.
func ByExtension(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldExtension, opts...).ToFunc()
}
