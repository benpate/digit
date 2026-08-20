# digit — Notes for AI Agents

A handful of non-obvious behaviors that bite if you don't know them:

- **`Lookup` blocks private/loopback IPs by default.** It uses `benpate/remote`, whose SSRF guard refuses to connect to non-public addresses unless `AllowPrivateIPs(true)` is set. This is correct for production (WebFinger queries arbitrary hosts), but tests against a `127.0.0.1` httptest server must pass a `remote.Option` that flips the guard — otherwise the request fails before it is sent.

- **`NewLink` silently moves `Href` into `Template` for two relation types.** For oStatus subscribe requests (`RelationTypeSubscribeRequest`) and FEP-3b86 intent links (`RelationTypeFEP3b86Prefix`), the constructed Link has an empty `Href` and the URL in `Template` instead. Read `Template`, not `Href`, for those links.

- **`LinkSet` mutators change the slice in place.** `Remove`, `RemoveBy`, `Apply`, and `ApplyBy` rewrite the receiver's backing array, so any other slice header aliasing the same data will be corrupted. Clone first if you need an independent copy. Note also the asymmetry: `Remove` deletes *all* matches; `RemoveBy` deletes only the *first*.

- **The chainable builder methods return copies; the `Set`/`LinkSet` methods mutate.** `Resource.Alias/Property/Link` and `Link.Title/Property` take value receivers and return a modified copy (safe to chain). The `*LinkSet` and `Link.SetString` methods take pointer receivers and mutate.
