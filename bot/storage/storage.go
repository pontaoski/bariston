package storage

import (
	"baritone/bot/logger"
	"baritone/ent"
	"context"

	_ "github.com/mattn/go-sqlite3"
)

var Client *ent.Client

func init() {
	var err error
	Client, err = ent.Open("sqlite3", "file:data.db?_fk=1")
	logger.FatalIfError(logger.DatabaseFailure, err)

	err = Client.Schema.Create(context.Background())
	logger.FatalIfError(logger.DatabaseFailure, err)
}
