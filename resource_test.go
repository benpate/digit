package digit

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestResource(t *testing.T) {

	resource := NewResource("acct:sarah@sky.net").
		Alias("http://sky.net/sarah").
		Alias("http://other.website.com/sarah-connor").
		Property("http://sky.net/ns/role", "employee").
		Link("http://webfinger.example/rel/profile-page", "text/html", "https://sky.net/sarah", "Sarah Connor")

	// Verify that all properties have been populated correctly.
	assert.Equal(t, resource.Subject, "acct:sarah@sky.net")
	assert.Equal(t, resource.Aliases[0], "http://sky.net/sarah")
	assert.Equal(t, resource.Aliases[1], "http://other.website.com/sarah-connor")
	assert.Equal(t, resource.Properties["http://sky.net/ns/role"], "employee")
	assert.Equal(t, resource.Links[0].RelationType, "http://webfinger.example/rel/profile-page")
	assert.Equal(t, resource.Links[0].MediaType, "text/html")
	assert.Equal(t, resource.Links[0].Href, "https://sky.net/sarah")

	link := NewLink("http://example.com", "text/html", "https://connor.com/john", "John Connor")

	resource.AppendLink(*link)

	assert.Equal(t, 2, len(resource.Links))
	assert.Equal(t, "text/html", resource.Links[1].MediaType)
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
