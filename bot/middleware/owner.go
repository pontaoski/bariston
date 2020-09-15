package middleware

import (
	"baritone/bot/logger"
	"baritone/bot/routing/types"

	"github.com/diamondburned/arikawa/discord"
)

var (
	// Owner is the middleware that ensures a user is the owner
	Owner = types.BasicMiddleware{
		Wrapper: func(next types.Action) types.Action {
			return func(c *types.Context) {
				if c.TriggerMessage.GuildID.IsNull() {
					c.Send(discord.Embed{
						Title: "This command can only be run in guilds.",
					}, "primary")
					return
				}
				guild, err := c.Session.Guild(c.TriggerMessage.GuildID)
				logger.LogIfError(err)
				if err != nil {
					c.Send(discord.Embed{
						Title: "There was an error verifiying permissions.",
					}, "primary")
					return
				}
				if guild.OwnerID != c.FromAuthor.ID {
					c.Send(discord.Embed{
						Title: "You must be the owner of the guild to run this command.",
					}, "primary")
					return
				}
				next(c)
			}
		},
	}
)
