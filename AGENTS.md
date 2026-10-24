# digit — Notes for AI Agents

A handful of non-obvious behaviors that bite if you don't know them:

- **`Lookup` blocks private/loopback IPs by default.** It uses `benpate/remote`, whose SSRF guard refuses to connect to non-public addresses unless `AllowPrivateIPs(true)` is set. This is correct for production (WebFinger queries arbitrary hosts), but tests against a `127.0.0.1` httptest server must pass a `remote.Option` that flips the guard — otherwise the request fails before it is sent.

- **`NewLink` silently moves `Href` into `Template` for two relation types.** For oStatus subscribe requests (`RelationTypeSubscribeRequest`) and FEP-3b86 intent links (`RelationTypeFEP3b86Prefix`), the constructed Link has an empty `Href` and the URL in `Template` instead. Read `Template`, not `Href`, for those links.

- **`LinkSet` mutators change the slice in place.** `Remove`, `RemoveBy`, `Apply`, and `ApplyBy` rewrite the receiver's backing array, so any other slice header aliasing the same data will be corrupted. Clone first if you need an independent copy. Note also the asymmetry: `Remove` deletes *all* matches; `RemoveBy` deletes only the *first*.

- **The chainable builder methods return copies; the `Set`/`LinkSet` methods mutate.** `Resource.Alias/Property/Link` and `Link.Title/Property` take value receivers and return a modified copy (safe to chain). The `*LinkSet` and `Link.SetString` methods take pointer receivers and mutate.

- **A chainable builder must clone a map and clip a slice before writing — a value receiver copies the header, not the data.** The map-valued builders call `maps.Clone`; the slice-valued ones use `append(slices.Clip(s), v)`. Without the clip, two chains branching from the same `Resource` wrote into one backing array as soon as it had spare capacity (len 3, cap 4) and clobbered each other. 100% statement coverage never caught it, because a single linear chain always passes — the probe has to branch twice from a base with `cap > len`. `Resource.Link` delegates to `AddLink` so there is exactly one guarded append; keep it that way, and apply the same rule to any new builder here.

- **`parseAccount_ResourceURL` prefixes `acct:` even onto a URL, and that is deliberate.** A URL-form lookup goes out as `resource=acct:https://connor.com/john`, which is not a valid `acct:` URI — RFC 7565 wants `acct:user@host`, and Mastodon expects the bare URL. Six `url_test.go` cases assert it verbatim, so it is locked in on purpose: changing it is a federation-behavior change across every `digit.Lookup` caller. Do not "fix" it in a quality sweep and do not re-flag it as new. Emissary hides the breakage in-house because `service/locator.go` trims the `acct:` prefix before deciding URL-vs-handle. The handle path in `parseAccount_AsHandle` is correct as-is.
