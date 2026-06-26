# typography-go

Pure-Go micro-typography for UI strings, organized by language. A small,
dependency-free engine plus per-language rule sets. The Russian package (`ru`) is
a behavioral port of the [Dart implementation](../typography-dart/) and matches
it byte-for-byte.

Part of the [`typography-packages`](https://github.com/vziger/typography-packages)
umbrella repository. The Go module lives in the `typography-go/` subdirectory, so
its module path and release tags carry that prefix.

## Install

```sh
go get github.com/vziger/typography-packages/typography-go/ru@v0.1.0
```

The version tag for this subdirectory module is `typography-go/v0.1.0` (a plain
`v0.1.0` belongs to the Dart part). `go get` with `@v0.1.0` resolves to it.

## Usage

```go
package main

import (
	"fmt"

	"github.com/vziger/typography-packages/typography-go/ru"
)

func main() {
	fmt.Println(ru.Ru("Соколы — это 10 дней и 50% скидки"))
	// Соколы — это 10 дней и 50 % скидки
	// (with U+00A0 before the dash and the noun, U+202F before %)
}
```

`ru.Ru` applies the full Russian rule set:

1. NBSP (U+00A0) before an em dash «—».
2. NBSP after every 1–2 letter Cyrillic word (hyphenated fragments like
   *какой-то* are excluded).
3. Narrow NBSP (U+202F) between a number and `%`.
4. Narrow NBSP (U+202F) between a number and a currency sign (`€ ₽ $ £ ¥`),
   preserving order (`€ 10` / `10 ₽`).
5. NBSP (U+00A0) between a numeral and its dependent noun, Cyrillic or Latin
   (`10 дней`, `3 items`).
6. Quotes: first level «», nested „“; straight `"…"` are converted.

The rules are idempotent: `ru.Ru(ru.Ru(x)) == ru.Ru(x)`.

## Specification

The canonical, authoritative rule specification (with ❌/✅ examples) is a single
local file and is **not** vendored into this repository:

```
/Users/muntu/.cursor/rules/ui-typography-ru.mdc
```

That file is the single source of truth for both the Go and Dart
implementations; this README only summarizes it.

## Packages

| Package | Path                                                            | Purpose                                          |
| ------- | --------------------------------------------------------------- | ------------------------------------------------ |
| `core`  | `github.com/vziger/typography-packages/typography-go/core`      | NBSP constants, the `Rule` type, `Apply` engine. |
| `ru`    | `github.com/vziger/typography-packages/typography-go/ru`        | Russian rules + quote scanner. Entry: `Ru`.      |
| `en`    | `github.com/vziger/typography-packages/typography-go/en`        | English placeholder (identity). Template.        |

## How to add a language

A language package is just an ordered set of `core.Rule` values plus, optionally,
any stateful pre/post passes (like the Russian quote scanner). To add one:

1. Create a new package directory (e.g. `de/`).
2. Build a `[]core.Rule` of regex rules. Because Go's `regexp` is RE2, it has
   **no look-ahead/look-behind** — express context with capture groups and a
   `core.Replacer` (`func(groups []string) string`), as `ru` does. Spell
   whitespace classes out as `[\s\x{00A0}\x{202F}]` (Go's `\s` is ASCII-only and
   excludes NBSP); this also keeps rules idempotent.
3. Expose a single entry point, e.g. `func De(s string) string`, that runs any
   stateful passes and then `core.Apply(s, rules)`. Use the NBSP constants from
   `core` — never paste a raw U+00A0/U+202F into source, as it is silently
   normalized to a plain space by some tooling.
4. Add table-driven tests (port reference cases 1:1) and an idempotency check.

See `ru/ru.go` for a complete worked example.

## Development

```sh
go vet ./...
go test ./...
```

## License

[MIT](LICENSE).
