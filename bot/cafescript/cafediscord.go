package cafescript

import (
	"reflect"

	"github.com/diamondburned/arikawa/discord"
	"github.com/mattn/anko/env"
)

func init() {
	env.Packages["discord"] = map[string]reflect.Value{
		"": reflect.ValueOf(nil),
	}
	env.PackageTypes["discord"] = map[string]reflect.Type{
		"Embed":          reflect.TypeOf(discord.Embed{}),
		"EmbedFooter":    reflect.TypeOf(discord.EmbedFooter{}),
		"EmbedImage":     reflect.TypeOf(discord.EmbedImage{}),
		"EmbedThumbnail": reflect.TypeOf(discord.EmbedThumbnail{}),
		"EmbedVideo":     reflect.TypeOf(discord.EmbedVideo{}),
		"EmbedProvider":  reflect.TypeOf(discord.EmbedProvider{}),
		"EmbedAuthor":    reflect.TypeOf(discord.EmbedAuthor{}),
		"EmbedField":     reflect.TypeOf(discord.EmbedField{}),
	}
}
