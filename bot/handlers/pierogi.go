package handlers

import (
	"baritone/bot/routing/handlers"
	"baritone/bot/routing/types"
	"baritone/bot/storage"
	"context"
	"fmt"
	"strings"

	"github.com/diamondburned/arikawa/discord"
	"github.com/diamondburned/arikawa/gateway"
	"github.com/diamondburned/arikawa/session"
)

func init() {
	handlers.RegisterHandler(types.Handler{
		MessageCreate: Pierogi,
	})
}

func Pierogi(s *session.Session, g *gateway.MessageCreateEvent) {
	if strings.Contains(strings.ToLower(g.Message.Content), "thanks") {
		for _, mention := range g.Message.Mentions {
			storage.RegisterUser(mention.ID)
			pierogiCount := storage.Client.User.
				GetX(context.Background(), mention.ID).
				Update().
				AddPierogi(1).
				SaveX(context.Background()).Pierogi

			s.SendEmbed(
				g.ChannelID,
				discord.Embed{
					Description: fmt.Sprintf("<@%s> got thanked by <@%s>! They now have %d pierogi.", mention.ID.String(), g.Author.ID.String(), pierogiCount),
				},
			)
		}
	}
}
