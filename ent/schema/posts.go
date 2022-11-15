package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Post holds the schema definition for the Post entity.
type Posts struct {
	ent.Schema
}

// Fields of the Post.
func (Posts) Fields() []ent.Field {
	return []ent.Field{
		field.String("title"),
		field.String("body"),
		field.Strings("tags"),
		field.Time("publishAt").
			StorageKey("publish_at").
			Default(time.Now),
	}
}

// Edges of the Post.
func (Posts) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", Users.Type).
			Ref("posts").
			Unique(),
	}
}
