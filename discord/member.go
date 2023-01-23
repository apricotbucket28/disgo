package discord

import (
	"time"

	"github.com/disgoorg/json"
	"github.com/disgoorg/snowflake/v2"
)

var _ Mentionable = (*Member)(nil)

// Member is a discord GuildMember
type Member struct {
	User                       User           `json:"user"`
	Nick                       *string        `json:"nick"`
	Avatar                     *string        `json:"avatar"`
	RoleIDs                    []snowflake.ID `json:"roles,omitempty"`
	JoinedAt                   time.Time      `json:"joined_at"`
	PremiumSince               *time.Time     `json:"premium_since,omitempty"`
	Deaf                       bool           `json:"deaf,omitempty"`
	Mute                       bool           `json:"mute,omitempty"`
	Flags                      MemberFlags    `json:"flags"`
	Pending                    bool           `json:"pending"`
	CommunicationDisabledUntil *time.Time     `json:"communication_disabled_until"`

	// This field is not present everywhere in the API and often populated by disgo
	GuildID snowflake.ID `json:"guild_id"`
}

func (m Member) String() string {
	return m.User.String()
}

func (m Member) Mention() string {
	return m.String()
}

// EffectiveName returns either the nickname or username depending on if the user has a nickname
func (m Member) EffectiveName() string {
	if m.Nick != nil {
		return *m.Nick
	}
	return m.User.Username
}

func (m Member) EffectiveAvatarURL(opts ...CDNOpt) string {
	if m.Avatar == nil {
		return m.User.EffectiveAvatarURL(opts...)
	}
	if avatar := m.AvatarURL(opts...); avatar != nil {
		return *avatar
	}
	return ""
}

func (m Member) AvatarURL(opts ...CDNOpt) *string {
	if m.Avatar == nil {
		return nil
	}
	url := formatAssetURL(MemberAvatar, opts, m.GuildID, m.User.ID, *m.Avatar)
	return &url
}

func (m Member) CreatedAt() time.Time {
	return m.User.CreatedAt()
}

// MemberAdd is used to add a member via the oauth2 access token to a guild
type MemberAdd struct {
	AccessToken string         `json:"access_token"`
	Nick        string         `json:"nick,omitempty"`
	Roles       []snowflake.ID `json:"roles,omitempty"`
	Mute        bool           `json:"mute,omitempty"`
	Deaf        bool           `json:"deaf,omitempty"`
}

// MemberUpdate is used to modify a member
type MemberUpdate struct {
	ChannelID                  *snowflake.ID             `json:"channel_id,omitempty"`
	Nick                       *string                   `json:"nick,omitempty"`
	Roles                      *[]snowflake.ID           `json:"roles,omitempty"`
	Mute                       *bool                     `json:"mute,omitempty"`
	Deaf                       *bool                     `json:"deaf,omitempty"`
	Flags                      *MemberFlags              `json:"flags,omitempty"`
	CommunicationDisabledUntil *json.Nullable[time.Time] `json:"communication_disabled_until,omitempty"`
}

// CurrentMemberUpdate is used to update the current member
type CurrentMemberUpdate struct {
	Nick string `json:"nick"`
}

type MemberFlags int

const (
	MemberFlagsDidRejoin MemberFlags = 1 << iota
	MemberFlagsCompletedOnboarding
	MemberFlagsBypassesVerification
	MemberFlagsStartedOnboarding
	MemberFlagsNone MemberFlags = 0
)

// Add allows you to add multiple bits together, producing a new bit
func (f MemberFlags) Add(bits ...MemberFlags) MemberFlags {
	for _, bit := range bits {
		f |= bit
	}
	return f
}

// Remove allows you to subtract multiple bits from the first, producing a new bit
func (f MemberFlags) Remove(bits ...MemberFlags) MemberFlags {
	for _, bit := range bits {
		f &^= bit
	}
	return f
}

// Has will ensure that the bit includes all the bits entered
func (f MemberFlags) Has(bits ...MemberFlags) bool {
	for _, bit := range bits {
		if (f & bit) != bit {
			return false
		}
	}
	return true
}

// Missing will check whether the bit is missing any one of the bits
func (f MemberFlags) Missing(bits ...MemberFlags) bool {
	for _, bit := range bits {
		if (f & bit) != bit {
			return true
		}
	}
	return false
}
