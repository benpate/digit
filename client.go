package digit

import (
	"github.com/benpate/derp"
	"github.com/benpate/remote"
)

func Lookup(url string, options ...remote.Option) (Resource, error) {

	webFingerServerURLs := ParseAccount(url)
	result := NewResource(url)

	for _, webFingerServerURL := range webFingerServerURLs {

		txn := remote.Get(webFingerServerURL).
			With(options...).
			Result(&result)

		if err := txn.Send(); err == nil {
			return result, nil
		}
	}

	return result, derp.InternalError("digit.Lookup", "Unable to load resource", url, webFingerServerURLs)
}
