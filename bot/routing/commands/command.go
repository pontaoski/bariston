package commands

import (
	"baritone/bot/routing/types"
	"strings"

	iradix "github.com/hashicorp/go-immutable-radix"
)

var commandRadix = iradix.New()
var commandList []types.Command

func RegisterCommand(command types.Command) {
	for _, match := range command.Matches {
		commandRadix, _, _ = commandRadix.Insert([]byte(match), command)
	}
	commandList = append(commandList, command)
}

func LexCommand(content string) (cmd types.Command, ctx types.Context, ok bool) {
	if content == "" {
		return
	}
	prefix, value, ok := commandRadix.Root().LongestPrefix([]byte(content))
	if !ok {
		return
	}
	content = strings.TrimSpace(strings.TrimPrefix(content, string(prefix)))
	ctx.RawContent = content
	cmd = value.(types.Command)
	for _, flag := range cmd.Flags {
		flag.Register(&ctx.FlagSet)
	}
	ctx.FlagSet.Parse(strings.Fields(ctx.RawContent))
	ok = true
	ctx.Init()
	return
}
