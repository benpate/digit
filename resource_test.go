package digit

import (
	"encoding/json"
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

func TestResourcePropertyNoAliasing(t *testing.T) {

	base := NewResource("acct:x").Property("a", "1")

	// Deriving a new Resource must not mutate the original's Properties map.
	derived := base.Property("b", "2")

	require.Equal(t, 1, len(base.Properties))
	_, ok := base.Properties["b"]
	require.False(t, ok)
	require.Equal(t, 2, len(derived.Properties))
	require.Equal(t, "2", derived.Properties["b"])
}

func TestResourcePropertyOnZeroValue(t *testing.T) {

	// Property must initialize the map when called on a zero-value Resource
	// (one not created via NewResource, whose Properties map is nil).
	resource := Resource{}.Property("a", "1")
	require.Equal(t, "1", resource.Properties["a"])
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
	require.Equal(t, "", missing.RelationType)
	require.Equal(t, "", missing.MediaType)
	require.Equal(t, "", missing.Href)

	// A miss must be detectable via IsEmpty(), consistent with LinkSet.Find.
	require.True(t, missing.IsEmpty())
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

// FuzzResourceUnmarshal feeds arbitrary bytes to the JSON decoder and, for any
// input that decodes successfully, verifies decode→encode→decode is idempotent.
func FuzzResourceUnmarshal(f *testing.F) {

	seeds := []string{
		`{}`,
		`{"subject":"acct:sarah@sky.net"}`,
		`{"subject":"acct:sarah@sky.net","aliases":["http://sky.net/sarah"]}`,
		`{"subject":"acct:sarah@sky.net","properties":{"role":"employee"}}`,
		`{"subject":"acct:sarah@sky.net","links":[{"rel":"self","type":"application/activity+json","href":"https://sky.net/users/sarah"}]}`,
		`not json`,
	}

	for _, seed := range seeds {
		f.Add(seed)
	}

	f.Fuzz(func(t *testing.T, data string) {

		var first Resource
		if err := json.Unmarshal([]byte(data), &first); err != nil {
			return // Ignore inputs that are not valid Resource JSON.
		}

		// Encode the decoded Resource, then decode and re-encode it. The encoded
		// form must be stable. (We compare the encoded bytes rather than the
		// structs because "omitempty" makes an empty slice/map indistinguishable
		// from a nil one after one round-trip.)
		encoded, err := json.Marshal(first)
		require.Nil(t, err)

		var second Resource
		require.Nil(t, json.Unmarshal(encoded, &second))

		reEncoded, err := json.Marshal(second)
		require.Nil(t, err)

		require.Equal(t, string(encoded), string(reEncoded))
	})
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

// TestResource_ChainedBuildersDoNotAlias verifies that two chains branching from
// the same Resource each get an independent copy of the Aliases and Links slices.
func TestResource_ChainedBuildersDoNotAlias(t *testing.T) {

	// Build a base Resource whose slices have spare capacity (len 3, cap 4),
	// which is the condition that lets append() write through a shared array.
	base := NewResource("acct:john@connor.com").
		Alias("one").Alias("two").Alias("three").
		Link("1", "text/html", "https://one").
		Link("2", "text/html", "https://two").
		Link("3", "text/html", "https://three")

	require.Greater(t, cap(base.Aliases), len(base.Aliases))
	require.Greater(t, cap(base.Links), len(base.Links))

	// Branch twice from the same base
	branchA := base.Alias("A").Link("a", "text/html", "https://A")
	branchB := base.Alias("B").Link("b", "text/html", "https://B")

	// Neither branch may see the other's values
	require.Equal(t, []string{"one", "two", "three", "A"}, branchA.Aliases)
	require.Equal(t, []string{"one", "two", "three", "B"}, branchB.Aliases)
	require.Equal(t, "https://A", branchA.Links[3].Href)
	require.Equal(t, "https://B", branchB.Links[3].Href)

	// ...and the base must be untouched by either branch
	require.Equal(t, []string{"one", "two", "three"}, base.Aliases)
	require.Equal(t, 3, len(base.Links))
}

// TestResource_AddLinkDoesNotAlias verifies the same independence for AddLink.
func TestResource_AddLinkDoesNotAlias(t *testing.T) {

	base := NewResource("acct:john@connor.com").
		AddLink(NewLink("1", "text/html", "https://one")).
		AddLink(NewLink("2", "text/html", "https://two")).
		AddLink(NewLink("3", "text/html", "https://three"))

	branchA := base.AddLink(NewLink("a", "text/html", "https://A"))
	branchB := base.AddLink(NewLink("b", "text/html", "https://B"))

	require.Equal(t, "https://A", branchA.Links[3].Href)
	require.Equal(t, "https://B", branchB.Links[3].Href)
	require.Equal(t, 3, len(base.Links))
}
