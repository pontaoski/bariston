package storage

import (
	"baritone/bot/commands/guildconfig"
	"context"

	"github.com/diamondburned/arikawa/discord"
)

func RegisterGuild(ID discord.GuildID) {
	if guild, _ := Client.Guild.Get(context.Background(), ID); guild != nil {
		return
	}
	Client.Guild.Create().
		SetID(ID).
		SetConfig(guildconfig.GuildConfig{}).
		SaveX(context.Background())
}
