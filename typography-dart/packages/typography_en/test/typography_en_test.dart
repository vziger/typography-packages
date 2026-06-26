import 'package:test/test.dart';
import 'package:typography_en/typography_en.dart';

void main() {
  test('passthrough leaves text unchanged for now', () {
    expect(en('Box Score'), 'Box Score');
    expect(en('10% off'), '10% off');
  });

  test('empty string', () {
    expect(en(''), '');
  });

  test('no rules are defined yet', () {
    expect(englishRules, isEmpty);
  });
}
