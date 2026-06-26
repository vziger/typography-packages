import 'package:test/test.dart';
import 'package:typography_core/typography_core.dart';

void main() {
  test('nbsp constants have the expected code points', () {
    expect(nbsp.codeUnitAt(0), 0x00A0);
    expect(narrowNbsp.codeUnitAt(0), 0x202F);
  });

  test('applyRules returns empty input unchanged', () {
    expect(applyRules('', [Rule.constant(RegExp('a'), 'b')]), '');
  });

  test('applyRules applies rules in order', () {
    final rules = [
      Rule.constant(RegExp('a'), 'b'),
      Rule.constant(RegExp('b'), 'c'),
    ];
    // First rule turns a->b, then second turns every b (incl. new ones)->c.
    expect(applyRules('a b', rules), 'c c');
  });

  test('Rule with a Replacer can use capture groups', () {
    final rule = Rule(RegExp(r'(\d+)px'), (m) => '${m.group(1)}rem');
    expect(applyRules('12px', [rule]), '12rem');
  });
}
