package routing

import (
	"baritone/bot/logger"
	"baritone/bot/routing/cache"
	"baritone/bot/routing/commands"
	"baritone/bot/routing/handlers"
	"baritone/bot/routing/types"
	"runtime/debug"

	"github.com/alecthomas/repr"
	"github.com/diamondburned/arikawa/discord"
	"github.com/diamondburned/arikawa/gateway"
	"github.com/diamondburned/arikawa/session"

	stripmd "github.com/writeas/go-strip-markdown"
)

func handleMessage(s *session.Session, m discord.Message, mem *discord.Member) {
	canPanic := func(f func()) {
		defer func() {
			if err := recover(); err != nil {
				s.Client.SendEmbed(m.ChannelID, discord.Embed{
					Title:       "EXTREME ERROR",
					Description: "There was a serious issue handling messages. Please report at once!",
					Color:       types.ErrorRed,
				})
				repr.Println(err)
				println(string(debug.Stack()))
			}
		}()
		f()
	}
	cmd, ctx, ok := commands.LexCommand(stripmd.Strip(m.Content))

	if !ok {
		if val, ok := cache.CommandCache.Get(m.ID); ok {
			data := val.(*types.Context)
			if data.Command.ID != cmd.ID {
				if data.Command.DeleteAction != nil {
					data.Command.DeleteAction(data)
				}
			}
			cache.CommandCache.Remove(m.ID)
		}
		return
	}

	if val, ok := cache.CommandCache.Get(m.ID); ok {
		data := val.(*types.Context)
		if data.Command.ID != cmd.ID {
			if data.Command.DeleteAction != nil {
				data.Command.DeleteAction(data)
			}
		}
		data.ApplyFrom(&ctx)
		data.Reason = types.EditMessage
		go canPanic(func() { cmd.GAction()(data) })
	} else {
		ctx.TriggerMessage = m
		ctx.Reason = types.CreateMessage
		ctx.FromAuthor = m.Author
		ctx.FromMember = mem
		ctx.Session = s
		cache.CommandCache.Add(m.ID, &ctx)
		go canPanic(func() { cmd.GAction()(&ctx) })
	}

}

func BeginRouting(s *session.Session) {
	s.AddHandler(func(m *gateway.MessageCreateEvent) {
		go logger.CanPanic(func() { handlers.MessageCreate(s, m) })
		handleMessage(s, m.Message, m.Member)
	})
	s.AddHandler(func(m *gateway.MessageUpdateEvent) {
		handleMessage(s, m.Message, m.Member)
	})
	logger.FatalIfError(logger.ConnectionFailure, s.Open())
	logger.Info("Started Baritone")
	select {}
}
