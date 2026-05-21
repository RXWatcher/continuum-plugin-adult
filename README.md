# Continuum Adult Metadata Plugin

Multi-source metadata aggregator for adult content. Registers a single
`metadata_provider.v1` capability and fans out lookups across the upstream
sources the operator has configured.

The plugin's content mapping follows TPDB's data model:

| TPDB entity | Host content type |
|-------------|-------------------|
| Movie       | Movie             |
| Site        | Series            |
| Scene       | Movie or Episode  |
| Performer   | Person            |

Sites map to TV series with a synthetic single season; scenes belonging to
the site become episodes ordered by release date. Standalone scenes (no
parent site context) are exposed as movies.

## Sources

Day-one source: [ThePornDB (TPDB)](https://theporndb.net). Configure via
the plugin admin UI by enabling the TPDB card and pasting an API key (TPDB
account → API).

Additional sources slot in as new `provider/sources/<name>/` packages
implementing the `provider.Source` interface and a corresponding entry in
`manifest.json`'s `global_config_schema`. Stash is the next planned source.

## Dependency Model

This repository consumes
`github.com/ContinuumApp/continuum-plugin-sdk` as a normal Go module
dependency. CI and release builds run with `GOWORK=off` and expect the SDK
version in `go.mod` to resolve from a published semver tag.

For local multi-repo development, use a temporary `replace` or a local
`go.work` that points at `dev/github/continuum-plugin-sdk`. Do not commit
machine-local filesystem replaces as the supported release path.

## Development

```sh
go test ./...
go build .
```

## License

`continuum-plugin-adult` is licensed under `AGPL-3.0-or-later`. See
[LICENSE](LICENSE).
