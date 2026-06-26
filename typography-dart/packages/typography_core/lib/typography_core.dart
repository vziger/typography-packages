/// Language-agnostic core for micro-typography rule engines.
///
/// Provides:
/// - non-breaking space constants ([nbsp], [narrowNbsp]),
/// - the [Rule] type (a [RegExp] plus a [Replacer]),
/// - the [applyRules] engine that runs an ordered list of rules over a string.
///
/// Language packages (e.g. `typography_ru`) build an ordered `List<Rule>` on
/// top of this and expose a small facade. Pure Dart, no Flutter.
library;

export 'src/engine.dart';
export 'src/nbsp.dart';
export 'src/rule.dart';
