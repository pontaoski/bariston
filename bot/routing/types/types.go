package types

import (
	flag "github.com/spf13/pflag"

	"github.com/diamondburned/arikawa/discord"
	"github.com/diamondburned/arikawa/session"
	"github.com/diamondburned/dgwidgets"
)

type Action func(c *Context)

type Middleware interface {
	Wrap(Action) Action
	IsWebsiteVisible() bool
	WebsiteDescription() string
}

type BasicMiddleware struct {
	Wrapper        func(Action) Action
	WebsiteVisible bool
	WebsiteDesc    string
}

func (b BasicMiddleware) Wrap(a Action) Action {
	return b.Wrapper(a)
}

func (b BasicMiddleware) IsWebsiteVisible() bool {
	return b.WebsiteVisible
}

func (b BasicMiddleware) WebsiteDescription() string {
	return b.WebsiteDesc
}

type FlagList []Flag

func (fl FlagList) GetFlagSet() *flag.FlagSet {
	var fs flag.FlagSet
	fs.Init("", flag.ContinueOnError)
	for _, flag := range fl {
		flag.Register(&fs)
	}
	return &fs
}

// The flag type.
type Flag interface {
	Long() string
	Short() string
	Usage() string
	Register(*flag.FlagSet)
}

type Command struct {
	Name  string
	Usage string

	Examples string

	ID      string
	Matches []string

	Middlewares []Middleware

	Flags        FlagList
	Action       Action
	DeleteAction Action
	Hidden       bool

	Subcommands []*Command
}

func (c Command) GAction() Action {
	action := c.Action
	for _, mw := range c.Middlewares {
		action = mw.Wrap(action)
	}
	return action
}

type TriggerKind int

const (
	Invalid TriggerKind = iota
	CreateMessage
	EditMessage
)

type Context struct {
	Session        *session.Session
	TriggerMessage discord.Message
	Reason         TriggerKind
	RawContent     string
	FlagSet        flag.FlagSet
	Command        Command
	Data           map[string]interface{}
	IsTag          bool
	FromAuthor     discord.User
	FromMember     *discord.Member
	Waiting        bool

	msgs       map[string]*discord.Message
	paginators map[string]*dgwidgets.Paginator
}

type EmbedList struct {
	ItemTypeName string
	Embeds       []discord.Embed
}

func (c Context) FlagValue(name string) string {
	if c.IsTag {
		return ""
	}
	return c.FlagSet.Lookup(name).Value.String()
}

func (c Context) Arg(i int) string {
	return c.FlagSet.Arg(i)
}

func (c Context) Args() []string {
	return c.FlagSet.Args()
}

func (c Context) IsFlagSet(name string) (yes bool) {
	c.FlagSet.Visit(func(f *flag.Flag) {
		if f.Name == name {
			yes = true
		}
	})
	return
}

func (c Context) NArgs() int {
	return c.FlagSet.NArg()
}

func (c Context) ChoiceFlags(flags ...string) string {
	for _, flag := range flags {
		if c.IsFlagSet(flag) {
			return flag
		}
	}
	return ""
}

func (c Context) AnySet(flags ...string) bool {
	for _, flag := range flags {
		if c.IsFlagSet(flag) {
			return true
		}
	}
	return false
}
