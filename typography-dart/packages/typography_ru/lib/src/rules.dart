/// Regex-based Russian micro-typography rules (spec rules 1–5).
///
/// Source of truth for the rules:
/// `/Users/muntu/.cursor/rules/ui-typography-ru.mdc`.
///
/// Rule 6 (quotes) is stateful and lives in `quotes.dart`; it is applied by the
/// `ru()` facade before these rules so that «/„ are present as opening context
/// for the short-word rule.
library;

import 'package:typography_core/typography_core.dart';

/// Currency symbols handled by the currency rule (rule 4).
const String _currency = r'€₽$£¥';

/// Ordered list of regex rules applied by `ru()` after quote normalization.
///
/// Order notes:
/// - The em-dash and number rules are independent of each other.
/// - The short-word rule runs last; its look-behind includes «/„ so a short
///   word directly after an opening quote also gets a non-breaking space.
final List<Rule> russianRules = [
  // Rule 1: NBSP before an em dash «—». Any whitespace run before the dash
  // collapses to a single NBSP. (Behavior preserved from the reference.)
  Rule.constant(RegExp(r'\s+—'), '$nbsp—'),

  // Rule 3: narrow NBSP (U+202F) between a number and the percent sign.
  Rule.constant(RegExp(r'(?<=\d)\s*%'), '$narrowNbsp%'),

  // Rule 4a: narrow NBSP between a number and a trailing currency sign
  // (e.g. `10₽` / `10 €` → `10 ₽`). Order is preserved.
  Rule(
    RegExp('(?<=\\d)\\s*([$_currency])'),
    (m) => '$narrowNbsp${m.group(1)}',
  ),
  // Rule 4b: narrow NBSP between a leading currency sign and a number
  // (e.g. `€10` / `€ 10` → `€ 10`).
  Rule(
    RegExp('([$_currency])\\s*(?=\\d)'),
    (m) => '${m.group(1)}$narrowNbsp',
  ),

  // Rule 5: NBSP (U+00A0) between a numeral and a dependent noun, Cyrillic or
  // Latin (e.g. `10 дней`, `3 items`). A whitespace run collapses to one NBSP.
  Rule.constant(RegExp(r'(?<=\d)\s+(?=[a-zA-Zа-яёА-ЯЁ])'), nbsp),

  // Rule 2: NBSP after every 1–2 letter Cyrillic word. The word must be at the
  // start of the string or preceded by whitespace / an opening bracket or
  // quote; a fragment after a hyphen (какой-то) is therefore excluded.
  // (Regex preserved from the reference, extended with „ as opening context.)
  Rule.constant(RegExp(r'(?<=(^|[\s(«„"])[а-яёА-ЯЁ]{1,2}) '), nbsp),
];
