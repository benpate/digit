package digit

type LinkSet []Link

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

// Remove removes all items from the set that match the given link (having identical "rel"
// and "type" properties)
func (set *LinkSet) Remove(link Link) {
	for index, target := range *set {
		if link.Matches(target) {
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

// RemoveBy removes the first link with a property that matches the given value
func (set *LinkSet) RemoveBy(name string, value string) {
	for index, link := range *set {
		if link.GetString(name) == value {
			(*set) = append((*set)[:index], (*set)[index+1:]...)
			break
		}
	}
}

// ApplyBy searches for a matching link, updates it if found, and appends it if not
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
