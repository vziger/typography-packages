/// The rule-application engine.
library;

import 'rule.dart';

/// Applies [rules] to [input] in order, returning the rewritten string.
///
/// Each rule is applied to the whole string (every match) before the next rule
/// runs, so later rules see the output of earlier ones. Order therefore
/// matters and is the caller's responsibility. An empty [input] is returned
/// unchanged.
String applyRules(String input, List<Rule> rules) {
  if (input.isEmpty) return input;
  var out = input;
  for (final rule in rules) {
    out = out.replaceAllMapped(rule.pattern, rule.replace);
  }
  return out;
}
