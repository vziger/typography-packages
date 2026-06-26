# typography-packages

Umbrella repository for micro-typography libraries, organized by language /
ecosystem. Each top-level directory is an independent project with its own
build, tests and versioning.

| Directory                          | Status      | Description                                  |
| ---------------------------------- | ----------- | -------------------------------------------- |
| [`typography-dart/`](typography-dart/) | ✅ available | Pure-Dart micro-typography (core + `ru`, `en`). |
| `typography-go/`                   | planned     | Go implementation.                           |

See each project's own `README.md` for usage, installation and contribution
notes. CI for the Dart project lives in [`.github/workflows/ci.yml`](.github/workflows/ci.yml)
and runs only when `typography-dart/**` changes.
