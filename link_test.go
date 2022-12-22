package digit

import (
	"testing"

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
	titles, ok := link.GetChild("titles")
	require.True(t, ok)
	require.Equal(t, make(MapOfString), titles)

	properties, ok := link.GetChild("properties")
	require.True(t, ok)
	require.Equal(t, make(MapOfString), properties)

	unknown, ok := link.GetChild("unknown")
	require.False(t, ok)
	require.Nil(t, unknown)

	// Test empty link
	require.True(t, link.SetString("rel", ""))
	require.True(t, link.SetString("type", ""))
	require.True(t, link.SetString("href", ""))
	require.True(t, link.IsEmpty())
}

func ExampleLink() {

	// Create a new link with a link relation, mime-type, and href
	link := NewLink("http://webfinger.example/rel/profile-page", "text/html", "https://www.example.com/~bob")

	// You can also set link titles in multiple languages
	link = link.Title("en-us", "The Magical World of Steve")
	link = link.Title("fr", "Le Mondo Magique de Steve")
}
