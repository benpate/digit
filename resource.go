package digit

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

// Alias adds an alias (additional URI) to this Resource.  It returns a pointer to the Resource so that calls can be chained.
func (resource Resource) Alias(URI string) Resource {
	resource.Aliases = append(resource.Aliases, URI)
	return resource
}

// Property adds a property to this Resource.  It returns a pointer to the Resource so that calls can be chained.
func (resource Resource) Property(name string, value string) Resource {
	resource.Properties[name] = value
	return resource
}

// Link adds a link to this Resource.  It returns a pointer to the Resource so that calls can be chained.
func (resource Resource) Link(relationType string, mediaType string, href string) Resource {
	resource.Links = append(resource.Links, NewLink(relationType, mediaType, href))
	return resource
}

// AddLink adds a link to this Resource.  It returns a pointer to the Resource so that calls can be chained.
func (resource Resource) AddLink(link Link) Resource {
	resource.Links = append(resource.Links, link)
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

	return NewLink(relationType, "", "")
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
