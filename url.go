package digit

import (
	"net/url"
	"strings"

	"github.com/benpate/rosetta/list"
	"github.com/benpate/uri"
)

// ParseAccount returns a slice of potential WebFinger URLs for a given username.
// If the hostname is a localhost URL, then both HTTP and HTTPS versions are returned.
// Otherwise, only HTTPS endpoints are returned.
func ParseAccount(account string) []string {

	// If the account already has a protocol prefix, then craft a WebFinger URL directly for that.
	if strings.HasPrefix(account, "http://") || strings.HasPrefix(account, "https://") {
		if webFingerURL := parseAccount_ResourceURL(account); webFingerURL != "" {
			return []string{webFingerURL}
		}
		return make([]string, 0)
	}

	// Otherwise, try to parse it like an Fediverse Address / Email Address
	if webFingerURLs := parseAccount_AsHandle(account); len(webFingerURLs) > 0 {
		return webFingerURLs
	}

	// Last Ditch, this might just be a regular URL without a PROTOCOL.  Guess the protocol(s) and continue.
	result := make([]string, 0, 2)

	// Always try https://
	if webFingerURL := parseAccount_ResourceURL("https://" + account); webFingerURL != "" {
		result = append(result, webFingerURL)
	}

	// If this is localhost, then try http://
	if uri.IsLocalHostname(account) {
		if webFingerURL := parseAccount_ResourceURL("http://" + account); webFingerURL != "" {
			result = append(result, webFingerURL)
		}
	}

	// This may be a glorious success, or abject failure. IDK. Deal with it.
	return result
}

// parseAccount_AsHandle identifies a username in the form of a Fediverse Handle or Email address
func parseAccount_AsHandle(account string) []string {

	// Remove the leading "@" from the account name (if it exists)
	account = strings.TrimPrefix(account, "@")

	// If the account doesn't LOOK like an email address, then skip this step
	if !strings.Contains(account, "@") {
		return make([]string, 0)
	}

	// Split into username and hostname
	username, hostname := list.Split(account, '@')

	// RULE: Username must not be empty
	if username == "" {
		return make([]string, 0)
	}

	// RULE: Strip Port number and see if the hostname is valid.
	// If not, we've read the account wrong, and it's not a handle
	if uri.NotValidHostname(uri.NormalizeHost(hostname)) {
		return make([]string, 0)
	}

	// Return the URL *without* the protocol (which will be handled later).
	// The resource value is percent-encoded so that special characters in the
	// account (such as "+" or ":") survive transport to the WebFinger server.
	query := url.Values{"resource": {"acct:" + account}}.Encode()
	urlWithoutProtocol := hostname + "/.well-known/webfinger?" + query

	// Always try the HTTPS version first
	result := []string{"https://" + urlWithoutProtocol}

	// Allow HTTP for localhost (for testing purposes)
	if uri.IsLocalHostname(hostname) {
		result = append(result, "http://"+urlWithoutProtocol)
	}

	return result
}

// parseAccount_ResourceURL translates a standard URL into a WebFinger
// account lookup using the standard "well-known" path and query parameters.
func parseAccount_ResourceURL(resource string) string {

	urlValue, err := url.Parse(resource)

	if err != nil {
		return ""
	}

	if uri.NotValidHostname(uri.NormalizeHost(urlValue.Host)) {
		return ""
	}

	urlValue.Path = ".well-known/webfinger"
	urlValue.RawQuery = url.Values{"resource": {"acct:" + resource}}.Encode()

	return urlValue.String()
}
