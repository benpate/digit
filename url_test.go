package digit

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestParseURL(t *testing.T) {

	var webFinger string
	var err error

	// Test URL
	webFinger, err = ParseUsername("https://connor.com/john")
	require.Nil(t, err)
	require.Equal(t, "https://connor.com/.well-known/webfinger?resource=https://connor.com/john", webFinger)

	// Test Fediverse @URL
	webFinger, err = ParseUsername("https://connor.com/@john")
	require.Nil(t, err)
	require.Equal(t, "https://connor.com/.well-known/webfinger?resource=acct:john@connor.com", webFinger)

	// Test simple email address
	webFinger, err = ParseUsername("john@connor.com")
	require.Nil(t, err)
	require.Equal(t, "https://connor.com/.well-known/webfinger?resource=acct:john@connor.com", webFinger)

	// Test Fediverse address
	webFinger, err = ParseUsername("@sarah@sky.net")
	require.Nil(t, err)
	require.Equal(t, "https://sky.net/.well-known/webfinger?resource=acct:sarah@sky.net", webFinger)

}

func TestParseURL_WeirdStuff(t *testing.T) {

	var webFinger string
	var err error

	// Test URL with port
	webFinger, err = ParseUsername("https://connor.com:8080/john")
	require.Nil(t, err)
	require.Equal(t, "https://connor.com:8080/.well-known/webfinger?resource=https://connor.com:8080/john", webFinger)

	// Test Malformed URL
	webFinger, err = ParseUsername("https:///@john")
	require.NotNil(t, err)
	require.Equal(t, "", webFinger)

	// Test email address with a "+"
	webFinger, err = ParseUsername("john+connor@connor.com")
	require.Nil(t, err)
	require.Equal(t, "https://connor.com/.well-known/webfinger?resource=acct:john+connor@connor.com", webFinger)
}
