// Code generated by ent, DO NOT EDIT.

package users

import (
	"time"
)

const (
	// Label holds the string label denoting the users type in the database.
	Label = "users"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldEmail holds the string denoting the email field in the database.
	FieldEmail = "email"
	// FieldUsername holds the string denoting the username field in the database.
	FieldUsername = "username"
	// FieldPassword holds the string denoting the password field in the database.
	FieldPassword = "password"
	// FieldCreateAt holds the string denoting the createat field in the database.
	FieldCreateAt = "create_at"
	// FieldUpdateAt holds the string denoting the updateat field in the database.
	FieldUpdateAt = "update_at"
	// EdgePosts holds the string denoting the posts edge name in mutations.
	EdgePosts = "posts"
	// Table holds the table name of the users in the database.
	Table = "users"
	// PostsTable is the table that holds the posts relation/edge.
	PostsTable = "posts"
	// PostsInverseTable is the table name for the Posts entity.
	// It exists in this package in order to avoid circular dependency with the "posts" package.
	PostsInverseTable = "posts"
	// PostsColumn is the table column denoting the posts relation/edge.
	PostsColumn = "users_posts"
)

// Columns holds all SQL columns for users fields.
var Columns = []string{
	FieldID,
	FieldEmail,
	FieldUsername,
	FieldPassword,
	FieldCreateAt,
	FieldUpdateAt,
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
	// EmailValidator is a validator for the "email" field. It is called by the builders before save.
	EmailValidator func(string) error
	// UsernameValidator is a validator for the "username" field. It is called by the builders before save.
	UsernameValidator func(string) error
	// DefaultCreateAt holds the default value on creation for the "createAt" field.
	DefaultCreateAt func() time.Time
	// DefaultUpdateAt holds the default value on creation for the "updateAt" field.
	DefaultUpdateAt func() time.Time
	// UpdateDefaultUpdateAt holds the default value on update for the "updateAt" field.
	UpdateDefaultUpdateAt func() time.Time
)