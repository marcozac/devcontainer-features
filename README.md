# Dev Container Features

This repository provides the following [Development Container](https://containers.dev/overview) [features](https://containers.dev/implementors/spec/#features).

## Contents

| Feature | Version | Description |
| ------- | ------- | ----------- |
| [atlas](./src/atlas/README.md) | 1.0.3 | [Atlas](https://atlasgo.io) is a language-independent tool for managing and migrating database schemas using modern DevOps principles. |
| [buf](./src/buf/README.md) | 1.0.0 | The Buf CLI enables building and management of Protobuf APIs |
| [shellcheck](./src/shellcheck/README.md) | 1.0.0 | [ShellCheck](https://github.com/koalaman/shellcheck) - A shell script static analysis tool. |

## Usage

To use this features, add them in your `devcontainer.json` as in the example below.

```json
{
    "image": "mcr.microsoft.com/devcontainers/base:debian",
    "features": {
        "ghcr.io/marcozac/devcontainer-features/atlas:1": {},
        "ghcr.io/marcozac/devcontainer-features/buf:1": {},
        "ghcr.io/marcozac/devcontainer-features/shellcheck:1": {}
    }
}
```

## Repo Structure

As in the [`devcontainers/features`](https://github.com/devcontainers/features) repository, each feature is located within a `src` sub-folder.

```
├── src
│   ├── atlas
│   │   ├── devcontainer-feature.json
│   │   └── install.sh
│   ├── buf
│   │   ├── devcontainer-feature.json
│   │   └── install.sh
│   ├── shellcheck
│   │   ├── devcontainer-feature.json
│   │   └── install.sh
```

### Note

:warning: This file is auto-generated. Do not edit it manually.
