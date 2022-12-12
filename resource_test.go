package digit

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestResource(t *testing.T) {

	resource := NewResource("acct:sarah@sky.net").
		Alias("http://sky.net/sarah").
		Alias("http://other.website.com/sarah-connor").
		Property("http://sky.net/ns/role", "employee").
		Link(RelationTypeProfile, "text/html", "https://sky.net/sarah")

	// Verify that all properties have been populated correctly.
	require.Equal(t, resource.Subject, "acct:sarah@sky.net")
	require.Equal(t, resource.Aliases[0], "http://sky.net/sarah")
	require.Equal(t, resource.Aliases[1], "http://other.website.com/sarah-connor")
	require.Equal(t, resource.Properties["http://sky.net/ns/role"], "employee")
	require.Equal(t, resource.Links[0].RelationType, RelationTypeProfile)
	require.Equal(t, resource.Links[0].MediaType, "text/html")
	require.Equal(t, resource.Links[0].Href, "https://sky.net/sarah")

	link := NewLink(RelationTypeProfile, "text/html", "https://john.connor.com").Title("John Connor", "en")

	resource = resource.AddLink(link)

	require.Equal(t, 2, len(resource.Links))
	require.Equal(t, "text/html", resource.Links[1].MediaType)

}

func TestFindLink(t *testing.T) {

	resource := NewResource("acct:sarah@sky.net").
		Link(RelationTypeAvatar, "img/webp", "https://sara.sky.net/profile.webp").
		Link(RelationTypeProfile, "text/html", "https://sara.sky.net/profile").
		Link(RelationTypeSelf, "application/activity+json", "https://sara.sky.net/activity.json")

	avatar := resource.FindLink(RelationTypeAvatar)
	require.Equal(t, "img/webp", avatar.MediaType)
	require.Equal(t, "https://sara.sky.net/profile.webp", avatar.Href)

	profile := resource.FindLink(RelationTypeProfile)
	require.Equal(t, "text/html", profile.MediaType)
	require.Equal(t, "https://sara.sky.net/profile", profile.Href)

	activity := resource.FindLink(RelationTypeSelf)
	require.Equal(t, "application/activity+json", activity.MediaType)
	require.Equal(t, "https://sara.sky.net/activity.json", activity.Href)

	missing := resource.FindLink("missing-type")
	require.Equal(t, "", missing.MediaType)
	require.Equal(t, "", missing.Href)
}

func TestResource_FilterLinks(t *testing.T) {

	resource := NewResource("acct:sarah@sky.net").
		Link(RelationTypeAvatar, "img/webp", "https://sara.sky.net/profile.webp").
		Link(RelationTypeProfile, "text/html", "https://sara.sky.net/profile").
		Link(RelationTypeSelf, "application/activity+json", "https://sara.sky.net/activity.json")

	require.Equal(t, 3, len(resource.Links))

	// FilterLinks with empty string is a no-op.
	resource.FilterLinks("")
	require.Equal(t, 3, len(resource.Links))

	// FilterLinks returns all matching values.
	resource.FilterLinks(RelationTypeAvatar)

	require.Equal(t, 1, len(resource.Links))
	require.Equal(t, RelationTypeAvatar, resource.Links[0].RelationType)
	require.Equal(t, "img/webp", resource.Links[0].MediaType)
	require.Equal(t, "https://sara.sky.net/profile.webp", resource.Links[0].Href)

}

func TestResource_FilterLinks_NonMatching(t *testing.T) {

	resource := NewResource("acct:sarah@sky.net").
		Link(RelationTypeAvatar, "img/webp", "https://sara.sky.net/profile.webp").
		Link(RelationTypeProfile, "text/html", "https://sara.sky.net/profile").
		Link(RelationTypeSelf, "application/activity+json", "https://sara.sky.net/activity.json")

	// FilterLinks with non-matching value returns an empty list
	resource.FilterLinks("Non-Matching-Value")
	require.Equal(t, 0, len(resource.Links))
}

func ExampleResource() {

	// Create and populate the resource object.
	resource := NewResource("acct:sarah@sky.net").
		Alias("http://sky.net/sarah").
		Alias("http://linkedin.com/in/sarah-connor").
		Property("http://sky.net/ns/role", "employee")

	fmt.Print(resource.Subject)
	// Output: acct:sarah@sky.net
}
