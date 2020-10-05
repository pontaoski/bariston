package schema

import (
	"github.com/diamondburned/arikawa/discord"
	"github.com/facebook/ent"
	"github.com/facebook/ent/schema/field"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.Uint64("id").Unique().GoType(discord.UserID(0)),
		field.Int64("pierogi").Default(0),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return nil
}
