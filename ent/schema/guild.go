package schema

import (
	"github.com/diamondburned/arikawa/discord"
	"github.com/facebook/ent"
	"github.com/facebook/ent/schema/edge"
	"github.com/facebook/ent/schema/field"
)

// Guild holds the schema definition for the Guild entity.
type Guild struct {
	ent.Schema
}

// Fields of the Guild.
func (Guild) Fields() []ent.Field {
	return []ent.Field{
		field.Uint64("id").Unique().GoType(discord.GuildID(0)),
	}
}

// Edges of the Guild.
func (Guild) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("warnings", Warning.Type),
	}
}
