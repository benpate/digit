package digit

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/benpate/remote"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// allowPrivateIPs lets a test reach a loopback httptest server. remote blocks
// non-public IPs by default (SSRF protection), so the guard must be disabled
// before the request is sent.
var allowPrivateIPs = remote.Option{
	BeforeRequest: func(txn *remote.Transaction) error {
		txn.AllowPrivateIPs(true)
		return nil
	},
}

func TestLookup(t *testing.T) {

	// Spin up a WebFinger server that returns a fixed resource.
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Use assert (not require) here: this handler runs in the server's
		// goroutine, where require's FailNow would call runtime.Goexit.
		assert.Equal(t, "/.well-known/webfinger", r.URL.Path)
		w.Header().Set("Content-Type", "application/json")
		_, err := w.Write([]byte(`{"subject":"acct:sarah@sky.net","links":[{"rel":"self","type":"application/activity+json","href":"https://sky.net/users/sarah"}]}`))
		assert.NoError(t, err)
	}))
	t.Cleanup(server.Close)

	resource, err := Lookup(server.URL+"/@sarah", allowPrivateIPs)

	require.Nil(t, err)
	require.Equal(t, "acct:sarah@sky.net", resource.Subject)
	require.Equal(t, 1, len(resource.Links))

	self := resource.FindLink("self")
	require.Equal(t, "https://sky.net/users/sarah", self.Href)
}

func TestLookup_NoWebFingerURLs(t *testing.T) {

	// An account that produces no candidate URLs cannot be looked up.
	_, err := Lookup("http://")
	require.NotNil(t, err)
}

func TestLookup_ServerError(t *testing.T) {

	// A server that returns an error status causes the lookup to fail.
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	}))
	t.Cleanup(server.Close)

	_, err := Lookup(server.URL+"/@sarah", allowPrivateIPs)
	require.NotNil(t, err)
}

/*
// These tests require live external servers and are kept for manual use.

func TestClient(t *testing.T) {
	resource, err := Lookup("http://localhost/@benpate")
	require.Nil(t, err)
	t.Log(resource)
}

func TestMitra(t *testing.T) {
	resource, err := Lookup("@benpate@wizard.casa")
	require.Nil(t, err)
	require.Equal(t, "acct:benpate@wizard.casa", resource.Subject)

	self := resource.FindLink("self")
	require.Equal(t, "https://wizard.casa/users/benpate", self.Href)
}
*/
