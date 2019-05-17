package digit

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLink(t *testing.T) {

	link := NewLink("example", "text/plain", "http://example.com", "Example")

	assert.Equal(t, 1, len(link.Titles))

	link.Title("", "")
	assert.Equal(t, 1, len(link.Titles))

	link.Title("en-us", "")
	assert.Equal(t, 1, len(link.Titles))

	link.Title("und", "New Title")
	assert.Equal(t, 1, len(link.Titles))
	assert.Equal(t, "New Title", link.Titles["und"])

	link.Title("es", "New Title En Español")
	assert.Equal(t, 2, len(link.Titles))
	assert.Equal(t, "New Title En Español", link.Titles["es"])

	link.Title("es", "")
	assert.Equal(t, 1, len(link.Titles))

	link.Property("", "")
	assert.Zero(t, len(link.Properties))
	assert.Equal(t, "", link.Properties["Author"])

	link.Property("Author", "John Connor")
	assert.Equal(t, 1, len(link.Properties))
	assert.Equal(t, "John Connor", link.Properties["Author"])

	link.Property("Author", "")
	assert.Equal(t, 0, len(link.Properties))
	assert.Equal(t, "", link.Properties["Author"])
}

func ExampleLink() {

	link := NewLink("http://webfinger.example/rel/profile-page", "text/html", "https://www.example.com/~bob", "Bob Smith")

	fmt.Print(link.Titles["und"]) // Default language is "und" for undetermined.
	// Output: Bob Smith

	// You can also set a specific language on the link object itself
	link.Title("en-us", "The Magical World of Steve")
	link.Title("fr", "Le Mondo Magique de Steve")
}
