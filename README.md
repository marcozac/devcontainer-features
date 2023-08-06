# Dev Container Features

This repository provides the following [Development Container](https://containers.dev/overview) [features](https://containers.dev/implementors/spec/#features).

## Contents

| Feature | Version | Description |
| ------- | ------- | ----------- |
| [buf](./src/buf/README.md) | 1.0.0 | The Buf CLI enables building and management of Protobuf APIs |

## Usage

To use this features, add them in your `devcontainer.json` as in the example below.

```json
{
    "image": "mcr.microsoft.com/devcontainers/base:debian",
    "features": {
        "ghcr.io/marcozac/devcontainer-features/buf:1": {}
    }
}
```

## Repo Structure

As in the [`devcontainers/features`](https://github.com/devcontainers/features) repository, each feature is located within a `src` sub-folder.

```
├── src
│   ├── buf
│   │   ├── devcontainer-feature.json
│   │   └── install.sh
```

### Note

:warning: This file is auto-generated. Do not edit it manually.
