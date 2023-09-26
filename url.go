package digit

import (
	"net/url"
	"strings"

	"github.com/benpate/domain"
	"github.com/benpate/rosetta/list"
)

// ParseAccount returns a slice of potential WebFinger URLs for a given username.
// If the hostname is a localhost URL, then both HTTP and HTTPS versions are returned.
// Otherwise, only HTTPS endpoints are returned.
func ParseAccount(account string) []string {

	// If the account already has an HTTP prefix, then craft a WebFinger URL directly for that.
	if strings.HasPrefix(account, "http://") {
		if webFingerURL := parseAccount_ResourceURL(account); webFingerURL != "" {
			return []string{webFingerURL}
		} else {
			return make([]string, 0)
		}
	}

	// If the account already has an HTTPS prefix, then craft a WebFinger URL directly for that.
	if strings.HasPrefix(account, "https://") {
		if webFingerURL := parseAccount_ResourceURL(account); webFingerURL != "" {
			return []string{webFingerURL}
		} else {
			return make([]string, 0)
		}
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
	if domain.IsLocalhost(account) {
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
	if domain.NotValidHostname(hostname) {
		return make([]string, 0)
	}

	// Return the URL *without* the protocol (which will be handled later)
	urlWithoutProtocol := hostname + "/.well-known/webfinger?resource=acct:" + account

	// Always try the HTTPS version first
	result := []string{"https://" + urlWithoutProtocol}

	// Allow HTTP for localhost (for testing purposes)
	if domain.IsLocalhost(hostname) {
		result = append(result, "http://"+urlWithoutProtocol)
	}

	return result
}

// parseAccount_AsResourceURL translates a standard URL into a WebFinger
// account lookup using the standard "well-known" path and query parameters.
func parseAccount_ResourceURL(resource string) string {

	urlValue, err := url.Parse(resource)

	if err != nil {
		return ""
	}

	if domain.NotValidHostname(urlValue.Host) {
		return ""
	}

	urlValue.Path = ".well-known/webfinger"
	urlValue.RawQuery = "resource=acct:" + resource

	return urlValue.String()
}
