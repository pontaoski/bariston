package commands

import (
	"baritone/bot/commands/guildconfig"
	"baritone/bot/routing/commands"
	"baritone/bot/routing/types"
	"baritone/bot/storage"
	"context"

	"github.com/diamondburned/arikawa/discord"
)

func init() {
	commands.RegisterCommand(types.Command{
		Name:    "Config",
		Usage:   "Configure your server",
		ID:      "config",
		Matches: []string{"Config-Util"},
		Action:  Config,
	})
}

// GetConfig gets a config for a guild
func GetConfig(g discord.GuildID) guildconfig.GuildConfig {
	storage.RegisterGuild(g)

	return storage.Client.Guild.
		GetX(context.Background(), g).
		Config
}

// SaveConfig saves a config for a guild
func SaveConfig(g discord.GuildID, conf guildconfig.GuildConfig) {
	storage.RegisterGuild(g)

	storage.Client.Guild.
		GetX(context.Background(), g).
		Update().
		SetConfig(conf).
		SaveX(context.Background())
}

// Config handles configuration
func Config(c *types.Context) {
	_ = GetConfig(c.TriggerMessage.GuildID)
}
