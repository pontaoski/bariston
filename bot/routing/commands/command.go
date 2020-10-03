package commands

import (
	"baritone/bot/routing/types"
	"strings"
	"unicode"

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

func ShellSplit(content string) (words []string) {
	currentWord := ""
	var parseUntil rune

	for _, char := range content {
		switch {
		case parseUntil != 0:
			if char == parseUntil {
				parseUntil = 0
			} else {
				currentWord += string(char)
			}
		case char == '\'':
			parseUntil = '\''
		case char == '"':
			parseUntil = '"'
		case unicode.IsSpace(char):
			if currentWord != "" {
				words = append(words, currentWord)
				currentWord = ""
			}
		default:
			currentWord += string(char)
		}
	}

	if currentWord != "" {
		words = append(words, currentWord)
	}

	return
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
	ctx.FlagSet.Parse(ShellSplit(ctx.RawContent))
	ok = true
	ctx.Init()
	return
}
