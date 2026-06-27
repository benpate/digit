# digit 👉

[![GoDoc](https://img.shields.io/badge/go-documentation-blue.svg?style=flat-square)](http://pkg.go.dev/github.com/benpate/digit)
[![Version](https://img.shields.io/github/v/release/benpate/digit?include_prereleases&style=flat-square&color=brightgreen)](https://github.com/benpate/digit/releases)
[![Build Status](https://img.shields.io/github/actions/workflow/status/benpate/digit/go.yml?style=flat-square)](https://github.com/benpate/digit/actions/workflows/go.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/benpate/digit?style=flat-square)](https://goreportcard.com/report/github.com/benpate/digit)
[![Codecov](https://img.shields.io/codecov/c/github/benpate/digit.svg?style=flat-square)](https://codecov.io/gh/benpate/digit)

## WebFinger for Go
Digit implements the WebFinger protocol.  It includes type definitions for WebFinger data structures, along with some utilities for sending and receiving WebFinger requests

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

## What matters here

A handful of non-obvious behaviors that bite if you don't know them:

- **`Lookup` blocks private/loopback IPs by default.** It uses `benpate/remote`, whose SSRF guard refuses to connect to non-public addresses unless `AllowPrivateIPs(true)` is set. This is correct for production (WebFinger queries arbitrary hosts), but tests against a `127.0.0.1` httptest server must pass a `remote.Option` that flips the guard — otherwise the request fails before it is sent.

- **`NewLink` silently moves `Href` into `Template` for two relation types.** For oStatus subscribe requests (`RelationTypeSubscribeRequest`) and FEP-3b86 intent links (`RelationTypeFEP3b86Prefix`), the constructed Link has an empty `Href` and the URL in `Template` instead. Read `Template`, not `Href`, for those links.

- **`LinkSet` mutators change the slice in place.** `Remove`, `RemoveBy`, `Apply`, and `ApplyBy` rewrite the receiver's backing array, so any other slice header aliasing the same data will be corrupted. Clone first if you need an independent copy. Note also the asymmetry: `Remove` deletes *all* matches; `RemoveBy` deletes only the *first*.

- **The chainable builder methods return copies; the `Set`/`LinkSet` methods mutate.** `Resource.Alias/Property/Link` and `Link.Title/Property` take value receivers and return a modified copy (safe to chain). The `*LinkSet` and `Link.SetString` methods take pointer receivers and mutate.

## WebFinger Resources

* [https://webfinger.net](https://webfinger.net) - primary website for WebFinger protocol
* [https://tools.ietf.org/html/rfc7033](https://tools.ietf.org/html/rfc7033) - IETF specification

## Pull Requests Welcome

Digit is relatively stable and is performing well in [Emissary](https://github.com/EmissarySocial/emissary).  However, it is still a work in progress, and will benefit from your experience reports, use cases, and contributions.  If you have an idea for making this library better, send in a pull request.  We're all in this together! 👉
