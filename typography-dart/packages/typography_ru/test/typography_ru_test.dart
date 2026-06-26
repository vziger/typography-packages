import 'package:test/test.dart';
import 'package:typography_ru/typography_ru.dart';

const nbsp = '\u00A0';
const nnbsp = '\u202F';

void main() {
  // --- Rules 1–2: ported verbatim from the reference test suite (9 cases) ---
  group('rules 1–2 (short words & em dash)', () {
    test('NBSP после однобуквенного слова', () {
      expect(ru('счёт в очереди'), 'счёт в$nbspочереди');
    });

    test('NBSP после двухбуквенного слова', () {
      expect(ru('из кэша'), 'из$nbspкэша');
    });

    test('NBSP перед длинным тире', () {
      expect(ru('Соколы — utility'), 'Соколы$nbsp— utility');
    });

    test('короткое слово в начале строки', () {
      expect(ru('в команде'), 'в$nbspкоманде');
    });

    test('длинные слова не трогаются', () {
      expect(ru('Аналитика команды'), 'Аналитика команды');
    });

    test('латиница (бренды) — обычные пробелы', () {
      expect(ru('Box Score'), 'Box Score');
    });

    test('цепочка коротких слов', () {
      expect(ru('я и ты'), 'я$nbspи$nbspты');
    });

    test('пустая строка', () {
      expect(ru(''), '');
    });

    test('комбинация: короткое слово и тире', () {
      expect(ru('Г — гол · П — потеря'), 'Г$nbsp— гол · П$nbsp— потеря');
    });

    test('дефисное слово: суффикс после дефиса не считается словом', () {
      expect(ru('какой-то текст'), 'какой-то текст');
      expect(ru('по-настоящему важно'), 'по-настоящему важно');
    });
  });

  // --- Rule 3: narrow NBSP before % ---
  group('rule 3 (percent)', () {
    test('без пробела', () {
      expect(ru('скидка 10%'), 'скидка 10$nnbsp%');
    });

    test('с обычным пробелом исправляется на узкий NBSP', () {
      expect(ru('10 %'), '10$nnbsp%');
    });

    test('идемпотентность', () {
      expect(ru(ru('10%')), '10$nnbsp%');
    });
  });

  // --- Rule 4: narrow NBSP between number and currency, order preserved ---
  group('rule 4 (currency)', () {
    test('знак после числа', () {
      expect(ru('10€'), '10$nnbsp€');
      expect(ru('10 €'), '10$nnbsp€');
      expect(ru('10₽'), '10$nnbsp₽');
    });

    test('знак перед числом, порядок сохраняется', () {
      expect(ru('€10'), '€${nnbsp}10');
      expect(ru('€ 10'), '€${nnbsp}10');
    });

    test('идемпотентность', () {
      expect(ru(ru('10₽')), '10$nnbsp₽');
    });
  });

  // --- Rule 5: NBSP between numeral and dependent noun ---
  group('rule 5 (numeral + noun)', () {
    test('10 дней', () {
      expect(ru('осталось 10 дней'), 'осталось 10$nbspдней');
    });

    test('5 яблок', () {
      expect(ru('5 яблок'), '5$nbspяблок');
    });

    test('английское слово после числа тоже получает NBSP', () {
      expect(ru('3 items'), '3${nbsp}items');
      expect(ru('10 px'), '10${nbsp}px');
    });
  });

  // --- Rule 6: quotes ---
  group('rule 6 (quotes)', () {
    test('первый уровень — ёлочки', () {
      expect(ru('"текст"'), '«текст»');
    });

    test('вложенный уровень — немецкие лапки', () {
      expect(
        ru('"уровень один "уровень два" снова один"'),
        '«уровень один „уровень два“ снова один»',
      );
    });

    test('две независимые цитаты', () {
      expect(ru('сказал "да" и "нет"'), 'сказал «да» и$nbsp«нет»');
    });

    test('короткое слово сразу после открывающей ёлочки', () {
      expect(ru('"в кавычки"'), '«в$nbspкавычки»');
    });
  });
}
