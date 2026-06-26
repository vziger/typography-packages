/// Russian quote normalization (rule 6 of the spec).
///
/// Converts straight double quotes `"..."` into typographic quotes:
/// first level → guillemets «», nested level → German low-high „“.
///
/// This is a small stateful scanner rather than a single regex, because the
/// same character `"` serves as both opener and closer and the correct glyph
/// depends on nesting depth and the preceding character.
library;

final RegExp _openingContext = RegExp(r'[\s(\[«„]');

/// Whether a `"` preceded by [prev] should be treated as an opening quote.
///
/// A quote opens at the start of the string or after whitespace / an opening
/// bracket / an already-open quote; otherwise it closes.
bool _opensAfter(String prev) => prev.isEmpty || _openingContext.hasMatch(prev);

/// Replaces straight double quotes with Russian typographic quotes.
///
/// First-level pairs become «…», nested pairs become „…“. Text without a `"`
/// is returned unchanged. Unbalanced quotes are handled defensively: a `"`
/// while nothing is open always opens.
String applyRussianQuotes(String input) {
  if (!input.contains('"')) return input;

  final sb = StringBuffer();
  var depth = 0;
  for (var i = 0; i < input.length; i++) {
    final ch = input[i];
    if (ch != '"') {
      sb.write(ch);
      continue;
    }

    final prev = i == 0 ? '' : input[i - 1];
    final isOpener = depth == 0 || _opensAfter(prev);
    if (isOpener) {
      sb.write(depth == 0 ? '«' : '„');
      depth++;
    } else if (depth >= 2) {
      sb.write('“');
      depth--;
    } else {
      sb.write('»');
      depth = 0;
    }
  }
  return sb.toString();
}
