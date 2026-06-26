/// A single micro-typography rule: a pattern plus how to rewrite each match.
library;

/// Builds the replacement string for one [Match].
///
/// Receives the full match (with capture groups available via [Match.group])
/// and returns the text that should replace it.
typedef Replacer = String Function(Match match);

/// A micro-typography rule.
///
/// A rule is a [RegExp] [pattern] and a [replace] function applied to every
/// match. Rules are intentionally tiny and composable: a language package is
/// just an ordered `List<Rule>` fed to `applyRules`.
class Rule {
  /// Pattern matched against the input string.
  final RegExp pattern;

  /// Produces the replacement for each match of [pattern].
  final Replacer replace;

  /// Creates a rule from a [pattern] and a per-match [replace] function.
  const Rule(this.pattern, this.replace);

  /// Convenience rule that replaces every match with a constant [replacement].
  ///
  /// Unlike [String.replaceAll], `$1` style placeholders are not interpreted;
  /// for group-aware replacements use the main [Rule] constructor with a
  /// [Replacer].
  Rule.constant(RegExp pattern, String replacement)
      : this(pattern, ((_) => replacement));
}
