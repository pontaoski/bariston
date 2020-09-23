package middleware

import (
	"baritone/bot/routing/types"

	"github.com/diamondburned/arikawa/discord"
)

var (
	// Guild is the middleware ensuring a command is only run in a guild
	Guild = types.BasicMiddleware{
		Wrapper: func(a types.Action) types.Action {
			return func(c *types.Context) {
				if c.TriggerMessage.GuildID.IsNull() {
					c.Send(discord.Embed{
						Title: "This command can only be run in guilds.",
					}, "primary")
					return
				}
				a(c)
			}
		},
	}
)
