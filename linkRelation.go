package digit

// LinkRelation represents a relationship as definedin RFC5988
type LinkRelation struct {
	RelationName string
	Description  string
	Reference    string
}

// LinkRelations returns an array of all link relations defined in RFC5988
func LinkRelations() []LinkRelation {

	return []LinkRelation{
		{
			RelationName: "alternate",
			Description:  "Designates a substitute for the link's context.",
			Reference:    "[W3C.REC-html401-19991224]",
		}, {
			RelationName: "appendix",
			Description:  "Refers to an appendix.",
			Reference:    "[W3C.REC-html401-19991224]",
		}, {
			RelationName: "bookmark",
			Description:  "Refers to a bookmark or entry point.",
			Reference:    "[W3C.REC-html401-19991224]",
		}, {
			RelationName: "chapter",
			Description:  "Refers to a chapter in a collection of resources.",
			Reference:    "[W3C.REC-html401-19991224]",
		}, {
			RelationName: "contents",
			Description:  "Refers to a table of contents.",
			Reference:    "[W3C.REC-html401-19991224]",
		}, {
			RelationName: "copyright",
			Description:  "Refers to a copyright statement that applies to the link's context.",
			Reference:    "[W3C.REC-html401-19991224]",
		}, {
			RelationName: "current",
			Description:  "Refers to a resource containing the most recent item(s) in a collection of resources.",
			Reference:    "[RFC5005]",
		}, {
			RelationName: "describedby",
			Description:  "Refers to a resource providing information about the link's context.",
			Reference:    "http://www.w3.org/TR/powder-dr/#assoc-linking>",
		}, {
			RelationName: "edit",
			Description:  "Refers to a resource that can be used to edit the link's context.",
			Reference:    "[RFC5023]",
		}, {
			RelationName: "edit-media",
			Description:  "Refers to a resource that can be used to edit media associated with the link's context.",
			Reference:    "[RFC5023]",
		}, {
			RelationName: "enclosure",
			Description:  "Identifies a related resource that is potentially large and might require special handling.",
			Reference:    "[RFC4287]",
		}, {
			RelationName: "first",
			Description:  "An IRI that refers to the furthest preceding resource in a series of resources.",
			Reference:    "[RFC5988]",
		}, {
			RelationName: "glossary",
			Description:  "Refers to a glossary of terms.",
			Reference:    "[W3C.REC-html401-19991224]",
		}, {
			RelationName: "help",
			Description:  "Refers to a resource offering help (more information, links to other sources information, etc.)",
			Reference:    "[W3C.REC-html401-19991224]",
		}, {
			RelationName: "hub",
			Description:  "Refers to a hub that enables registration for notification of updates to the context.",
			Reference:    "<http://pubsubhubbub.googlecode.com/> <http://pubsubhubbub.googlecode.com/svn/trunk/pubsubhubbub-core-0.3.html>",
		}, {
			RelationName: "index",
			Description:  "Refers to an index.",
			Reference:    "[W3C.REC-html401-19991224]",
		}, {
			RelationName: "last",
			Description:  "An IRI that refers to the furthest following resource in a series of resources.",
			Reference:    "[RFC5988]",
		}, {
			RelationName: "latest-version",
			Description:  "Points to a resource containing the latest (e.g., current) version of the context.",
			Reference:    "[RFC5829]",
		}, {
			RelationName: "license",
			Description:  "Refers to a license associated with the link's context.",
			Reference:    "[RFC4946]",
		}, {
			RelationName: "next",
			Description:  "Refers to the next resource in a ordered series of resources.",
			Reference:    "[W3C.REC-html401-19991224]",
		}, {
			RelationName: "next-archive",
			Description:  "Refers to the immediately following archive resource.",
			Reference:    "[RFC5005]",
		}, {
			RelationName: "payment",
			Description:  "indicates a resource where payment is accepted.",
			Reference:    "[RFC5988]",
		}, {
			RelationName: "prev",
			Description:  "Refers to the previous resource in an ordered series of resources.  Synonym for \"previous\".",
			Reference:    "[W3C.REC-html401-19991224]",
		}, {
			RelationName: "predecessor-version",
			Description:  "Points to a resource containing the predecessor version in the version history.",
			Reference:    "[RFC5829]",
		}, {
			RelationName: "previous",
			Description:  "Refers to the previous resource in an ordered series of resources.  Synonym for \"prev\".",
			Reference:    "[W3C.REC-html401-19991224]",
		}, {
			RelationName: "prev-archive",
			Description:  "Refers to the immediately preceding archive resource.",
			Reference:    "[RFC5005]",
		}, {
			RelationName: "related",
			Description:  "Identifies a related resource.",
			Reference:    "[RFC4287]",
		}, {
			RelationName: "replies",
			Description:  "Identifies a resource that is a reply to the context of the link.",
			Reference:    "[RFC4685]",
		}, {
			RelationName: "section",
			Description:  "Refers to a section in a collection of resources.",
			Reference:    "[W3C.REC-html401-19991224]",
		}, {
			RelationName: "self",
			Description:  "Conveys an identifier for the link's context.",
			Reference:    "[RFC4287]",
		}, {
			RelationName: "service",
			Description:  "Indicates a URI that can be used to retrieve a service document.",
			Reference:    "[RFC5023]",
		}, {
			RelationName: "start",
			Description:  "Refers to the first resource in a collection of resources.",
			Reference:    "[W3C.REC-html401-19991224]",
		}, {
			RelationName: "stylesheet",
			Description:  "Refers to an external style sheet.",
			Reference:    "[W3C.REC-html401-19991224]",
		}, {
			RelationName: "subsection",
			Description:  "Refers to a resource serving as a subsection in a collection of resources.",
			Reference:    "[W3C.REC-html401-19991224]",
		}, {
			RelationName: "successor-version",
			Description:  "Points to a resource containing the successor version in the version history.",
			Reference:    "[RFC5829]",
		}, {
			RelationName: "up",
			Description:  "Refers to a parent document in a hierarchy of documents.",
			Reference:    "[RFC5988]",
		}, {
			RelationName: "version-history",
			Description:  "points to a resource containing the version history for the context.",
			Reference:    "[RFC5829]",
		}, {
			RelationName: "via",
			Description:  "Identifies a resource that is the source of the information in the link's context.",
			Reference:    "[RFC4287]",
		}, {
			RelationName: "working-copy",
			Description:  "Points to a working copy for this resource.",
			Reference:    "[RFC5829]",
		}, {
			RelationName: "working-copy-of",
			Description:  "Points to the versioned resource from which this working copy was obtained.",
			Reference:    "[RFC5829]",
		},
	}
}
