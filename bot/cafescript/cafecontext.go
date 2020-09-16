package cafescript

import (
	"baritone/bot/routing/types"
	"time"
)

type CafeContext struct {
	ctx *types.Context
}

// RecallData recalls data stored in the given context
func (p CafeContext) RecallData(key string) (v interface{}, ok bool) {
	value, ok := p.ctx.Data["pomo-"+key]
	return value, ok
}

// SetData stores data in the context across invocations
func (p CafeContext) SetData(key string, v interface{}) {
	p.ctx.Data["pomo-"+key] = v
}

// NextResponse creates a channel with the contents of the next response
func (p CafeContext) NextResponse() chan string {
	return p.ctx.NextResponse()
}

// AwaitResponse waits for a response
func (p CafeContext) AwaitResponse(time time.Duration) (content string, ok bool) {
	return p.ctx.AwaitResponse(time)
}

// SendMessage sends a message
func (p CafeContext) SendMessage(id string, content interface{}) {
	p.ctx.Send(content, "pomo-"+id)
}
