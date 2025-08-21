package digit

import (
	"testing"

	"github.com/benpate/domain"
	"github.com/stretchr/testify/require"
)

func TestParseURL(t *testing.T) {

	var webFingerURLs []string

	// Test URL
	webFingerURLs = ParseAccount("https://connor.com/john")
	require.Equal(t, 1, len(webFingerURLs))
	require.Equal(t, "https://connor.com/.well-known/webfinger?resource=acct:https://connor.com/john", webFingerURLs[0])

	// Test Fediverse @URL
	webFingerURLs = ParseAccount("https://connor.com/@john")
	require.Equal(t, 1, len(webFingerURLs))
	require.Equal(t, "https://connor.com/.well-known/webfinger?resource=acct:https://connor.com/@john", webFingerURLs[0])

	// Test simple email address
	webFingerURLs = ParseAccount("john@connor.com")
	require.Equal(t, 1, len(webFingerURLs))
	require.Equal(t, "https://connor.com/.well-known/webfinger?resource=acct:john@connor.com", webFingerURLs[0])

	// Test Fediverse address
	webFingerURLs = ParseAccount("@sarah@sky.net")
	require.Equal(t, 1, len(webFingerURLs))
	require.Equal(t, "https://sky.net/.well-known/webfinger?resource=acct:sarah@sky.net", webFingerURLs[0])

	// Test Localhost addresses
	webFingerURLs = ParseAccount("http://localhost/john")
	require.Equal(t, 1, len(webFingerURLs))
	require.Equal(t, "http://localhost/.well-known/webfinger?resource=acct:http://localhost/john", webFingerURLs[0])

}

func TestParseURL_Error(t *testing.T) {

	var webFingerURLs []string

	// Test previously failed URL
	webFingerURLs = ParseAccount("@first-group@127.0.0.1")
	require.Equal(t, 2, len(webFingerURLs))
	require.Equal(t, "https://127.0.0.1/.well-known/webfinger?resource=acct:first-group@127.0.0.1", webFingerURLs[0])
	require.Equal(t, "http://127.0.0.1/.well-known/webfinger?resource=acct:first-group@127.0.0.1", webFingerURLs[1])

	// Test previously failed URL
	webFingerURLs = ParseAccount("first-group@127.0.0.1")
	require.Equal(t, 2, len(webFingerURLs))
	require.Equal(t, "https://127.0.0.1/.well-known/webfinger?resource=acct:first-group@127.0.0.1", webFingerURLs[0])
	require.Equal(t, "http://127.0.0.1/.well-known/webfinger?resource=acct:first-group@127.0.0.1", webFingerURLs[1])
}

func TestParseURL_ActivityPub(t *testing.T) {

	// Test Fediverse localhost address
	webFingerURLs := ParseAccount("@sarah@localhost:3000")
	require.Equal(t, 2, len(webFingerURLs))
	require.Equal(t, "https://localhost:3000/.well-known/webfinger?resource=acct:sarah@localhost:3000", webFingerURLs[0])
	require.Equal(t, "http://localhost:3000/.well-known/webfinger?resource=acct:sarah@localhost:3000", webFingerURLs[1])
}

func TestParseURL_WeirdStuff(t *testing.T) {

	// Test URL with port
	webFingerURLs := ParseAccount("https://connor.com:8080/john")
	require.Equal(t, 1, len(webFingerURLs))
	require.Equal(t, "https://connor.com:8080/.well-known/webfinger?resource=acct:https://connor.com:8080/john", webFingerURLs[0])
}

func TestParseURL_WeirdStuff2(t *testing.T) {

	// This is actually a valid URL
	webFingerURLs := ParseAccount("https://@john")
	require.Equal(t, 1, len(webFingerURLs))
}

func TestParseURL_WeirdStuff3(t *testing.T) {

	// But this one isn't because the host is missing
	webFingerURLs := ParseAccount("https://@john")
	require.Equal(t, 1, len(webFingerURLs))
	require.Equal(t, "https://@john/.well-known/webfinger?resource=acct:https://@john", webFingerURLs[0])
}

func TestParseURL_WeirdStuff4(t *testing.T) {

	// Test email address with a "+"
	webFingerURLs := ParseAccount("john+connor@connor.com")
	require.Equal(t, 1, len(webFingerURLs))
	require.Equal(t, "https://connor.com/.well-known/webfinger?resource=acct:john+connor@connor.com", webFingerURLs[0])
}

func TestParseURL_WeirdStuff5(t *testing.T) {

	// Test Local Address without a protocol
	webFingerURLs := ParseAccount("localhost/john")
	require.Equal(t, 2, len(webFingerURLs))
	require.Equal(t, "https://localhost/.well-known/webfinger?resource=acct:https://localhost/john", webFingerURLs[0])
	require.Equal(t, "http://localhost/.well-known/webfinger?resource=acct:http://localhost/john", webFingerURLs[1])

	// Test Remote Address without a protocol
	webFingerURLs = ParseAccount("sky.net/sarah")
	require.Equal(t, 1, len(webFingerURLs))
	require.Equal(t, "https://sky.net/.well-known/webfinger?resource=acct:https://sky.net/sarah", webFingerURLs[0])

}

func TestIsValidhostName(t *testing.T) {
	require.True(t, domain.IsValidHostname("localhost"))
	require.True(t, domain.IsValidHostname("127.0.0.1"))
}
