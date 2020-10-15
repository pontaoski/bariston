package state

import (
	"sync"

	"github.com/diamondburned/arikawa/discord"
)

// Guild is the live state of a guild
type Guild struct {
	sync.RWMutex

	Werewolf struct {
		Running bool
		Joined  []discord.UserID
		Channel discord.ChannelID
	}
	Data map[string]interface{}
}

var guildStates map[discord.GuildID]*Guild = make(map[discord.GuildID]*Guild)
var guildStatesMutex sync.RWMutex

// GetGuild gets the state for a guild
func GetGuild(id discord.GuildID) *Guild {
	if !id.IsValid() {
		return nil
	}

	guildStatesMutex.RLock()

	if val, ok := guildStates[id]; ok {
		guildStatesMutex.RUnlock()
		return val
	}

	guildStatesMutex.RUnlock()

	guildStatesMutex.Lock()
	defer guildStatesMutex.Unlock()

	guild := new(Guild)
	guildStates[id] = guild

	return guild
}
