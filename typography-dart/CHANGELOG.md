# Changelog

All notable changes to this monorepo are documented here. Packages are
versioned together; tags use the `vMAJOR.MINOR.PATCH` form.

## [0.1.0] - 2026-06-26

Initial release.

### typography_core
- `nbsp` (U+00A0) and `narrowNbsp` (U+202F) constants.
- `Rule` type (`RegExp` + `Replacer`) with a `Rule.constant` convenience.
- `applyRules(input, rules)` engine.

### typography_ru
- Full Russian UI micro-typography per the canonical spec, public `ru(String)`:
  - NBSP before em dash «—»;
  - NBSP after 1–2 letter Cyrillic words (hyphenated fragments excluded);
  - narrow NBSP between a number and `%`;
  - narrow NBSP between a number and a currency sign (order preserved);
  - NBSP between a numeral and its dependent noun (Cyrillic or Latin);
  - quotes «» (first level) / „“ (nested), straight `"` converted.
- Backwards compatible with the reference implementation for rules 1–2.

### typography_en
- Minimal scaffold (`en(String)` identity passthrough) as a new-language template.
