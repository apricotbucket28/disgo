package handlers

import (
	"github.com/DisgoOrg/disgo/core"
	"github.com/DisgoOrg/disgo/core/events"
	"github.com/DisgoOrg/disgo/discord"
)

// GuildMemberAddHandler handles api.GuildMemberAddGatewayEvent
type GuildMemberAddHandler struct{}

// EventType returns the api.GatewayGatewayEventType
func (h *GuildMemberAddHandler) EventType() discord.GatewayEventType {
	return discord.GatewayEventTypeGuildMemberAdd
}

// New constructs a new payload receiver for the raw gateway event
func (h *GuildMemberAddHandler) New() interface{} {
	return discord.Member{}
}

// HandleGatewayEvent handles the specific raw gateway event
func (h *GuildMemberAddHandler) HandleGatewayEvent(disgo core.Disgo, eventManager core.EventManager, sequenceNumber int, v interface{}) {
	member, ok := v.(discord.Member)
	if !ok {
		return
	}

	eventManager.Dispatch(&events.GuildMemberJoinEvent{
		GenericGuildMemberEvent: &events.GenericGuildMemberEvent{
			GenericGuildEvent: &events.GenericGuildEvent{
				GenericEvent: events.NewGenericEvent(disgo, sequenceNumber),
				Guild:        disgo.Cache().GuildCache().Get(member.GuildID),
			},
			Member: disgo.EntityBuilder().CreateMember(member.GuildID, member, core.CacheStrategyYes),
		},
	})
}
