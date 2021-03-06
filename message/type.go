package message

// These are the message types.
const (
	// This is a type for testing purposes. It is not sent by slack.
	TypeHelloWorld = "hello_world"

	TypeHello        = "hello"
	TypeError        = "error"
	TypeReconnectURL = "reconnect_url"
	TypeMessage      = "message"
	TypeTyping       = "user_typing"

	TypeChannelMarked         = "channel_marked"
	TypeChannelCreated        = "channel_created"
	TypeChannelJoined         = "channel_joined"
	TypeChannelLeft           = "channel_left"
	TypeChannelDeleted        = "channel_deleted"
	TypeChannelRename         = "channel_rename"
	TypeChannelArchive        = "channel_archive"
	TypeChannelUnarchive      = "channel_unarchive"
	TypeChannelHistoryChanged = "channel_history_changed"

	TypeDNDUpdated     = "dnd_updated"
	TypeDNDUpdatedUser = "dnd_updated_user"

	TypeIMCreated        = "im_created"
	TypeIMOpen           = "im_open"
	TypeIMClose          = "im_close"
	TypeIMMarked         = "im_marked"
	TypeIMHistoryChanged = "im_history_changed"

	TypeGroupJoined         = "group_joined"
	TypeGroupLeft           = "group_left"
	TypeGroupOpen           = "group_open"
	TypeGroupClose          = "group_close"
	TypeGroupArchive        = "group_archive"
	TypeGroupUnarchive      = "group_unarchive"
	TypeGroupRename         = "group_rename"
	TypeGroupMarked         = "group_marked"
	TypeGroupHistoryChanged = "group_history_changed"

	TypeFileCreated        = "file_created"
	TypeFileShared         = "file_shared"
	TypeFileUnshared       = "file_unshared"
	TypeFilePublic         = "file_public"
	TypeFilePrivate        = "file_private"
	TypeFileChange         = "file_change"
	TypeFileDeleted        = "file_deleted"
	TypeFileCommentAdded   = "file_comment_added"
	TypeFileCommentEdited  = "file_comment_edited"
	TypeFileCommentDeleted = "file_comment_deleted"

	TypePinAdded   = "pin_added"
	TypePinRemoved = "pin_removed"

	TypePresenceChange       = "presence_change"
	TypeManualPresenceChange = "manual_presence_change"

	TypePrefChange      = "pref_change"
	TypeUserChange      = "user_change"
	TypeTeamJoin        = "team_join"
	TypeStarAdded       = "star_added"
	TypeStarRemoved     = "star_removed"
	TypeReactionAdded   = "reaction_added"
	TypeReactionRemoved = "reaction_removed"
	TypeEmojiChanged    = "emoji_changed"
	TypeCommandsChanged = "commands_changed"

	TypeTeamPlanChange       = "team_plan_change"
	TypeTeamPrefChange       = "team_pref_change"
	TypeTeamRename           = "team_rename"
	TypeTeamDomainChange     = "team_domain_change"
	TypeEmailDomainChanged   = "email_domain_changed"
	TypeTeamProfileChange    = "team_profile_change"
	TypeTeamProfileDelete    = "team_profile_delete"
	TypeTeamProfileReorder   = "team_profile_reorder"
	TypeTeamMigrationStarted = "team_migration_started"

	TypeBotAdded        = "bot_added"
	TypeBotChanged      = "bot_changed"
	TypeAccountsChanged = "accounts_changed"

	TypeSubteamCreated     = "subteam_created"
	TypeSubteamUpdated     = "subteam_updated"
	TypeSubteamSelfAdded   = "subteam_self_added"
	TypeSubteamSelfRemoved = "subteam_self_removed"
)
