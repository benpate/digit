package digit

// RelationTypeAvatar identifies a person’s avatar and may be in any standard image format (e.g., PNG, JPEG, GIF).
const RelationTypeAvatar = "https://webfinger.net/rel/avatar/"

// RelationTypeProfile identifies the main home/profile page that a human should visit when getting info about that webfinger account. It says nothing about the content-type (or microformats), but it’s likely text/html if it’s for users.
const RelationTypeProfile = "https://webfinger.net/rel/profile-page/"

// RelationTypeActivityJSON identifies documents confirming to the Activity Streams 2.0 format. See https://www.w3.org/TR/activitystreams-core/#media-type
const RelationTypeSelf = "self"

// Additional values to consider adding: http://microformats.org/wiki/existing-rel-values
