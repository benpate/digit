package digit

// Link represents a link, or relationship, to another resource on the Internet.
type Link struct {
	RelationType string            `json:"rel"`                  // Either a URI or a registered relation type (RFC5988)
	MediaType    string            `json:"type,omitempty"`       // Media Type of the target resource (RFC 3986)
	Href         string            `json:"href,omitempty"`       // URI of the target resource
	Titles       map[string]string `json:"titles,omitempty"`     // Map keys are either language tag (or the string "und"), values are the title of this object in that language.  If the language is unknown or unspecified, then the name is "und".
	Properties   map[string]string `json:"properties,omitempty"` // Zero or more name/value pairs whose names are URIs and whose values are strings.  properties are used to convey additional information about the link relationship.
}

// NewLink returns a fully initialized Link object.
func NewLink(relationType string, mediaType string, href string, title string) *Link {
	result := &Link{
		RelationType: relationType,
		MediaType:    mediaType,
		Href:         href,
		Titles:       map[string]string{},
		Properties:   map[string]string{},
	}

	result.Title("und", title)
	return result
}

// Title populates a title value for the Link.
func (link *Link) Title(name string, value string) *Link {

	if name == "" {
		return link
	}

	if value == "" {
		delete(link.Titles, name)
		return link
	}

	link.Titles[name] = value
	return link
}

// Property populates a property of the link.  Name must be a URI (called a property identifier) and value must be a string.
func (link *Link) Property(name string, value string) *Link {

	if name == "" {
		return link
	}

	if value == "" {
		delete(link.Properties, name)
		return link
	}

	link.Properties[name] = value
	return link
}
