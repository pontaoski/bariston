package handlers

import (
	"baritone/bot/routing/types"

	"github.com/diamondburned/arikawa/gateway"
	"github.com/diamondburned/arikawa/session"
)

var handlerList []types.Handler

func RegisterHandler(h types.Handler) {
	handlerList = append(handlerList, h)
}

func MessageCreate(s *session.Session, e *gateway.MessageCreateEvent) {
	for _, handler := range handlerList {
		if handler.MessageCreate != nil {
			handler.MessageCreate(s, e)
		}
	}
}
