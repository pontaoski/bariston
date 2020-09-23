package storage

import (
	"context"

	"github.com/diamondburned/arikawa/discord"
)

func RegisterUser(ID discord.UserID) {
	if user, _ := Client.User.Get(context.Background(), ID); user != nil {
		return
	}
	Client.User.Create().SetID(ID).Save(context.Background())
}
