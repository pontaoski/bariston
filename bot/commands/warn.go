package commands

import (
	"baritone/bot/logger"
	"baritone/bot/middleware"
	"baritone/bot/routing/commands"
	"baritone/bot/routing/types"
	"baritone/bot/storage"
	"context"
	"time"
)

func init() {
	commands.RegisterCommand(types.Command{
		Name:    "Warn",
		Usage:   "Warn a user",
		ID:      "warn",
		Matches: []string{"o warn"},
		Flags: types.FlagList{
			commands.StringFlag{
				ShortFlag: "r",
				LongFlag:  "reason",
				FlagUsage: "Set a warn reason.",
				Value:     "",
			},
		},
		Action:      Warn,
		Middlewares: []types.Middleware{middleware.Admin},
	})
}

func Warn(c *types.Context) {
	logger.Info("Got a warn command!")
	if len(c.TriggerMessage.Mentions) < 1 {
		c.Send(c.ErrorEmbed("Please mention a user to warn."), "primary")
	}
	reason := c.FlagValue("reason")
	if reason == "" {
		c.Send(c.StatusEmbed("Please provide a warn reason", ""), "query")
		var ok bool
		reason, ok = c.AwaitResponse(time.Minute)
		if !ok {
			c.Send(c.ErrorEmbed("Timeout expired."), "error")
		}
	}

	id := uint64(c.TriggerMessage.Mentions[0].ID)

	storage.RegisterUser(id)
	_, err := storage.Client.Warning.Create().SetUserID(id).SetReason(reason).SetDate(time.Now()).Save(context.Background())

	if err != nil {
		c.Send(c.ErrorEmbed("There was an error issuing the warn. Please try again later."), "result")
	}

	c.Send(c.SuccessEmbed("User has been warned!", ""), "result")
}
