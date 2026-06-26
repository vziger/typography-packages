// Package ru applies Russian UI micro-typography.
//
// It implements the full set of rules from the canonical spec
// /Users/muntu/.cursor/rules/ui-typography-ru.mdc:
//
//  1. NBSP (U+00A0) before an em dash «—».
//  2. NBSP after every 1–2 letter Cyrillic word (hyphenated fragments excluded).
//  3. Narrow NBSP (U+202F) between a number and `%`.
//  4. Narrow NBSP (U+202F) between a number and a currency sign (order kept).
//  5. NBSP (U+00A0) between a numeral and its dependent noun (Cyrillic or Latin).
//  6. Quotes: first level «», nested „“ (straight `"` are converted).
//
// The public entry point is Ru. The rules are idempotent: Ru(Ru(x)) == Ru(x).
//
// Notes on the Go port (RE2 has no look-around, and Go's \s is ASCII-only):
//   - Rules 3–5 are rewritten with capture groups instead of look-behind/ahead.
//   - Whitespace classes are spelled out as [\s\x{00A0}\x{202F}] so the rules
//     stay idempotent (a previously inserted NBSP is recognized as whitespace).
//   - Rule 2 (short words) is a manual rune scanner, not a regex: a capturing
//     regex would consume the boundary character and break chains like «я и ты».
package ru

import (
	"regexp"
	"unicode"

	"github.com/vziger/typography-packages/typography-go/core"
)

// currency lists the currency symbols handled by rule 4 (matches the Dart set).
const currency = `€₽$£¥`

// ws is a whitespace character class that, unlike Go's ASCII-only \s, also
// covers the non-breaking spaces this package inserts. Including them keeps the
// regex rules idempotent.
const ws = `[\s\x{00A0}\x{202F}]`

// RussianRules is the ordered list of regex rules applied by Ru after quote
// normalization and before the short-word scanner. Exported for reuse/testing.
//
// Order matters: percent and currency run before the numeral+noun rule so that
// e.g. `10 €` gets a narrow NBSP rather than a regular one.
var RussianRules = []core.Rule{
	// Rule 1: NBSP before an em dash «—». A whitespace run before the dash
	// collapses to a single NBSP.
	core.Const(regexp.MustCompile(ws+`+—`), core.Nbsp+"—"),

	// Rule 3: narrow NBSP between a number and the percent sign.
	{
		Pattern: regexp.MustCompile(`(\d)` + ws + `*%`),
		Replace: func(g []string) string { return g[1] + core.NarrowNbsp + "%" },
	},

	// Rule 4a: narrow NBSP between a number and a trailing currency sign
	// (`10₽` / `10 €` → `10 ₽`). Order is preserved.
	{
		Pattern: regexp.MustCompile(`(\d)` + ws + `*([` + currency + `])`),
		Replace: func(g []string) string { return g[1] + core.NarrowNbsp + g[2] },
	},
	// Rule 4b: narrow NBSP between a leading currency sign and a number
	// (`€10` / `€ 10` → `€ 10`).
	{
		Pattern: regexp.MustCompile(`([` + currency + `])` + ws + `*(\d)`),
		Replace: func(g []string) string { return g[1] + core.NarrowNbsp + g[2] },
	},

	// Rule 5: NBSP between a numeral and a dependent noun, Cyrillic or Latin
	// (`10 дней`, `3 items`). A whitespace run collapses to one NBSP.
	{
		Pattern: regexp.MustCompile(`(\d)` + ws + `+([A-Za-zа-яёА-ЯЁ])`),
		Replace: func(g []string) string { return g[1] + core.Nbsp + g[2] },
	},
}

// Ru returns s with Russian UI micro-typography applied. An empty string is
// returned unchanged. Latin text and code are left untouched.
//
// Pipeline order (quotes first so «/„ are present as opening context for the
// short-word rule): quotes → em dash → % → currency → numeral+noun → short words.
func Ru(s string) string {
	if s == "" {
		return s
	}
	out := applyRussianQuotes(s)
	out = core.Apply(out, RussianRules)
	out = applyShortWords(out)
	return out
}

// applyShortWords implements rule 2: a regular space (0x20) after a 1–2 letter
// Cyrillic word becomes an NBSP, provided that word is at the start of the
// string or preceded by a boundary (whitespace / `(` / `«` / `„` / `"`). A
// fragment after a hyphen (какой-то) is excluded because `-` is not a boundary.
//
// Implemented as a single left-to-right rune scan over the original string so
// that chains like «я и ты» all get NBSPs (a capturing regex would eat the
// boundary char and break the next match). Only 0x20 is replaced, which keeps
// the rule idempotent.
func applyShortWords(s string) string {
	runes := []rune(s)
	out := make([]rune, len(runes))
	copy(out, runes)

	for i, ch := range runes {
		if ch != ' ' {
			continue
		}
		// Count Cyrillic letters immediately before the space, capped at 3:
		// only a run of exactly 1 or 2 qualifies as a short word.
		j := i - 1
		cnt := 0
		for j >= 0 && cnt < 3 && isCyrillic(runes[j]) {
			cnt++
			j--
		}
		if cnt != 1 && cnt != 2 {
			continue
		}
		// j now points just before the word; it must be the string start or a
		// boundary character (so a hyphen suffix like `-то` does not qualify).
		if j >= 0 && !isBoundary(runes[j]) {
			continue
		}
		out[i] = '\u00A0'
	}
	return string(out)
}

// isCyrillic reports whether r is a Cyrillic letter handled by rule 2 (matches
// the Dart class [а-яёА-ЯЁ]).
func isCyrillic(r rune) bool {
	return (r >= 'а' && r <= 'я') || (r >= 'А' && r <= 'Я') || r == 'ё' || r == 'Ё'
}

// isBoundary reports whether r counts as a left boundary for a short word:
// whitespace (including non-breaking spaces) or an opening bracket / quote.
func isBoundary(r rune) bool {
	if unicode.IsSpace(r) {
		return true
	}
	switch r {
	case '(', '«', '„', '"':
		return true
	}
	return false
}
