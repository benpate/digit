package digit

// LinkSet is a collection of Links with helper methods for finding, applying,
// and removing links by their "rel" and "type" properties.
type LinkSet []Link

// NewLinkSet returns an empty LinkSet with the given pre-allocated capacity.
func NewLinkSet(capacity int) LinkSet {
	return make(LinkSet, 0, capacity)
}

// Append adds a new link into the set without inspecting its contents
func (set *LinkSet) Append(links ...Link) {
	(*set) = append(*set, links...)
}

// Find returns the first link that matches the provided link (having identical
// "rel" and "type" properties)
func (set LinkSet) Find(link Link) Link {
	for _, target := range set {
		if target.Matches(link) {
			return target
		}
	}

	return NewLink("", "", "")
}

// Apply searches for the first link that matches (with identical "rel" and "type"
// properties) the given link. If found, then the first matching item is updated.
// If not, then a new link is inserted
func (set *LinkSet) Apply(link Link) {
	for index, target := range *set {
		if link.Matches(target) {
			(*set)[index] = link
			return
		}
	}
	*set = append(*set, link)
}

// Remove deletes ALL items that match the given link (identical "rel" and "type").
// It mutates the set in place, so callers must not rely on slices that alias the
// same backing array.
func (set *LinkSet) Remove(link Link) {
	// Iterate backwards so that removing an item does not shift the indexes
	// of the items we have not yet inspected.
	for index := len(*set) - 1; index >= 0; index-- {
		if link.Matches((*set)[index]) {
			*set = append((*set)[:index], (*set)[index+1:]...)
		}
	}
}

// FindBy returns the first link with a property that matches the given value
func (set LinkSet) FindBy(name string, value string) Link {
	for _, link := range set {
		if link.GetString(name) == value {
			return link
		}
	}

	return NewLink("", "", "")
}

// RemoveBy deletes the FIRST link whose named property matches the given value.
// Unlike Remove, it stops after one match. It mutates the set in place.
func (set *LinkSet) RemoveBy(name string, value string) {
	for index, link := range *set {
		if link.GetString(name) == value {
			(*set) = append((*set)[:index], (*set)[index+1:]...)
			break
		}
	}
}

// ApplyBy updates the first link whose named property matches the given link's,
// or appends the link if none match. It mutates the set in place.
func (set *LinkSet) ApplyBy(name string, link Link) {
	keyID := link.GetString(name)
	for index, target := range *set {
		if target.GetString(name) == keyID {
			(*set)[index] = link
			return
		}
	}

	(*set) = append(*set, link)
}
