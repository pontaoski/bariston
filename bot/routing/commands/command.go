package commands

import (
	"baritone/bot/routing/types"
	"strings"

	iradix "github.com/hashicorp/go-immutable-radix"
)

var commandRadix = iradix.New()
var commandList []types.Command

func recurseRegistration(prefix string, command *types.Command) {
	if prefix == "" {
		for _, match := range command.Matches {
			commandRadix, _, _ = commandRadix.Insert([]byte(match), command)
			for _, sub := range command.Subcommands {
				recurseRegistration(match, sub)
			}
		}
	} else {
		for _, match := range command.Matches {
			commandRadix, _, _ = commandRadix.Insert([]byte(prefix+" "+match), command)
			for _, sub := range command.Subcommands {
				recurseRegistration(prefix+" "+match, sub)
			}
		}
	}
}

func RegisterCommand(command types.Command) {
	recurseRegistration("", &command)
	commandList = append(commandList, command)
}

func LexCommand(content string) (cmd *types.Command, ctx types.Context, ok bool) {
	if content == "" {
		return
	}
	prefix, value, ok := commandRadix.Root().LongestPrefix([]byte(content))
	if !ok {
		return
	}
	content = strings.TrimSpace(strings.TrimPrefix(content, string(prefix)))
	ctx.RawContent = content
	ctx.Data = make(map[string]interface{})
	cmd = value.(*types.Command)
	if cmd.Action == nil {
		ok = false
		return
	}
	for _, flag := range cmd.Flags {
		flag.Register(&ctx.FlagSet)
	}
	ctx.FlagSet.Parse(strings.Fields(ctx.RawContent))
	ok = true
	ctx.Init()
	return
}
