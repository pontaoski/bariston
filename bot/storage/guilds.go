package storage

import (
	"context"

	"github.com/diamondburned/arikawa/discord"
)

func RegisterGuild(ID discord.GuildID) {
	if guild, _ := Client.Guild.Get(context.Background(), ID); guild != nil {
		return
	}
	Client.Guild.Create().SetID(ID).Save(context.Background())
}
