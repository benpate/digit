package digit

import (
	"github.com/benpate/rosetta/mapof"
)

// Link represents a link, or relationship, to another resource on the Internet.
// https://datatracker.ietf.org/doc/html/rfc7033#section-4.4.4
type Link struct {
	RelationType string       `json:"rel,omitempty"        bson:"rel,omitempty"`        // Either a URI or a registered relation type (RFC5988)
	MediaType    string       `json:"type,omitempty"       bson:"type,omitempty"`       // Media Type of the target resource (RFC 3986)
	Href         string       `json:"href,omitempty"       bson:"href,omitempty"`       // URI of the target resource
	Titles       mapof.String `json:"titles,omitempty"     bson:"titles,omitempty"`     // Map keys are either language tag (or the string "und"), values are the title of this object in that language.  If the language is unknown or unspecified, then the name is "und".
	Properties   mapof.String `json:"properties,omitempty" bson:"properties,omitempty"` // Zero or more name/value pairs whose names are URIs and whose values are strings.  Properties are used to convey additional information about the link relationship.
	Template     string       `json:"template,omitempty"   bson:"template,omitempty"`   // Non-standard URI template for the target resource (added by oStatus remote follows)
}

// NewLink returns a fully initialized Link object.
func NewLink(relationType string, mediaType string, href string) Link {
	result := Link{
		RelationType: relationType,
		MediaType:    mediaType,
		Href:         href,
		Titles:       mapof.NewString(),
		Properties:   mapof.NewString(),
	}

	// Special case for oStatus subscribe requests
	if relationType == RelationTypeSubscribeRequest {
		result.Template = result.Href
		result.Href = ""
	}

	return result
}

// IsEmpty returns TRUE if the Link object has no values set.
func (link Link) IsEmpty() bool {
	return link.RelationType == "" && link.MediaType == "" && link.Href == ""
}

// NotEmpty returns TRUE if the Link object has at least one value set.
func (link Link) NotEmpty() bool {
	return !link.IsEmpty()
}

// Title populates a title value for the Link.
func (link Link) Title(title string, language string) Link {

	if language == "" {
		return link
	}

	if title == "" {
		delete(link.Titles, language)
		return link
	}

	link.Titles[language] = title
	return link
}

// Property populates a property of the link.  Name must be a URI (called a property identifier) and value must be a string.
func (link Link) Property(name string, value string) Link {

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

// Matches returns TRUE if the "otherLink" has the same type and rel as this link
func (link Link) Matches(otherLink Link) bool {
	return (link.MediaType == otherLink.MediaType) && (link.RelationType == otherLink.RelationType)
}

// GetString returns string values of this Link object
func (link Link) GetString(name string) string {
	result, _ := link.GetStringOK(name)
	return result
}

// GetStringOK returns string values of this Link object
func (link Link) GetStringOK(name string) (string, bool) {
	switch name {

	case "rel":
		return link.RelationType, true

	case "type":
		return link.MediaType, true

	case "href":
		return link.Href, true

	case "template":
		return link.Template, true

	default:
		return "", false
	}
}

// SetString sets string values of this Link object
func (link *Link) SetString(name string, value string) bool {
	switch name {

	case "rel":
		link.RelationType = value
		return true

	case "type":
		link.MediaType = value
		return true

	case "href":
		link.Href = value
		return true

	case "template":
		link.Template = value
		return true

	default:
		return false
	}
}

// GetPointer returns a reference to struct values of this Link object
func (link *Link) GetPointer(name string) (interface{}, bool) {
	switch name {

	case "titles":
		return &link.Titles, true

	case "properties":
		return &link.Properties, true

	default:
		return nil, false
	}
}
