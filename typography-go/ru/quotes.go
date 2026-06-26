package ru

import (
	"strings"
	"unicode"
)

// applyRussianQuotes replaces straight double quotes `"..."` with Russian
// typographic quotes: first-level pairs become «…», nested pairs become „…“.
//
// This is a small stateful scanner rather than a single regex, because the same
// character `"` serves as both opener and closer and the correct glyph depends
// on nesting depth and the preceding character. Text without a `"` is returned
// unchanged. Unbalanced quotes are handled defensively: a `"` while nothing is
// open always opens. Ported from the Dart reference quotes.dart.
func applyRussianQuotes(input string) string {
	if !strings.ContainsRune(input, '"') {
		return input
	}

	var b strings.Builder
	runes := []rune(input)
	depth := 0
	for i, ch := range runes {
		if ch != '"' {
			b.WriteRune(ch)
			continue
		}

		hasPrev := i > 0
		var prev rune
		if hasPrev {
			prev = runes[i-1]
		}
		isOpener := depth == 0 || opensAfter(hasPrev, prev)
		switch {
		case isOpener:
			if depth == 0 {
				b.WriteRune('«')
			} else {
				b.WriteRune('„')
			}
			depth++
		case depth >= 2:
			b.WriteRune('“')
			depth--
		default:
			b.WriteRune('»')
			depth = 0
		}
	}
	return b.String()
}

// opensAfter reports whether a `"` preceded by prev should open a quote.
//
// A quote opens at the start of the string or after whitespace / an opening
// bracket / an already-open quote; otherwise it closes.
func opensAfter(hasPrev bool, prev rune) bool {
	if !hasPrev {
		return true
	}
	if unicode.IsSpace(prev) {
		return true
	}
	switch prev {
	case '(', '[', '«', '„':
		return true
	}
	return false
}
