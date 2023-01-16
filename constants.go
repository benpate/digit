package digit

// RelationTypeAvatar identifies a person’s avatar and may be in any standard image
// format (e.g., PNG, JPEG, GIF).
const RelationTypeAvatar = "https://webfinger.net/rel/avatar/"

// RelationTypeProfile identifies the main home/profile page that a human should
// visit when getting info about that webfinger account. It says nothing about
// the content-type (or microformats), but it’s likely text/html if it’s for users.
const RelationTypeProfile = "https://webfinger.net/rel/profile-page/"

// RelationTypeSubscribeRequest is used by Mastodon to initiate a remote follow.
// See: https://www.hughrundle.net/how-to-implement-remote-following-for-your-activitypub-project/
// See: http://ostatus.github.io/spec/OStatus%201.0%20Draft%202.html#anchor10
const RelationTypeSubscribeRequest = "http://ostatus.org/schema/1.0/subscribe"

// RelationTypeActivityJSON identifies documents confirming to the Activity Streams 2.0 format.
// See: https://www.w3.org/TR/activitystreams-core/#media-type
const RelationTypeSelf = "self"

// RelationTypeHub is used by PubSubHubbub to identify the hub for a given resource.
// See: https://www.w3.org/TR/websub/#discovery
// See: https://pubsubhubbub.github.io/PubSubHubbub/pubsubhubbub-core-0.4.html#anchor4
const RelationTypeHub = "hub"

// Additional values to consider adding: http://microformats.org/wiki/existing-rel-values
