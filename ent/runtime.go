// Code generated by ent, DO NOT EDIT.

package ent

import (
	"time"

	"github.com/kyh0703/stock-server/ent/post"
	"github.com/kyh0703/stock-server/ent/schema"
	"github.com/kyh0703/stock-server/ent/user"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	postFields := schema.Post{}.Fields()
	_ = postFields
	// postDescPublishAt is the schema descriptor for publish_at field.
	postDescPublishAt := postFields[3].Descriptor()
	// post.DefaultPublishAt holds the default value on creation for the publish_at field.
	post.DefaultPublishAt = postDescPublishAt.Default.(func() time.Time)
	userFields := schema.User{}.Fields()
	_ = userFields
	// userDescEmail is the schema descriptor for email field.
	userDescEmail := userFields[0].Descriptor()
	// user.EmailValidator is a validator for the "email" field. It is called by the builders before save.
	user.EmailValidator = userDescEmail.Validators[0].(func(string) error)
	// userDescUsername is the schema descriptor for username field.
	userDescUsername := userFields[1].Descriptor()
	// user.UsernameValidator is a validator for the "username" field. It is called by the builders before save.
	user.UsernameValidator = userDescUsername.Validators[0].(func(string) error)
	// userDescCreateAt is the schema descriptor for create_at field.
	userDescCreateAt := userFields[3].Descriptor()
	// user.DefaultCreateAt holds the default value on creation for the create_at field.
	user.DefaultCreateAt = userDescCreateAt.Default.(func() time.Time)
	// userDescUpdateAt is the schema descriptor for update_at field.
	userDescUpdateAt := userFields[4].Descriptor()
	// user.DefaultUpdateAt holds the default value on creation for the update_at field.
	user.DefaultUpdateAt = userDescUpdateAt.Default.(func() time.Time)
	// user.UpdateDefaultUpdateAt holds the default value on update for the update_at field.
	user.UpdateDefaultUpdateAt = userDescUpdateAt.UpdateDefault.(func() time.Time)
}
