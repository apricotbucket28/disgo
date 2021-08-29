package handlers

import (
	"github.com/DisgoOrg/disgo/core"
	"github.com/DisgoOrg/disgo/discord"
	"github.com/DisgoOrg/disgo/gateway/handlers"
)

func init() {
	core.HTTPServerEventHandler = &InteractionCreateWebhookHandler{}
}

// InteractionCreateWebhookHandler handles api.InteractionCreateWebhookEvent
type InteractionCreateWebhookHandler struct{}

// EventType returns the discord.GatewayEventType
func (h *InteractionCreateWebhookHandler) EventType() discord.GatewayEventType {
	return discord.GatewayEventTypeInteractionCreate
}

// New constructs a new payload receiver for the raw gateway event
func (h *InteractionCreateWebhookHandler) New() interface{} {
	return discord.UnmarshalInteraction{}
}

// HandleHTTPEvent handles the specific raw gateway event
func (h *InteractionCreateWebhookHandler) HandleHTTPEvent(disgo core.Disgo, eventManager core.EventManager, c chan discord.InteractionResponse, v interface{}) {
	unmarshalInteraction, ok := v.(discord.UnmarshalInteraction)
	if !ok {
		return
	}

	// we just want to pong all pings
	// no need for any event
	if unmarshalInteraction.Type == discord.InteractionTypePing {
		disgo.Logger().Debugf("received interaction ping")
		c <- discord.InteractionResponse{
			Type: discord.InteractionResponseTypePong,
		}
		return
	}
	handlers.HandleInteraction(disgo, eventManager, -1, c, unmarshalInteraction)
}
