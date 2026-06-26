# typography-packages

Umbrella repository for micro-typography libraries, organized by language /
ecosystem. Each top-level directory is an independent project with its own
build, tests and versioning.

| Directory                          | Status      | Description                                  |
| ---------------------------------- | ----------- | -------------------------------------------- |
| [`typography-dart/`](typography-dart/) | ✅ available | Pure-Dart micro-typography (core + `ru`, `en`). |
| [`typography-go/`](typography-go/) | ✅ available | Pure-Go micro-typography (core + `ru`, `en`). |

See each project's own `README.md` for usage, installation and contribution
notes. CI for the Dart project lives in [`.github/workflows/ci.yml`](.github/workflows/ci.yml)
(runs on `typography-dart/**`); the Go project's CI is in
[`.github/workflows/ci-go.yml`](.github/workflows/ci-go.yml) (runs on
`typography-go/**`).
