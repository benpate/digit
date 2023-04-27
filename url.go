package digit

import (
	"errors"
	"net/mail"
	"net/url"
	"strings"

	"github.com/benpate/domain"
)

// GetWebFingerServer returns the default WebFinger server for a given username.
// It works with email addresses (username@server.com) as well as URLs (https://server.com/username)
func ParseUsername(username string) (string, error) {

	// If the username LOOKS like a URL, then try to parse it like a URL
	if strings.HasPrefix(username, "http://") || strings.HasPrefix(username, "https://") {
		return parseUsernameURL(username)
	}

	// Otherwise, try to parse it like an Email Address

	// In case we've received the Fediverse "@username@server.com" format,
	// remove the leading "@" before parsing like an email
	parsed := strings.TrimPrefix(username, "@")

	// Try to parse the username like an email
	if emailValue, err := mail.ParseAddress(parsed); err == nil {

		// TODO: At some point we may want a more robust way of parsing the email address.
		// Check out http://emailregex.com/
		index := strings.LastIndex(emailValue.Address, `@`)
		hostname := emailValue.Address[index+1:]

		urlValue := url.URL{
			Scheme:   "https",
			Host:     hostname,
			Path:     ".well-known/webfinger",
			RawQuery: "resource=acct:" + emailValue.Address,
		}

		if domain.IsLocalhost(hostname) {
			urlValue.Scheme = "http"
		}

		return urlValue.String(), nil

	}

	// Last Ditch, try to parse it as a URL without a protocol
	if result, err := parseUsernameURL(domain.AddProtocol(username)); err == nil {
		return result, nil
	}

	// Otherwise, we don't recognize this format.

	//lint:ignore ST1005 This is likely a user-facing error message
	return "", errors.New("Username must be a valid URL or Email Address.")
}

func parseUsernameURL(username string) (string, error) {

	urlValue, err := url.Parse(username)

	if err != nil {
		//lint:ignore ST1005 This is likely a user-facing error message
		return "", errors.New("URL must be formatted correctly.")
	}

	if urlValue.Host == "" {
		//lint:ignore ST1005 This is likely a user-facing error message
		return "", errors.New("URL must be formatted correctly.")
	}

	// Build a username that looks like username@domain.com
	parsed := urlValue.Path
	parsed = strings.TrimPrefix(parsed, "/")
	parsed = strings.TrimSuffix(parsed, "/")
	parsed = parsed + "@" + urlValue.Host

	if strings.HasPrefix(parsed, "@") {
		parsed = strings.TrimPrefix(parsed, "@")
		urlValue.Path = ".well-known/webfinger"
		urlValue.RawQuery = "resource=acct:" + parsed
	} else {
		urlValue.Path = ".well-known/webfinger"
		urlValue.RawQuery = "resource=" + username
	}

	return urlValue.String(), nil
}
