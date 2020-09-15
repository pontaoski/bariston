package middleware

import (
	"baritone/bot/logger"
	"baritone/bot/routing/types"

	"github.com/diamondburned/arikawa/discord"
)

var (
	// Admin is the middleware that ensures a user has admin perms
	Admin = types.BasicMiddleware{
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
				for _, userRole := range c.FromMember.RoleIDs {
					for _, role := range guild.Roles {
						if role.ID == userRole {
							if role.Permissions.Has(discord.PermissionAdministrator) {
								next(c)
							}
						}
					}
				}
				if guild.OwnerID != c.FromAuthor.ID {
					c.Send(discord.Embed{
						Title: "You must be an admin of the guild to run this command.",
					}, "primary")
					return
				}
				next(c)
			}
		},
	}
)
