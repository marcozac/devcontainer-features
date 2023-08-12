
# Atlas CLI (atlas)

[Atlas](https://atlasgo.io) is a language-independent tool for managing and migrating database schemas using modern DevOps principles.

## Example Usage

```json
"features": {
    "ghcr.io/marcozac/devcontainer-features/atlas:1": {}
}
```

## Options

| Options Id | Description | Type | Default Value |
|-----|-----|-----|-----|
| version | Select or enter an Atlas version. | string | latest |
| communityEdition | Install the Community Edition of Atlas building from source. | boolean | true |
| goVersion | The Go version to use to build a specific (community edition) version of the Atlas CLI | string | 1.20.7 |

## Community Edition

Only the latest release of the [Community Edition of Atlas](https://atlasgo.io/community-edition) is available as pre-built binary.

Choosing a specific version will cause it to be built from source.

### Building from source

[Go](https://go.dev/) is required to build the Atlas CLI.

If `go` is neither in the `PATH` nor `/usr/local/go/bin/go`, it is installed from the [Go Feature](https://github.com/devcontainers/features/tree/main/src/go) (skipping tools installation) and is currently available for Debian/Ubuntu-based distributions with the apt package manager installed.

## Third Party

-   The [Atlas installation script](./atlas.sh) is a copy of the original one from https://atlasgo.sh.
-   The [library_scripts.sh](./library_scripts.sh) file is a copy of one of the many present in the [devcontainers-contrib repository](https://github.com/devcontainers-contrib/features/).


---

_Note: This file was auto-generated from the [devcontainer-feature.json](https://github.com/marcozac/devcontainer-features/blob/main/src/atlas/devcontainer-feature.json).  Add additional notes to a `NOTES.md`._
