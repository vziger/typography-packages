// Package core holds the language-agnostic building blocks for the
// micro-typography rules: the non-breaking space constants, the Rule type and
// the rule-application engine. Language packages (such as ru) are just an
// ordered slice of Rule values fed to Apply.
//
// Source of truth for the rules:
// /Users/muntu/.cursor/rules/ui-typography-ru.mdc
package core

import "regexp"

// Non-breaking space characters used by the typography rules.
//
// They are written as escape sequences on purpose: a raw U+00A0 / U+202F byte
// in source is easily and silently normalized to a plain space (0x20) by
// editors and tooling, which would break the rules. Always build expected
// values from these constants, never from a pasted literal.
const (
	// Nbsp is the regular non-breaking space, U+00A0. Used between short words,
	// before an em dash, and between a numeral and its dependent noun.
	Nbsp = "\u00A0"
	// NarrowNbsp is the narrow (half) non-breaking space, U+202F. Used between a
	// number and a % sign or a currency symbol.
	NarrowNbsp = "\u202F"
)

// Replacer builds the replacement string for one match.
//
// It receives the result of regexp.FindStringSubmatch for the match: groups[0]
// is the full match and groups[i] is the i-th capture group. This mirrors the
// Dart reference's per-match Replacer and lets rules be group-aware, which is
// required because Go's RE2 regexp has no look-around.
type Replacer func(groups []string) string

// Rule is a single micro-typography rule: a Pattern plus how to rewrite each
// match of it.
type Rule struct {
	// Pattern is matched against the input; every match is rewritten.
	Pattern *regexp.Regexp
	// Replace produces the replacement for each match of Pattern.
	Replace Replacer
}

// Const builds a Rule that replaces every match with a fixed replacement,
// ignoring capture groups.
func Const(pattern *regexp.Regexp, replacement string) Rule {
	return Rule{
		Pattern: pattern,
		Replace: func([]string) string { return replacement },
	}
}

// Apply runs rules against s in order, returning the rewritten string.
//
// Each rule is applied to the whole string (every match) before the next rule
// runs, so later rules see the output of earlier ones. Order therefore matters
// and is the caller's responsibility. An empty input is returned unchanged.
func Apply(s string, rules []Rule) string {
	if s == "" {
		return s
	}
	out := s
	for _, r := range rules {
		// ReplaceAllStringFunc only hands us the whole match, so re-match it to
		// recover the capture groups for the Replacer. Each m is exactly one
		// full match of Pattern, so FindStringSubmatch returns its groups.
		out = r.Pattern.ReplaceAllStringFunc(out, func(m string) string {
			return r.Replace(r.Pattern.FindStringSubmatch(m))
		})
	}
	return out
}
