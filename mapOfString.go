package digit

type MapOfString map[string]string

// GetString returns a named option as a string type.
func (mapOfString MapOfString) GetString(name string) string {
	return mapOfString[name]
}

// SetString adds a string value into the map
func (mapOfString *MapOfString) SetString(name string, value string) bool {
	(*mapOfString)[name] = value
	return true
}

// Delete removes a named option from the map
func (mapOfString *MapOfString) Delete(name string) {
	delete(*mapOfString, name)
}
