/// English UI micro-typography.
///
/// Minimal scaffold built on `typography_core`, included to demonstrate how a
/// new language package is structured: define an ordered `List<Rule>` and run
/// it through [applyRules] in a small facade. With no rules defined yet, [en]
/// is an identity passthrough. Pure Dart, no Flutter.
library;

import 'package:typography_core/typography_core.dart';

import 'src/rules.dart';

export 'src/rules.dart' show englishRules;

/// Returns [s] with English UI micro-typography applied.
///
/// Currently an identity passthrough (no rules defined). Add `Rule`s to
/// [englishRules] to enable transformations.
String en(String s) {
  if (s.isEmpty) return s;
  return applyRules(s, englishRules);
}
