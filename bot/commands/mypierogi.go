package commands

import (
	"baritone/bot/routing/commands"
	"baritone/bot/routing/types"
	"baritone/bot/storage"
	"context"
	"fmt"

	"github.com/diamondburned/arikawa/discord"
)

func init() {
	commands.RegisterCommand(types.Command{
		Name:    "Pierogi",
		Usage:   "Get pierogi",
		ID:      "pierogi",
		Matches: []string{"Get-Pierogi"},
		Action:  Pierogi,
	})
}

func Pierogi(c *types.Context) {
	if len(c.TriggerMessage.Mentions) == 0 {
		storage.RegisterUser(c.FromAuthor.ID)
		c.Send(discord.Embed{
			Color: c.AuthorColor(),
			Title: fmt.Sprintf("You have %d pierogi", storage.Client.User.GetX(context.Background(), c.FromAuthor.ID).Pierogi),
			Author: &discord.EmbedAuthor{
				Name: c.AuthorDisplayName(),
				Icon: c.FromAuthor.AvatarURL(),
			},
		}, "main")
	} else {
		for _, mention := range c.TriggerMessage.Mentions {
			storage.RegisterUser(mention.ID)
			c.Send(discord.Embed{
				Color:       c.AuthorColor(),
				Description: fmt.Sprintf("<@%s> has %d pierogi", mention.ID, storage.Client.User.GetX(context.Background(), mention.ID).Pierogi),
				Author: &discord.EmbedAuthor{
					Name: c.AuthorDisplayName(),
					Icon: c.FromAuthor.AvatarURL(),
				},
			}, "main")
		}
	}
}
