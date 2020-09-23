package moderation

import (
	"baritone/bot/storage"
	"baritone/ent"
	"baritone/ent/guild"
	"baritone/ent/user"
	"baritone/ent/warning"
	"context"
	"time"

	"github.com/diamondburned/arikawa/discord"
)

// WarnUser applies a warning to a user
func WarnUser(from, to discord.UserID, in discord.GuildID, reason string) (int, error) {
	storage.RegisterUser(from)
	storage.RegisterUser(to)
	storage.RegisterGuild(in)
	warning, err := storage.Client.Warning.Create().SetUserID(to).SetIssuedByID(from).SetGuildID(in).SetReason(reason).SetDate(time.Now()).Save(context.Background())
	return warning.ID, err
}

// UserWarnings gets the warnings of a user
func UserWarnings(userID discord.UserID, in discord.GuildID) ([]*ent.Warning, error) {
	storage.RegisterUser(userID)
	storage.RegisterGuild(in)
	return storage.Client.Warning.Query().Where(warning.HasGuildWith(guild.ID(in)), warning.HasUserWith(user.ID(userID))).All(context.Background())
}
