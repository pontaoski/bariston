// Code generated by entc, DO NOT EDIT.

package migrate

import (
	"github.com/facebook/ent/dialect/sql/schema"
	"github.com/facebook/ent/schema/field"
)

var (
	// GuildsColumns holds the columns for the "guilds" table.
	GuildsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
	}
	// GuildsTable holds the schema information for the "guilds" table.
	GuildsTable = &schema.Table{
		Name:        "guilds",
		Columns:     GuildsColumns,
		PrimaryKey:  []*schema.Column{GuildsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{},
	}
	// UsersColumns holds the columns for the "users" table.
	UsersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
	}
	// UsersTable holds the schema information for the "users" table.
	UsersTable = &schema.Table{
		Name:        "users",
		Columns:     UsersColumns,
		PrimaryKey:  []*schema.Column{UsersColumns[0]},
		ForeignKeys: []*schema.ForeignKey{},
	}
	// WarningsColumns holds the columns for the "warnings" table.
	WarningsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "reason", Type: field.TypeString},
		{Name: "date", Type: field.TypeTime},
		{Name: "guild_warnings", Type: field.TypeInt, Nullable: true},
		{Name: "warning_user", Type: field.TypeInt, Nullable: true},
	}
	// WarningsTable holds the schema information for the "warnings" table.
	WarningsTable = &schema.Table{
		Name:       "warnings",
		Columns:    WarningsColumns,
		PrimaryKey: []*schema.Column{WarningsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:  "warnings_guilds_warnings",
				Columns: []*schema.Column{WarningsColumns[3]},

				RefColumns: []*schema.Column{GuildsColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:  "warnings_users_user",
				Columns: []*schema.Column{WarningsColumns[4]},

				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		GuildsTable,
		UsersTable,
		WarningsTable,
	}
)

func init() {
	WarningsTable.ForeignKeys[0].RefTable = GuildsTable
	WarningsTable.ForeignKeys[1].RefTable = UsersTable
}
