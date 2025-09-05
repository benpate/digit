# digit 👉

[![GoDoc](https://img.shields.io/badge/go-documentation-blue.svg?style=flat-square)](http://pkg.go.dev/github.com/benpate/digit)
[![Version](https://img.shields.io/github/v/release/benpate/digit?include_prereleases&style=flat-square&color=brightgreen)](https://github.com/benpate/digit/releases)
[![Build Status](https://img.shields.io/github/actions/workflow/status/benpate/digit/go.yml?style=flat-square)](https://github.com/benpate/digit/actions/workflows/go.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/benpate/digit?style=flat-square)](https://goreportcard.com/report/github.com/benpate/digit)
[![Codecov](https://img.shields.io/codecov/c/github/benpate/digit.svg?style=flat-square)](https://codecov.io/gh/benpate/digit)

## WebFinger for Go
Digit implements the WebFinger protocol.  It includes type definitions for WebFinger data structures, along with some utilities for sending and receiving WebFinger requests.

## Generating WebFinger Data

Digit provides data types with a simple, chainable API for creating new resources.

``` go
resource := digit.NewResource("acct:sarah@sky.net").
	Alias("http://sky.net/sarah").
	Alias("http://other.website.com/sarah-connor").
	Property("http://sky.net/ns/role", "employee").
	Link(RelationTypeProfile, "text/html", "https://sky.net/sarah")

result, err := json.Marshal(resource)
```

## Retrieving WebFinger Data

Digit can look up WebFinger metadata using a variety of identifiers

``` go
resource, err := digit.Lookup("sarah@sky.net") // Email construction
resource, err := digit.Lookup("sarah@sky.net") // Fediverse "@username" construction
resource, err := digit.Lookup("http://sky.net/sarah") // Canonical URL construction
```

## WebFinger Resources

* [https://webfinger.net](https://wefinger.net) - primary website for WebFinger protocol
* [https://tools.ietf.org/html/rfc7033](https://tools.ietf.org/html/rfc7033) - IETF specification

## Pull Requests Welcome

Digit is relatively stable and is performing well in [Emissary](https://github.com/EmissarySocial/emissary).  However, it is still a work in progress, and will benefit from your experience reports, use cases, and contributions.  If you have an idea for making this library better, send in a pull request.  We're all in this together! 👉
