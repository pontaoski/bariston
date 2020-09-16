package commands

import (
	"baritone/bot/commands/moderation"
	"baritone/bot/logger"
	"baritone/bot/middleware"
	"baritone/bot/routing/commands"
	"baritone/bot/routing/types"
	"time"
)

func init() {
	commands.RegisterCommand(types.Command{
		Name:    "Warn",
		Usage:   "Warn a user",
		ID:      "warn",
		Matches: []string{"o warn", "Warn-User"},
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
	if c.Waiting {
		c.Send(c.ErrorEmbed("Uh, you edited your message while I was waiting on you. Not very nice of you."), "scolding")
		return
	}
	logger.Info("Got a warn command!")
	if len(c.TriggerMessage.Mentions) < 1 {
		c.Send(c.ErrorEmbed("Please mention a user to warn. You can edit your message, and I'll respond to it."), "primary")
	}
	reason := c.FlagValue("reason")
	if reason == "" {
		c.Send(c.StatusEmbed("Please reply with your warn reason within 60 seconds:", ""), "query")
		var ok bool
		c.Waiting = true
		reason, ok = c.AwaitResponse(time.Minute)
		c.Waiting = false
		if !ok {
			c.Send(c.ErrorEmbed("Timeout expired."), "error")
		}
	}

	id := uint64(c.TriggerMessage.Mentions[0].ID)
	warnID, err := moderation.WarnUser(id, reason)
	c.Data["warn-id"] = warnID

	if err != nil {
		c.Send(c.ErrorEmbed("There was an error issuing the warn. Please try again later."), "result")
	}

	c.Send(c.SuccessEmbed("User has been warned!", ""), "result")
}
