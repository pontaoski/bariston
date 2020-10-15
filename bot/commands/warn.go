package commands

import (
	"baritone/bot/commands/moderation"
	"baritone/bot/middleware"
	"baritone/bot/routing/commands"
	"baritone/bot/routing/types"
	"baritone/ent"
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/diamondburned/arikawa/discord"
	"github.com/xeonx/timeago"
)

func init() {
	commands.RegisterCommand(types.Command{
		Name:    "Warnings",
		Usage:   "Manage warnings",
		ID:      "warn",
		Matches: []string{"o warnings", "User-Warnings"},
		Subcommands: []*types.Command{
			{
				Matches:     []string{"add", "Add"},
				Middlewares: []types.Middleware{middleware.Admin},
				Flags: types.FlagList{
					commands.StringFlag{
						ShortFlag: "r",
						LongFlag:  "reason",
						FlagUsage: "Set a warn reason.",
						Value:     "",
					},
				},
				Action: Warn,
			},
			{
				Matches:     []string{"get", "Get"},
				Middlewares: []types.Middleware{middleware.Admin},
				Flags: types.FlagList{
					commands.BoolFlag{
						ShortFlag: "l",
						LongFlag:  "list",
						FlagUsage: "Show warnings as a list.",
						Value:     false,
					},
				},
				Action: GetWarns,
			},
		},
		Middlewares: []types.Middleware{middleware.Guild},
		Action:      MyWarnings,
	})
}

// WarningListToString converts a list of warnings to a string
func WarningListToString(list []*ent.Warning) string {
	var sb strings.Builder
	for _, warning := range list {
		sb.WriteString(fmt.Sprintf("- %s", warning.Reason))
		sb.WriteString(fmt.Sprintf(" - Issued %s", timeago.English.Format(warning.Date)))
		sb.WriteString(fmt.Sprintf(" - Issued by <@%d>\n", warning.QueryIssuedBy().FirstX(context.Background()).ID))
	}
	return sb.String()
}

// WarningListToEmbedList converts a list of warnings to a list of embeds
func WarningListToEmbedList(list []*ent.Warning) types.EmbedList {
	return types.EmbedList{
		ItemTypeName: "Warning",
		Embeds: func() (ret []discord.Embed) {
			for _, warning := range list {
				ret = append(ret, discord.Embed{
					Title: "Warning",
					Fields: []discord.EmbedField{
						{
							Name:   "Reason",
							Value:  warning.Reason,
							Inline: true,
						},
						{
							Name:   "Issued By",
							Value:  fmt.Sprintf("<@%d>", warning.QueryIssuedBy().FirstX(context.Background()).ID),
							Inline: true,
						},
						{
							Name:   "Issued",
							Value:  timeago.English.Format(warning.Date),
							Inline: true,
						},
					},
					Color: types.NeutralBlue,
				})
			}
			return
		}(),
	}
}

// MyWarnings handles the command where users get information about their warnings
func MyWarnings(c *types.Context) {
	warnings, err := moderation.UserWarnings(c.FromAuthor.ID, c.TriggerMessage.GuildID)
	if err != nil {
		c.Send(c.ErrorEmbed("There was an error getting your warnings."), "result")
		return
	}
	if len(warnings) == 0 {
		c.Send(c.SuccessEmbed("No Warnings", "You have no warnings here! *At least for now...*"), "result")
		return
	}
	c.Send(WarningListToEmbedList(warnings), "result")
}

// GetWarns handles the command that allows a user to get warnings
func GetWarns(c *types.Context) {
	if len(c.TriggerMessage.Mentions) < 1 {
		c.Send(c.ErrorEmbed("Please mention a user to get warnings for. You can edit your message, and I'll respond to it."), "primary")
	}
	warnings, err := moderation.UserWarnings(c.TriggerMessage.Mentions[0].ID, c.TriggerMessage.GuildID)
	if err != nil {
		c.Send(c.ErrorEmbed("There was an error getting warnings."), "result")
		return
	}
	if len(warnings) == 0 {
		c.Send(c.SuccessEmbed("No Warnings", fmt.Sprintf("<@%d> has no warnings here! *At least for now...*", c.TriggerMessage.Mentions[0].ID)), "result")
		return
	}
	if c.IsFlagSet("list") {
		c.Send(c.StatusEmbed("Warnings", WarningListToString(warnings)), "result")
		return
	}
	c.Send(WarningListToEmbedList(warnings), "result")
}

// Warn warns a user
func Warn(c *types.Context) {
	if c.Waiting {
		c.Send(c.ErrorEmbed("Uh, you edited your message while I was waiting on you. Not very nice of you."), "scolding")
		return
	}
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

	warnID, err := moderation.WarnUser(c.FromAuthor.ID, c.TriggerMessage.Mentions[0].ID, c.TriggerMessage.GuildID, reason)
	c.Data["warn-id"] = warnID

	if err != nil {
		c.Send(c.ErrorEmbed("There was an error issuing the warn. Please try again later."), "result")
	}

	c.Send(c.SuccessEmbed("User has been warned!", ""), "result")
}
