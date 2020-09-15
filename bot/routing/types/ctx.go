package types

import (
	"baritone/bot/logger"
	"fmt"
	"time"

	"github.com/diamondburned/arikawa/discord"
	"github.com/diamondburned/arikawa/gateway"
	"github.com/diamondburned/arikawa/session"
	"github.com/diamondburned/dgwidgets"
)

func (c *Context) Init() {
	c.msgs = make(map[string]*discord.Message)
	c.paginators = make(map[string]*dgwidgets.Paginator)
}

func (c *Context) cleanID(id string) {
	if val, ok := c.paginators[id]; ok {
		val.Widget.Close <- struct{}{}
		delete(c.paginators, id)
	}
}

func waitForMessage(s *session.Session) chan *gateway.MessageCreateEvent {
	channel := make(chan *gateway.MessageCreateEvent)
	var rm func()
	rm = s.AddHandler(func(m *gateway.MessageCreateEvent) {
		channel <- m
		rm()
	})
	return channel
}

func (c *Context) NextResponse() (out chan string) {
	out = make(chan string)
	go func() {
		for {
			select {
			case usermsg := <-waitForMessage(c.Session):
				if usermsg.Author.ID == c.TriggerMessage.Author.ID && usermsg.ChannelID == c.TriggerMessage.ChannelID {
					out <- usermsg.Content
					return
				}
			}
		}
	}()
	return out
}

func (c *Context) AwaitResponse(tenpo time.Duration) (response string, ok bool) {
	timeoutChan := make(chan struct{})
	go func() {
		time.Sleep(tenpo)
		timeoutChan <- struct{}{}
	}()
	for {
		select {
		case msg := <-c.NextResponse():
			return msg, true
		case <-timeoutChan:
			return "", false
		}
	}
}

func (c *Context) SetData(data string, v interface{}) {
	c.data[data] = v
}

func (c *Context) RecallData(data string) (v interface{}, ok bool) {
	v, ok = c.data[data]
	return
}

func (c *Context) sendPaginators(list EmbedList, id string) {
	if val, ok := c.msgs[id]; ok && val != nil {
		c.Session.DeleteMessage(val.ChannelID, val.ID)
	}
	paginator := dgwidgets.NewPaginator(c.Session, c.TriggerMessage.ChannelID)
	c.paginators[id] = paginator
	title := "Item"
	if list.ItemTypeName != "" {
		title = list.ItemTypeName
	}
	for idx, page := range list.Embeds {
		page.Footer.Text = fmt.Sprintf("%s %d/%d", title, idx+1, len(list.Embeds))
		paginator.Add(page)
	}
	paginator.DeleteMessageWhenDone = true
	go paginator.Spawn()
}

func (lhs *Context) ApplyFrom(rhs *Context) {
	lhs.RawContent = rhs.RawContent
	lhs.FlagSet = rhs.FlagSet
}

func (c *Context) AuthorColor() (kule discord.Color) {
	if c.FromMember != nil {
		if len(c.FromMember.RoleIDs) > 0 {
			ok, err := c.Session.Roles(c.TriggerMessage.GuildID)
			logger.LogIfError(err)
			for _, role := range ok {
				if role.ID == c.FromMember.RoleIDs[len(c.FromMember.RoleIDs)-1] {
					kule = role.Color
				}
			}
		}
	}
	return
}

func (c *Context) AuthorDisplayName() string {
	if c.FromMember == nil || c.FromMember.Nick == "" {
		return c.FromAuthor.Username
	}
	return c.FromMember.Nick
}

func (c *Context) ErrorEmbed(content string) discord.Embed {
	return discord.Embed{
		Title:       "Error",
		Description: content,
		Color:       0xff0000,
	}
}

func (c *Context) StatusEmbed(title, content string) discord.Embed {
	return discord.Embed{
		Title:       title,
		Description: content,
		Color:       0x00d485,
	}
}

func (c *Context) SuccessEmbed(title, content string) discord.Embed {
	return discord.Embed{
		Title:       title,
		Description: content,
		Color:       0x3dd425,
	}
}

func (c *Context) Send(content interface{}, id string) {
	var err error
	defer func() {
		logger.LogIfError(err)
	}()

	if val, ok := c.msgs[id]; ok {
		c.cleanID(id)
		switch data := content.(type) {
		case string:
			c.msgs[id], err = c.Session.EditMessage(c.TriggerMessage.ChannelID, val.ID, data, nil, false)
		case discord.Embed:
			c.msgs[id], err = c.Session.EditMessage(c.TriggerMessage.ChannelID, val.ID, "", &data, false)
		case EmbedList:
			c.sendPaginators(data, id)
		}
	} else {
		switch data := content.(type) {
		case string:
			c.msgs[id], err = c.Session.SendMessage(c.TriggerMessage.ChannelID, data, nil)
		case discord.Embed:
			c.msgs[id], err = c.Session.SendMessage(c.TriggerMessage.ChannelID, "", &data)
		case EmbedList:
			c.sendPaginators(data, id)
		}
	}
}
