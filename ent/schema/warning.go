package schema

import (
	"github.com/facebook/ent"
	"github.com/facebook/ent/schema/edge"
	"github.com/facebook/ent/schema/field"
)

// Warning holds the schema definition for the Warning entity.
type Warning struct {
	ent.Schema
}

// Fields of the Warning.
func (Warning) Fields() []ent.Field {
	return []ent.Field{
		field.String("reason"),
		field.Time("date"),
	}
}

// Edges of the Warning.
func (Warning) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("user", User.Type).Unique().Required(),
		edge.To("issuedBy", User.Type).Unique().Required(),
		edge.To("guild", Guild.Type).Unique().Required(),
	}
}
