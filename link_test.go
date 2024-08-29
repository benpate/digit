package digit

import (
	"encoding/json"
	"testing"

	"github.com/benpate/rosetta/mapof"
	"github.com/stretchr/testify/require"
)

func TestLink(t *testing.T) {

	link := NewLink("example", "text/plain", "http://example.com").Title("Example", "und")

	require.Equal(t, 1, len(link.Titles))

	link = link.Title("", "")
	require.Equal(t, 1, len(link.Titles))

	link = link.Title("", "en-us")
	require.Equal(t, 1, len(link.Titles))

	link = link.Title("New Title", "und")
	require.Equal(t, 1, len(link.Titles))
	require.Equal(t, "New Title", link.Titles["und"])

	link = link.Title("New Title En Español", "es")
	require.Equal(t, 2, len(link.Titles))
	require.Equal(t, "New Title En Español", link.Titles["es"])

	link = link.Title("", "es")
	require.Equal(t, 1, len(link.Titles))

	link = link.Property("", "")
	require.Zero(t, len(link.Properties))
	require.Equal(t, "", link.Properties["Author"])

	link = link.Property("Author", "John Connor")
	require.Equal(t, 1, len(link.Properties))
	require.Equal(t, "John Connor", link.Properties["Author"])

	link = link.Property("Author", "")
	require.Equal(t, 0, len(link.Properties))
	require.Equal(t, "", link.Properties["Author"])
}

func TestLinkProperties(t *testing.T) {

	link := NewLink("example", "text/plain", "http://example.com")

	// Test Getters/Setters
	require.False(t, link.IsEmpty())
	require.True(t, link.SetString("rel", "new-relationship"))
	require.True(t, link.SetString("type", "new-type"))
	require.True(t, link.SetString("href", "new-href"))
	require.False(t, link.SetString("unknown", "new-unknown"))

	require.Equal(t, "new-relationship", link.GetString("rel"))
	require.Equal(t, "new-type", link.GetString("type"))
	require.Equal(t, "new-href", link.GetString("href"))
	require.Equal(t, "", link.GetString("unknown"))

	// Test GetChildren
	titles, ok := link.GetPointer("titles")
	require.True(t, ok)
	require.Equal(t, &mapof.String{}, titles)

	properties, ok := link.GetPointer("properties")
	require.True(t, ok)
	require.Equal(t, &mapof.String{}, properties)

	unknown, ok := link.GetPointer("unknown")
	require.False(t, ok)
	require.Nil(t, unknown)

	// Test empty link
	require.True(t, link.SetString("rel", ""))
	require.True(t, link.SetString("type", ""))
	require.True(t, link.SetString("href", ""))
	require.True(t, link.IsEmpty())
}

func TestMatches(t *testing.T) {
	require.True(t, NewLink("example", "text/plain", "http://example.com").Matches(NewLink("example", "text/plain", "http://new.example.com")))
	require.False(t, NewLink("not-example", "text/plain", "http://example.com").Matches(NewLink("example", "text/plain", "http://new.example.com")))
	require.False(t, NewLink("example", "not/text/plain", "http://example.com").Matches(NewLink("example", "text/plain", "http://new.example.com")))
}

func TestUnmarshalLinkWithTitles(t *testing.T) {

	link := Link{}
	linkJSON := `{"href":"example.com", "rel":"example", "type":"text/plain", "titles":{"und":"Example", "es":"Ejemplo"}}`

	err := json.Unmarshal([]byte(linkJSON), &link)

	require.Nil(t, err)
	require.Equal(t, "example.com", link.Href)
	require.Equal(t, "example", link.RelationType)
	require.Equal(t, "text/plain", link.MediaType)
	require.Equal(t, "Example", link.Titles["und"])
	require.Equal(t, "Ejemplo", link.Titles["es"])
}

func TestUnmarshalLinkWithProperties(t *testing.T) {

	link := Link{}
	linkJSON := `{"href":"example.com", "rel":"example", "type":"text/plain", "properties":{"one":"ONE", "two":"TWO"}}`

	err := json.Unmarshal([]byte(linkJSON), &link)

	require.Nil(t, err)
	require.Equal(t, "example.com", link.Href)
	require.Equal(t, "example", link.RelationType)
	require.Equal(t, "text/plain", link.MediaType)
	require.Equal(t, "ONE", link.Properties["one"])
	require.Equal(t, "TWO", link.Properties["two"])
}

func TestUnmarshalLinkWithTemplate(t *testing.T) {

	link := Link{}
	linkJSON := `{"template":"example.com?one={one}", "rel":"example", "type":"text/plain"}`

	err := json.Unmarshal([]byte(linkJSON), &link)

	require.Nil(t, err)
	require.Equal(t, "example.com?one={one}", link.Template)
	require.Equal(t, "example", link.RelationType)
	require.Equal(t, "text/plain", link.MediaType)
}

func ExampleLink() {

	// Create a new link with a link relation, mime-type, and href
	link := NewLink("http://webfinger.example/rel/profile-page", "text/html", "https://www.example.com/~bob")

	// You can also set link titles in multiple languages
	link = link.Title("en-us", "The Magical World of Steve")
	link = link.Title("fr", "Le Mondo Magique de Steve")
}

func TestCreateLinkWithSubscribeRequest(t *testing.T) {

	link := NewLink("http://ostatus.org/schema/1.0/subscribe", "application/activity+json", "https://www.example.com/subscribe?uri={uri}")

	require.Equal(t, "https://www.example.com/subscribe?uri={uri}", link.Template)
	require.Equal(t, "", link.Href)
	require.Equal(t, RelationTypeSubscribeRequest, link.RelationType)
}

func TestUnmarshalLinkWithSubscribeRequest(t *testing.T) {

	link := Link{}
	linkJSON := `{"rel":"http://ostatus.org/schema/1.0/subscribe", "template":"https://www.example.com/subscribe?uri={uri}", "type":"application/activity+json"}`

	err := json.Unmarshal([]byte(linkJSON), &link)

	require.Nil(t, err)
	require.Equal(t, "https://www.example.com/subscribe?uri={uri}", link.Template)
	require.Equal(t, "", link.Href)
	require.Equal(t, RelationTypeSubscribeRequest, link.RelationType)
	require.Equal(t, "application/activity+json", link.MediaType)
}
