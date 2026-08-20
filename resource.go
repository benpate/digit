package digit

import (
	"maps"
	"slices"
)

// Resource defines a single resource (such as a user or web page) that is being queried using the WebFinger protocol
// https://datatracker.ietf.org/doc/html/rfc7033#section-4.4
type Resource struct {
	Subject    string            `json:"subject"`              // REQUIRED: URI that identifies the entity.
	Aliases    []string          `json:"aliases,omitempty"`    // Zero or more  URI strings that identify the same entity as the "subject" URI
	Properties map[string]string `json:"properties,omitempty"` // Zero of more name/value pairs whose names are URIs and whose values are strings.  Properties are used to convey additional information about the subject of the JRD.
	Links      []Link            `json:"links,omitempty"`      // Links to resources that are related or connected to this one.
}

// NewResource returns a fully initialized resource.  The "subject" is a URI that identifies the entity.
func NewResource(subject string) Resource {
	return Resource{
		Subject:    subject,
		Aliases:    make([]string, 0),
		Properties: make(map[string]string),
		Links:      make([]Link, 0),
	}
}

// Alias adds an alias (additional URI) to this Resource.  It returns the updated Resource so that calls can be chained.
func (resource Resource) Alias(URI string) Resource {
	// Clip the slice so that append always allocates. Without this, two chained
	// calls branching from the same Resource would write into a shared backing
	// array and clobber each other.
	resource.Aliases = append(slices.Clip(resource.Aliases), URI)
	return resource
}

// Property adds a property to this Resource.  It returns the updated Resource so that calls can be chained.
func (resource Resource) Property(name string, value string) Resource {
	// Copy the map so chained calls don't mutate the original Resource's data.
	resource.Properties = maps.Clone(resource.Properties)

	if resource.Properties == nil {
		resource.Properties = make(map[string]string)
	}

	resource.Properties[name] = value
	return resource
}

// Link adds a link to this Resource.  It returns the updated Resource so that calls can be chained.
func (resource Resource) Link(relationType string, mediaType string, href string) Resource {
	return resource.AddLink(NewLink(relationType, mediaType, href))
}

// AddLink adds a link to this Resource.  It returns the updated Resource so that calls can be chained.
func (resource Resource) AddLink(link Link) Resource {
	// Clip the slice so that append always allocates, keeping chained calls
	// that branch from the same Resource from clobbering each other.
	resource.Links = append(slices.Clip(resource.Links), link)
	return resource
}

// FindLink searches links to find one that matches the provided relationType.
// If none is found, then an empty link is returned
func (resource Resource) FindLink(relationType string) Link {

	for _, link := range resource.Links {
		if link.RelationType == relationType {
			return link
		}
	}

	// Return a fully-empty link so that callers can detect a miss via IsEmpty().
	return NewLink("", "", "")
}

// FilterLinks updates the resource to only include links that match the provided relationType.
// If the provided relationType is empty, then no change is performed.
func (resource *Resource) FilterLinks(relationType string) {

	if relationType == "" {
		return
	}

	filteredLinks := make([]Link, 0, len(resource.Links))

	for _, link := range resource.Links {
		if link.RelationType == relationType {
			filteredLinks = append(filteredLinks, link)
		}
	}

	resource.Links = filteredLinks
}
