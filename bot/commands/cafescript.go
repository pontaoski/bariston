package commands

import (
	"baritone/bot/cafescript"
	"baritone/bot/middleware"
	"baritone/bot/routing/commands"
	"baritone/bot/routing/types"
)

func init() {
	commands.RegisterCommand(types.Command{
		Name:        "Cafescript",
		Usage:       "Program Bariston",
		ID:          "cafescript",
		Matches:     []string{"cafescript execute"},
		Action:      ExecuteScript,
		Middlewares: []types.Middleware{middleware.Owner},
	})
}

func ExecuteScript(c *types.Context) {
	cafescript.ExecuteScript(c, c.RawContent)
}
