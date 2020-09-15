package commands

import (
	"baritone/bot/routing/commands"
	"baritone/bot/routing/types"

	"github.com/diamondburned/arikawa/discord"
)

func init() {
	commands.RegisterCommand(types.Command{
		Name:    "Echo",
		Usage:   "Repeat what you say",
		ID:      "echo",
		Matches: []string{"echo"},
		Action:  Echo,
	})
}

func Echo(c *types.Context) {
	if c.FromMember != nil {
		c.Send(discord.Embed{
			Color:       c.AuthorColor(),
			Description: c.RawContent,
			Author: &discord.EmbedAuthor{
				Name: c.AuthorDisplayName(),
				Icon: c.FromAuthor.AvatarURL(),
			},
		}, "main")
	} else {
		c.Send(discord.Embed{
			Description: c.RawContent,
			Author: &discord.EmbedAuthor{
				Name: c.AuthorDisplayName(),
				Icon: c.FromAuthor.AvatarURL(),
			},
		}, "main")
	}
}
