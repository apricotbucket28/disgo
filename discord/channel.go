package discord

import (
	"github.com/DisgoOrg/disgo/rest/route"
)

// ChannelType for interacting with discord's channels
type ChannelType int

// Channel constants
//goland:noinspection GoUnusedConst
const (
	ChannelTypeGuildText ChannelType = iota
	ChannelTypeDM
	ChannelTypeGuildVoice
	ChannelTypeGroupDM
	ChannelTypeGuildCategory
	ChannelTypeGuildNews
	ChannelTypeGuildStore
	_
	_
	_
	ChannelTypeGuildNewsThread
	ChannelTypeGuildPublicThread
	ChannelTypeGuildPrivateThread
	ChannelTypeGuildStageVoice
)

// VideoQualityMode https://discord.com/developers/docs/resources/channel#channel-object-video-quality-modes
type VideoQualityMode int

//goland:noinspection GoUnusedConst
const (
	VideoQualityModeAuto = iota + 1
	VideoQualityModeFull
)

// Channel is a generic discord channel object
type Channel struct {
	ID                         Snowflake             `json:"id"`
	Type                       ChannelType           `json:"type"`
	GuildID                    Snowflake             `json:"guild_id,omitempty"`
	Position                   int                   `json:"position,omitempty"`
	PermissionOverwrites       []PermissionOverwrite `json:"permission_overwrites"`
	Name                       string                `json:"name,omitempty"`
	Topic                      *string               `json:"topic,omitempty"`
	NSFW                       bool                  `json:"nsfw,omitempty"`
	LastMessageID              *Snowflake            `json:"last_message_id,omitempty"`
	Bitrate                    int                   `json:"bitrate,omitempty"`
	UserLimit                  int                   `json:"user_limit,omitempty"`
	RateLimitPerUser           int                   `json:"rate_limit_per_user,omitempty"`
	Recipients                 []User                `json:"recipients,omitempty"`
	Icon                       *string               `json:"icon,omitempty"`
	OwnerID                    Snowflake             `json:"owner_id,omitempty"`
	ApplicationID              Snowflake             `json:"application_id,omitempty"`
	ParentID                   *Snowflake            `json:"parent_id,omitempty"`
	LastPinTimestamp           *Time                 `json:"last_pin_timestamp,omitempty"`
	RTCRegion                  string                `json:"rtc_region"`
	VideoQualityMode           VideoQualityMode      `json:"video_quality_mode"`
	MessageCount               int                   `json:"message_count"`
	MemberCount                int                   `json:"member_count"`
	ThreadMetadata             ThreadMetadata        `json:"thread_metadata"`
	DefaultAutoArchiveDuration AutoArchiveDuration   `json:"default_auto_archive_duration"`
	InteractionPermissions     Permissions           `json:"permissions,omitempty"`
}

type ThreadMetadata struct {
	Archived            bool
	AutoArchiveDuration AutoArchiveDuration
	ArchiveTimestamp    Time
	Locked              bool
	Invitable           bool
}

type AutoArchiveDuration int

const (
	AutoArchiveDuration1h  AutoArchiveDuration = 60
	AutoArchiveDuration24h AutoArchiveDuration = 1440
	AutoArchiveDuration3d  AutoArchiveDuration = 4320
	AutoArchiveDuration1w  AutoArchiveDuration = 10080
)

type ChannelCreate struct {
	Name                 string                `json:"name"`
	Type                 ChannelType           `json:"type,omitempty"`
	Topic                string                `json:"topic,omitempty"`
	Bitrate              int                   `json:"bitrate,omitempty"`
	UserLimit            int                   `json:"user_limit,omitempty"`
	RateLimitPerUser     int                   `json:"rate_limit_per_user,omitempty"`
	Position             int                   `json:"position,omitempty"`
	PermissionOverwrites []PermissionOverwrite `json:"permission_overwrites,omitempty"`
	ParentID             Snowflake             `json:"parent_id,omitempty"`
	NSFW                 bool                  `json:"nsfw,omitempty"`
}

type DMChannelCreate struct {
	RecipientID Snowflake `json:"recipient_id"`
}

type ChannelUpdate struct {
	Name                       *string                `json:"name,omitempty"`
	Type                       *ChannelType           `json:"type,omitempty"`
	Position                   *int                   `json:"position,omitempty"`
	Topic                      *string                `json:"topic,omitempty"`
	NSFW                       *bool                  `json:"nsfw,omitempty"`
	RateLimitPerUser           *int                   `json:"rate_limit_per_user,omitempty"`
	Bitrate                    *int                   `json:"bitrate,omitempty"`
	UserLimit                  *int                   `json:"user_limit,omitempty"`
	PermissionOverwrites       *[]PermissionOverwrite `json:"permission_overwrites,omitempty"`
	ParentID                   *Snowflake             `json:"parent_id,omitempty"`
	RTCRegion                  *string                `json:"rtc_region"`
	VideoQualityMode           *VideoQualityMode      `json:"video_quality_mode"`
	DefaultAutoArchiveDuration *int                   `json:"default_auto_archive_duration"`
}

// PartialChannel contains basic info about a Channel
type PartialChannel struct {
	ID   Snowflake   `json:"id"`
	Type ChannelType `json:"type"`
	Name string      `json:"name"`
	Icon *string     `json:"icon,omitempty"`
}

// GetIconURL returns the Icon URL of this channel.
// This will be nil for every discord.ChannelType except discord.ChannelTypeGroupDM
func (c *PartialChannel) GetIconURL(size int) *string {
	return FormatAssetURL(route.ChannelIcon, c.ID, c.Icon, size)
}
