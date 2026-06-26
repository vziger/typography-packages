/// Russian UI micro-typography.
///
/// Applies the full set of Russian UI typography rules from the canonical spec
/// `/Users/muntu/.cursor/rules/ui-typography-ru.mdc`:
///
/// 1. NBSP (U+00A0) before an em dash «—».
/// 2. NBSP after every 1–2 letter Cyrillic word (hyphenated fragments excluded).
/// 3. Narrow NBSP (U+202F) between a number and `%`.
/// 4. Narrow NBSP (U+202F) between a number and a currency sign (order kept).
/// 5. NBSP (U+00A0) between a numeral and its dependent noun.
/// 6. Quotes: first level «», nested „“ (straight `"` are converted).
///
/// Pure Dart, no Flutter. The public entry point is [ru].
library;

import 'package:typography_core/typography_core.dart';

import 'src/quotes.dart';
import 'src/rules.dart';

export 'src/quotes.dart' show applyRussianQuotes;
export 'src/rules.dart' show russianRules;

/// Returns [s] with Russian UI micro-typography applied.
///
/// Backwards compatible with the original reference implementation for rules
/// 1–2 (short-word and em-dash non-breaking spaces); additionally applies the
/// percent, currency, numeral-noun and quote rules. An empty string is returned
/// unchanged. Latin text and code are left untouched.
String ru(String s) {
  if (s.isEmpty) return s;
  final quoted = applyRussianQuotes(s);
  return applyRules(quoted, russianRules);
}
