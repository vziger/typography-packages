# Changelog

All notable changes to the Go implementation are documented here. This project
adheres to [Semantic Versioning](https://semver.org/). Because the Go module
lives in a subdirectory of the umbrella repository, release tags are prefixed
with the module path: `typography-go/vX.Y.Z`.

## [0.1.0] - 2026-06-26

Initial release. Go port of the Russian micro-typography rules, behaviorally
matching the Dart reference byte-for-byte.

### Added

- `core` package: NBSP constants (`Nbsp`, `NarrowNbsp`), the `Rule` type and the
  `Apply` rule-application engine.
- `ru` package: `Ru(string) string` implementing the full rule set — NBSP before
  an em dash, NBSP after 1–2 letter Cyrillic words, narrow NBSP before `%` and
  currency signs (order preserved), NBSP between a numeral and its dependent noun
  (Cyrillic or Latin), and stateful quote normalization («», nested „“). Rules
  are idempotent.
- `en` package: `En(string) string` identity passthrough as a language template.
- Test suite porting the 23 reference cases from the Dart package plus an
  idempotency check.
