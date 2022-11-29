package digit

// RelationTypeAvatar identifies a person’s avatar and may be in any standard image format (e.g., PNG, JPEG, GIF).
const RelationTypeAvatar = "https://webfinger.net/rel/avatar/"

// RelationTypeProfile identifies the main home/profile page that a human should visit when getting info about that webfinger account. It says nothing about the content-type (or microformats), but it’s likely text/html if it’s for users.
const RelationTypeProfile = "https://webfinger.net/rel/profile-page/"

// RelationTypeSubscribeRequest is used by Mastodon to initiate a follow/subscribe workflow.  TODO: HIGH: Figure this out.
const RelationTypeSubscribeRequest = "http://ostatus.org/schema/1.0/subscribe"

// RelationTypeActivityJSON identifies documents confirming to the Activity Streams 2.0 format. See https://www.w3.org/TR/activitystreams-core/#media-type
const RelationTypeSelf = "self"

// Additional values to consider adding: http://microformats.org/wiki/existing-rel-values
