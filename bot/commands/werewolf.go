package commands

import (
	"baritone/bot/middleware"
	"baritone/bot/routing/commands"
	"baritone/bot/routing/types"
)

func init() {
	commands.RegisterCommand(types.Command{
		Name:        "Werewolf",
		Usage:       "Play the werewolf game",
		ID:          "werewolf",
		Matches:     []string{"Werewolf-Game"},
		Middlewares: []types.Middleware{middleware.Guild},
		Action:      Werewolf,
	})
}

// Werewolf manages the werewolf game
func Werewolf(c *types.Context) {
	switch c.Arg(0) {
	case "Status":
		if c.GuildState.Werewolf.Running {
			c.Send("There is a werewolf game running", "main")
		} else {
			c.Send("There is not a werewolf game running", "main")
		}
	case "Start":
		if c.GuildState.Werewolf.Running {
			c.Send("There is already a werewolf game running", "main")
		} else {
			c.GuildState.Werewolf.Running = true
			c.Send("Starting a game of werewolf...", "main")
		}
	case "Stop":
		if c.GuildState.Werewolf.Running {
			c.Send("Stopping the game of werewolf...", "main")
			c.GuildState.Werewolf.Running = false
		} else {
			c.Send("There is no werewolf game running", "main")
		}
	}
}
