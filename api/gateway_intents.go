package api

// GatewayIntents is an extension of the Bit structure used when identifying with discord
type GatewayIntents int64

// Add allows you to add multiple bits together, producing a new bit
func (p GatewayIntents) Add(bits ...Bit) Bit {
	total := GatewayIntents(0)
	for _, bit := range bits {
		total |= bit.(GatewayIntents)
	}
	p |= total
	return p
}

// Remove allows you to subtract multiple bits from the first, producing a new bit
func (p GatewayIntents) Remove(bits ...Bit) Bit {
	total := GatewayIntents(0)
	for _, bit := range bits {
		total |= bit.(GatewayIntents)
	}
	p &^= total
	return p
}

// HasAll will ensure that the bit includes all of the bits entered
func (p GatewayIntents) HasAll(bits ...Bit) bool {
	for _, bit := range bits {
		if !p.Has(bit) {
			return false
		}
	}
	return true
}

// Has will check whether the Bit contains another bit
func (p GatewayIntents) Has(bit Bit) bool {
	return (p & bit.(GatewayIntents)) == bit
}

// MissingAny will check whether the bit is missing any one of the bits
func (p GatewayIntents) MissingAny(bits ...Bit) bool {
	for _, bit := range bits {
		if !p.Has(bit) {
			return true
		}
	}
	return false
}

// Missing will do the inverse of Bit.Has
func (p GatewayIntents) Missing(bit Bit) bool {
	return !p.Has(bit)
}

// Constants for the different bit offsets of GatewayIntents
const (
	GatewayIntentsGuilds GatewayIntents = 1 << iota
	GatewayIntentsGuildMembers
	GatewayIntentsGuildBans
	GatewayIntentsGuildEmojis
	GatewayIntentsGuildIntegrations
	GatewayIntentsGuildWebhooks
	GatewayIntentsGuildInvites
	GatewayIntentsGuildVoiceStates
	GatewayIntentsGuildPresences
	GatewayIntentsGuildMessages
	GatewayIntentsGuildMessageReactions
	GatewayIntentsGuildMessageTyping
	GatewayIntentsDirectMessages
	GatewayIntentsDirectMessageReactions
	GatewayIntentsDirectMessageTyping

	GatewayIntentsNonPrivileged = GatewayIntentsGuilds |
		GatewayIntentsGuildBans |
		GatewayIntentsGuildEmojis |
		GatewayIntentsGuildIntegrations |
		GatewayIntentsGuildWebhooks |
		GatewayIntentsGuildInvites |
		GatewayIntentsGuildVoiceStates |
		GatewayIntentsGuildMessages |
		GatewayIntentsGuildMessageReactions |
		GatewayIntentsGuildMessageTyping |
		GatewayIntentsDirectMessages |
		GatewayIntentsDirectMessageReactions |
		GatewayIntentsDirectMessageTyping
	GatewayIntentsPrivileged = GatewayIntentsGuildMembers |
		GatewayIntentsGuildPresences
	GatewayIntentsAll = GatewayIntentsNonPrivileged |
		GatewayIntentsPrivileged
	GatewayIntentsNone GatewayIntents = 0
)
