package cafescript

import (
	"baritone/bot/routing/types"
	"context"
	"fmt"
	"time"

	"github.com/alecthomas/repr"
	"github.com/diamondburned/arikawa/discord"
	"github.com/mattn/anko/env"
	"github.com/mattn/anko/vm"
)

func ExecuteScript(c *types.Context, script string) {
	ctx, canc := context.WithDeadline(context.Background(), time.Now().Add(time.Second*30))
	defer canc()

	pomo := CafeContext{c}
	environ := env.NewEnv()
	environ.Define("bariston", pomo)

	_, err := vm.ExecuteContext(ctx, environ, nil, script)

	if err != nil {
		c.Send(discord.Embed{
			Title:       "There was an error running your script:",
			Color:       0xff0000,
			Description: fmt.Sprintf("```\n%s\n```", repr.String(err)),
		}, "primary")
	}
}
