# typography-packages

Pure-Dart (no Flutter) micro-typography libraries, structured as a monorepo with
one package per language on top of a shared engine.

| Package            | What it does                                                        |
| ------------------ | ------------------------------------------------------------------- |
| `typography_core`  | Engine: NBSP constants, `Rule` type (regex + replacement), `applyRules`. |
| `typography_ru`    | Full Russian UI micro-typography. Public entry point: `ru(String)`. |
| `typography_en`    | Minimal scaffold (`en(String)` passthrough) — template for new languages. |

No package depends on Flutter; everything is plain Dart and tested with
`package:test`.

## Rules (source of truth)

The Russian rules are implemented from a single canonical specification — **do
not duplicate it**. The only source of truth is:

```
/Users/muntu/.cursor/rules/ui-typography-ru.mdc
```

`typography_ru` implements the full set described there:

1. NBSP (U+00A0) before an em dash «—».
2. NBSP after every 1–2 letter Cyrillic word (abbreviations, code, and
   hyphenated fragments like *какой-то* excluded).
3. Narrow NBSP (U+202F) between a number and `%`.
4. Narrow NBSP (U+202F) between a number and a currency sign, order preserved
   (`€ 10` or `10 ₽`).
5. NBSP (U+00A0) between a numeral and its dependent noun, Cyrillic or
   Latin (`10 дней`, `3 items`).
6. Quotes: first level «», nested „“; straight `"…"` are converted.

## Usage

```dart
import 'package:typography_ru/typography_ru.dart';

void main() {
  print(ru('в команде'));   // → в команде        (NBSP after «в»)
  print(ru('скидка 10%'));  // → скидка 10 %       (narrow NBSP)
  print(ru('"текст"'));     // → «текст»
}
```

## Adding it as a git dependency

This lives inside the umbrella repo `typography-packages`, so point pub at the
package subdirectory with `path:` (note the `typography-dart/` prefix) and pin to
a tag with `ref:`. Replace `<OWNER>/<REPO>` with the actual repository.

```yaml
dependencies:
  typography_ru:
    git:
      url: https://github.com/<OWNER>/<REPO>.git
      ref: v0.1.0
      path: typography-dart/packages/typography_ru
```

`typography_core` is pulled in automatically: `typography_ru` depends on it via a
relative `path: ../typography_core`, which resolves inside the same git checkout.

## How to add a new language

Each language is a new `typography_xx` package built on `typography_core`. Use
`typography_en` as a template:

1. `mkdir -p packages/typography_xx/lib/src packages/typography_xx/test`.
2. Add `packages/typography_xx/pubspec.yaml` depending on
   `typography_core: { path: ../typography_core }` plus dev deps `lints` and
   `test`, and `analysis_options.yaml` with `include: ../../analysis_options.yaml`.
3. In `lib/src/rules.dart`, build an ordered `final List<Rule> xxRules = [...]`.
   Each rule is a `RegExp` plus a `Replacer` (or `Rule.constant(pattern, text)`).
   For stateful transforms that a single regex cannot express (like Russian
   quote nesting), add a dedicated function and call it inside the facade.
4. In `lib/typography_xx.dart`, expose a facade:
   ```dart
   String xx(String s) => s.isEmpty ? s : applyRules(s, xxRules);
   ```
5. Add tests under `test/` using `package:test`. Define NBSP literals as
   `'\u00A0'` / `'\u202F'` escapes (raw NBSP characters are fragile in source).
6. Add the package directory to the CI matrix in `.github/workflows/ci.yml`.

## Development

```bash
# per package
cd packages/typography_ru
dart pub get
dart analyze
dart test
```

## License

[MIT](LICENSE).
