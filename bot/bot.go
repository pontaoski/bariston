package bot

import (
	_ "baritone/bot/commands"
	. "baritone/bot/configuration"
	_ "baritone/bot/handlers"
	"baritone/bot/logger"
	"baritone/bot/routing"
	"baritone/bot/storage"
	"baritone/bot/web"

	"github.com/diamondburned/arikawa/session"
)

// Main is the entrypoint to the bot
func Main() {
	defer storage.Client.Close()
	go web.Host()
	session, err := session.New("Bot " + Config.Token)
	logger.FatalIfError(logger.ConnectionFailure, err)
	routing.BeginRouting(session)
}
