package moderation

import (
	"baritone/bot/storage"
	"context"
	"time"
)

// WarnUser applies a warning to a user
func WarnUser(userID uint64, reason string) (int, error) {
	storage.RegisterUser(userID)
	warning, err := storage.Client.Warning.Create().SetUserID(userID).SetReason(reason).SetDate(time.Now()).Save(context.Background())
	return warning.ID, err
}
