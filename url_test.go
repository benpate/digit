package digit

import (
	"testing"

	"github.com/benpate/uri"
	"github.com/stretchr/testify/require"
)

func TestParseURL(t *testing.T) {

	var webFingerURLs []string

	// Test URL
	webFingerURLs = ParseAccount("https://connor.com/john")
	require.Equal(t, 1, len(webFingerURLs))
	require.Equal(t, "https://connor.com/.well-known/webfinger?resource=acct%3Ahttps%3A%2F%2Fconnor.com%2Fjohn", webFingerURLs[0])

	// Test Fediverse @URL
	webFingerURLs = ParseAccount("https://connor.com/@john")
	require.Equal(t, 1, len(webFingerURLs))
	require.Equal(t, "https://connor.com/.well-known/webfinger?resource=acct%3Ahttps%3A%2F%2Fconnor.com%2F%40john", webFingerURLs[0])

	// Test simple email address
	webFingerURLs = ParseAccount("john@connor.com")
	require.Equal(t, 1, len(webFingerURLs))
	require.Equal(t, "https://connor.com/.well-known/webfinger?resource=acct%3Ajohn%40connor.com", webFingerURLs[0])

	// Test Fediverse address
	webFingerURLs = ParseAccount("@sarah@sky.net")
	require.Equal(t, 1, len(webFingerURLs))
	require.Equal(t, "https://sky.net/.well-known/webfinger?resource=acct%3Asarah%40sky.net", webFingerURLs[0])

	// Test Localhost addresses
	webFingerURLs = ParseAccount("http://localhost/john")
	require.Equal(t, 1, len(webFingerURLs))
	require.Equal(t, "http://localhost/.well-known/webfinger?resource=acct%3Ahttp%3A%2F%2Flocalhost%2Fjohn", webFingerURLs[0])

}

func TestParseURL_Error(t *testing.T) {

	var webFingerURLs []string

	// Test previously failed URL
	webFingerURLs = ParseAccount("@first-group@127.0.0.1")
	require.Equal(t, 2, len(webFingerURLs))
	require.Equal(t, "https://127.0.0.1/.well-known/webfinger?resource=acct%3Afirst-group%40127.0.0.1", webFingerURLs[0])
	require.Equal(t, "http://127.0.0.1/.well-known/webfinger?resource=acct%3Afirst-group%40127.0.0.1", webFingerURLs[1])

	// Test previously failed URL
	webFingerURLs = ParseAccount("first-group@127.0.0.1")
	require.Equal(t, 2, len(webFingerURLs))
	require.Equal(t, "https://127.0.0.1/.well-known/webfinger?resource=acct%3Afirst-group%40127.0.0.1", webFingerURLs[0])
	require.Equal(t, "http://127.0.0.1/.well-known/webfinger?resource=acct%3Afirst-group%40127.0.0.1", webFingerURLs[1])
}

func TestParseURL_ActivityPub(t *testing.T) {

	// Test Fediverse localhost address
	webFingerURLs := ParseAccount("@sarah@localhost:3000")
	require.Equal(t, 2, len(webFingerURLs))
	require.Equal(t, "https://localhost:3000/.well-known/webfinger?resource=acct%3Asarah%40localhost%3A3000", webFingerURLs[0])
	require.Equal(t, "http://localhost:3000/.well-known/webfinger?resource=acct%3Asarah%40localhost%3A3000", webFingerURLs[1])
}

func TestParseURL_WeirdStuff(t *testing.T) {

	// Test URL with port
	webFingerURLs := ParseAccount("https://connor.com:8080/john")
	require.Equal(t, 1, len(webFingerURLs))
	require.Equal(t, "https://connor.com:8080/.well-known/webfinger?resource=acct%3Ahttps%3A%2F%2Fconnor.com%3A8080%2Fjohn", webFingerURLs[0])
}

func TestParseURL_WeirdStuff2(t *testing.T) {

	// "@john" resolves to the single-label host "john", which is not a valid
	// hostname (a public domain needs at least two segments), so no WebFinger URL.
	webFingerURLs := ParseAccount("https://@john")
	require.Equal(t, 0, len(webFingerURLs))
}

func TestParseURL_WeirdStuff4(t *testing.T) {

	// Test email address with a "+" (the "+" must be percent-encoded as %2B so
	// the server does not decode it as a space).
	webFingerURLs := ParseAccount("john+connor@connor.com")
	require.Equal(t, 1, len(webFingerURLs))
	require.Equal(t, "https://connor.com/.well-known/webfinger?resource=acct%3Ajohn%2Bconnor%40connor.com", webFingerURLs[0])
}

func TestParseURL_WeirdStuff5(t *testing.T) {

	// Test Local Address without a protocol
	webFingerURLs := ParseAccount("localhost/john")
	require.Equal(t, 2, len(webFingerURLs))
	require.Equal(t, "https://localhost/.well-known/webfinger?resource=acct%3Ahttps%3A%2F%2Flocalhost%2Fjohn", webFingerURLs[0])
	require.Equal(t, "http://localhost/.well-known/webfinger?resource=acct%3Ahttp%3A%2F%2Flocalhost%2Fjohn", webFingerURLs[1])

	// Test Remote Address without a protocol
	webFingerURLs = ParseAccount("sky.net/sarah")
	require.Equal(t, 1, len(webFingerURLs))
	require.Equal(t, "https://sky.net/.well-known/webfinger?resource=acct%3Ahttps%3A%2F%2Fsky.net%2Fsarah", webFingerURLs[0])

}

func TestParseAccount_EmptyHost(t *testing.T) {

	// An "http://" prefix with no hostname cannot become a WebFinger URL.
	require.Equal(t, 0, len(ParseAccount("http://")))

	// Likewise for an "https://" prefix with no hostname.
	require.Equal(t, 0, len(ParseAccount("https://")))
}

func TestParseAccount_AsHandle_EmptyUsername(t *testing.T) {

	// A handle with an empty username (everything before the "@") is not a valid handle.
	// The leading "@" is trimmed, leaving "@connor.com", which splits into an empty username.
	require.Equal(t, 0, len(parseAccount_AsHandle("@@connor.com")))
}

func TestParseAccount_AsHandle_InvalidHostname(t *testing.T) {

	// A handle whose hostname is empty/invalid is not a valid handle.
	require.Equal(t, 0, len(parseAccount_AsHandle("john@")))
}

func TestParseAccount_ResourceURL(t *testing.T) {

	// A valid URL produces a WebFinger lookup URL with a percent-encoded resource.
	require.Equal(t,
		"https://connor.com/.well-known/webfinger?resource=acct%3Ahttps%3A%2F%2Fconnor.com%2Fjohn",
		parseAccount_ResourceURL("https://connor.com/john"),
	)

	// Characters that are significant in a query string ("&", "?", "=") must be
	// escaped so they cannot corrupt the "resource" parameter.
	require.Equal(t,
		"https://host.com/.well-known/webfinger?resource=acct%3Ahttps%3A%2F%2Fhost.com%2Fa%3Fb%3Dc%26d%3De",
		parseAccount_ResourceURL("https://host.com/a?b=c&d=e"),
	)

	// A URL with no hostname returns an empty string.
	require.Equal(t, "", parseAccount_ResourceURL("https://"))

	// A URL that cannot be parsed (control character) returns an empty string.
	require.Equal(t, "", parseAccount_ResourceURL("https://example.com/\x7f"))
}

func TestIsValidhostName(t *testing.T) {
	require.True(t, uri.IsValidHostname("localhost"))
	require.True(t, uri.IsValidHostname("127.0.0.1"))
}

// FuzzParseAccount throws arbitrary account strings at the parser and verifies
// that the invariants of its output always hold (and that it never panics).
func FuzzParseAccount(f *testing.F) {

	// Seed the corpus with the shapes exercised by the unit tests above.
	seeds := []string{
		"",
		"@",
		"@@",
		"http://",
		"https://",
		"https://connor.com/john",
		"https://connor.com/@john",
		"john@connor.com",
		"@sarah@sky.net",
		"http://localhost/john",
		"@sarah@localhost:3000",
		"localhost/john",
		"sky.net/sarah",
		"https://@john",
		"john+connor@connor.com",
		"@first-group@127.0.0.1",
	}

	for _, seed := range seeds {
		f.Add(seed)
	}

	f.Fuzz(func(t *testing.T, account string) {

		results := ParseAccount(account)

		// ParseAccount returns at most two candidate URLs (https and http).
		require.LessOrEqual(t, len(results), 2)

		for _, result := range results {
			// Every candidate must be a non-empty WebFinger endpoint.
			require.NotEmpty(t, result)
			require.Contains(t, result, "/.well-known/webfinger")
		}

		// The parser must be deterministic.
		require.Equal(t, results, ParseAccount(account))
	})
}
