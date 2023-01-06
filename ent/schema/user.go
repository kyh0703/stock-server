package schema

import (
	"regexp"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.String("email").
			Match(regexp.MustCompile(`(?m)^[\w\.]+@[\w\.]+\.[\w]+$`)).
			MaxLen(100).
			Unique(),
		field.String("username").
			MaxLen(255),
		field.String("password"),
		field.Time("createAt").
			StorageKey("create_at").
			Default(time.Now).
			Immutable(),
		field.Time("updateAt").
			StorageKey("update_at").
			Default(time.Now).
			UpdateDefault(time.Now),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("posts", Post.Type),
	}
}
