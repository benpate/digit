package digit

import (
	"github.com/benpate/derp"
	"github.com/benpate/remote"
)

func Lookup(url string) (Resource, error) {

	webFingerServerURL, err := ParseUsername(url)

	if err != nil {
		return NewResource(url), derp.Wrap(err, "digit.Lookup", "Unable to parse resource url", url)
	}

	result := NewResource(url)

	if err := remote.Get(webFingerServerURL).Response(&result, nil).Send(); err != nil {
		return result, derp.Wrap(err, "digit.Lookup", "Error connecting to WebFinger server", webFingerServerURL)
	}

	return result, nil
}
