package digit

// Resource defines a single resource (such as a user or web page) that is being queried using the WebFinger protocol
type Resource struct {
	Subject    string            `json:"subject"`              // REQUIRED: URI that identifies the entity.
	Aliases    []string          `json:"aliases,omitempty"`    // Zero or more  URI strings that identify the same entity as the "subject" URI
	Properties map[string]string `json:"properties,omitempty"` // Zero of more name/value pairs whose names are URIs and whose values are strings.  Properties are used to convey additional information about the subject of the JRD.
	Links      []Link            `json:"links,omitempty"`      // Links to resources that are related or connected to this one.
}

// NewResource returns a fully initialized resource.  The "subject" is a URI that identifies the entity.
func NewResource(subject string) *Resource {
	return &Resource{
		Subject:    subject,
		Aliases:    []string{},
		Properties: map[string]string{},
		Links:      []Link{},
	}
}

// Alias adds an alias (additional URI) to this Resource.  It returns a pointer to the Resource so that calls can be chained.
func (resource *Resource) Alias(URI string) *Resource {
	resource.Aliases = append(resource.Aliases, URI)
	return resource
}

// Property adds a property to this Resource.  It returns a pointer to the Resource so that calls can be chained.
func (resource *Resource) Property(name string, value string) *Resource {
	resource.Properties[name] = value
	return resource
}

// Link adds a link to this Resource.  It returns a pointer to the Resource so that calls can be chained.
func (resource *Resource) Link(relationType string, mediaType string, href string, title string) *Resource {
	link := NewLink(relationType, mediaType, href, title)
	resource.Links = append(resource.Links, *link)
	return resource
}

// AppendLink adds a link to this Resource.  It returns a pointer to the Resource so that calls can be chained.
func (resource *Resource) AppendLink(link Link) *Resource {
	resource.Links = append(resource.Links, link)
	return resource
}
