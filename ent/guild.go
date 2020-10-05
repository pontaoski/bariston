// Code generated by entc, DO NOT EDIT.

package ent

import (
	"baritone/bot/commands/guildconfig"
	"baritone/ent/guild"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/diamondburned/arikawa/discord"
	"github.com/facebook/ent/dialect/sql"
)

// Guild is the model entity for the Guild schema.
type Guild struct {
	config `json:"-"`
	// ID of the ent.
	ID discord.GuildID `json:"id,omitempty"`
	// Config holds the value of the "config" field.
	Config guildconfig.GuildConfig `json:"config,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the GuildQuery when eager-loading is set.
	Edges GuildEdges `json:"edges"`
}

// GuildEdges holds the relations/edges for other nodes in the graph.
type GuildEdges struct {
	// Warnings holds the value of the warnings edge.
	Warnings []*Warning
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// WarningsOrErr returns the Warnings value or an error if the edge
// was not loaded in eager-loading.
func (e GuildEdges) WarningsOrErr() ([]*Warning, error) {
	if e.loadedTypes[0] {
		return e.Warnings, nil
	}
	return nil, &NotLoadedError{edge: "warnings"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Guild) scanValues() []interface{} {
	return []interface{}{
		&sql.NullInt64{}, // id
		&[]byte{},        // config
	}
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Guild fields.
func (gu *Guild) assignValues(values ...interface{}) error {
	if m, n := len(values), len(guild.Columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	value, ok := values[0].(*sql.NullInt64)
	if !ok {
		return fmt.Errorf("unexpected type %T for field id", value)
	}
	gu.ID = discord.GuildID(value.Int64)
	values = values[1:]

	if value, ok := values[0].(*[]byte); !ok {
		return fmt.Errorf("unexpected type %T for field config", values[0])
	} else if value != nil && len(*value) > 0 {
		if err := json.Unmarshal(*value, &gu.Config); err != nil {
			return fmt.Errorf("unmarshal field config: %v", err)
		}
	}
	return nil
}

// QueryWarnings queries the warnings edge of the Guild.
func (gu *Guild) QueryWarnings() *WarningQuery {
	return (&GuildClient{config: gu.config}).QueryWarnings(gu)
}

// Update returns a builder for updating this Guild.
// Note that, you need to call Guild.Unwrap() before calling this method, if this Guild
// was returned from a transaction, and the transaction was committed or rolled back.
func (gu *Guild) Update() *GuildUpdateOne {
	return (&GuildClient{config: gu.config}).UpdateOne(gu)
}

// Unwrap unwraps the entity that was returned from a transaction after it was closed,
// so that all next queries will be executed through the driver which created the transaction.
func (gu *Guild) Unwrap() *Guild {
	tx, ok := gu.config.driver.(*txDriver)
	if !ok {
		panic("ent: Guild is not a transactional entity")
	}
	gu.config.driver = tx.drv
	return gu
}

// String implements the fmt.Stringer.
func (gu *Guild) String() string {
	var builder strings.Builder
	builder.WriteString("Guild(")
	builder.WriteString(fmt.Sprintf("id=%v", gu.ID))
	builder.WriteString(", config=")
	builder.WriteString(fmt.Sprintf("%v", gu.Config))
	builder.WriteByte(')')
	return builder.String()
}

// Guilds is a parsable slice of Guild.
type Guilds []*Guild

func (gu Guilds) config(cfg config) {
	for _i := range gu {
		gu[_i].config = cfg
	}
}
