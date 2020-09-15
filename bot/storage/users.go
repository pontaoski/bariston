package storage

import (
	"context"
)

func RegisterUser(ID uint64) {
	if user, _ := Client.User.Get(context.Background(), ID); user != nil {
		return
	}
	Client.User.Create().SetID(ID).Save(context.Background())
}
