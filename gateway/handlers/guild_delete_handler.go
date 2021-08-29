package handlers

import (
	"github.com/DisgoOrg/disgo/core"
	"github.com/DisgoOrg/disgo/core/events"
	"github.com/DisgoOrg/disgo/discord"
)

// GuildDeleteHandler handles api.GuildDeleteGatewayEvent
type GuildDeleteHandler struct{}

// EventType returns the api.GatewayGatewayEventType
func (h *GuildDeleteHandler) EventType() discord.GatewayEventType {
	return discord.GatewayEventTypeGuildDelete
}

// New constructs a new payload receiver for the raw gateway event
func (h *GuildDeleteHandler) New() interface{} {
	return discord.UnavailableGuild{}
}

// HandleGatewayEvent handles the specific raw gateway event
func (h *GuildDeleteHandler) HandleGatewayEvent(disgo core.Disgo, eventManager core.EventManager, sequenceNumber int, v interface{}) {
	guild, ok := v.(discord.UnavailableGuild)
	if !ok {
		return
	}

	if guild.Unavailable {
		coreGuild := disgo.Cache().GuildCache().Get(guild.ID)
		if coreGuild != nil {
			coreGuild.Unavailable = true
		}
	}

	genericGuildEvent := &events.GenericGuildEvent{
		GenericEvent: events.NewGenericEvent(disgo, sequenceNumber),
		Guild:        disgo.Cache().GuildCache().GetCopy(guild.ID),
	}

	if guild.Unavailable {
		eventManager.Dispatch(&events.GuildUnavailableEvent{
			GenericGuildEvent: genericGuildEvent,
		})
	} else {
		disgo.Cache().GuildCache().Uncache(guild.ID)

		eventManager.Dispatch(&events.GuildLeaveEvent{
			GenericGuildEvent: genericGuildEvent,
		})
	}
}
