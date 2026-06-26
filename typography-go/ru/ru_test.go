package ru

import "testing"

// Non-breaking spaces as escapes, never as pasted literals (a raw U+00A0 /
// U+202F in source is silently normalized to 0x20 by some tooling).
const (
	nbsp  = "\u00A0"
	nnbsp = "\u202F"
)

// cases ports the 23 reference cases from the Dart suite
// (typography-dart/.../test/typography_ru_test.dart) 1:1. Multi-assert Dart
// tests are expanded into one row each. Results must match byte-for-byte.
var cases = []struct {
	name string
	in   string
	want string
}{
	// --- Rules 1–2: short words & em dash (ported verbatim) ---
	{"NBSP после однобуквенного слова", "счёт в очереди", "счёт в" + nbsp + "очереди"},
	{"NBSP после двухбуквенного слова", "из кэша", "из" + nbsp + "кэша"},
	{"NBSP перед длинным тире", "Соколы — utility", "Соколы" + nbsp + "— utility"},
	{"короткое слово в начале строки", "в команде", "в" + nbsp + "команде"},
	{"длинные слова не трогаются", "Аналитика команды", "Аналитика команды"},
	{"латиница (бренды) — обычные пробелы", "Box Score", "Box Score"},
	{"цепочка коротких слов", "я и ты", "я" + nbsp + "и" + nbsp + "ты"},
	{"пустая строка", "", ""},
	{"комбинация: короткое слово и тире", "Г — гол · П — потеря", "Г" + nbsp + "— гол · П" + nbsp + "— потеря"},
	{"дефисное слово: какой-то", "какой-то текст", "какой-то текст"},
	{"дефисное слово: по-настоящему", "по-настоящему важно", "по-настоящему важно"},

	// --- Rule 3: narrow NBSP before % ---
	{"процент без пробела", "скидка 10%", "скидка 10" + nnbsp + "%"},
	{"процент с обычным пробелом", "10 %", "10" + nnbsp + "%"},

	// --- Rule 4: narrow NBSP between number and currency, order preserved ---
	{"валюта после числа без пробела", "10€", "10" + nnbsp + "€"},
	{"валюта после числа с пробелом", "10 €", "10" + nnbsp + "€"},
	{"валюта рубль после числа", "10₽", "10" + nnbsp + "₽"},
	{"валюта перед числом без пробела", "€10", "€" + nnbsp + "10"},
	{"валюта перед числом с пробелом", "€ 10", "€" + nnbsp + "10"},

	// --- Rule 5: NBSP between numeral and dependent noun ---
	{"10 дней", "осталось 10 дней", "осталось 10" + nbsp + "дней"},
	{"5 яблок", "5 яблок", "5" + nbsp + "яблок"},
	{"английское слово после числа: items", "3 items", "3" + nbsp + "items"},
	{"английское слово после числа: px", "10 px", "10" + nbsp + "px"},

	// --- Rule 6: quotes ---
	{"кавычки первый уровень — ёлочки", `"текст"`, "«текст»"},
	{"кавычки вложенный уровень", `"уровень один "уровень два" снова один"`, "«уровень один „уровень два“ снова один»"},
	{"две независимые цитаты", `сказал "да" и "нет"`, "сказал «да» и" + nbsp + "«нет»"},
	{"короткое слово сразу после ёлочки", `"в кавычки"`, "«в" + nbsp + "кавычки»"},
}

func TestRu(t *testing.T) {
	for _, c := range cases {
		if got := Ru(c.in); got != c.want {
			t.Errorf("%s: Ru(%q) = %q, want %q", c.name, c.in, got, c.want)
		}
	}
}

// TestRuIdempotent asserts Ru(Ru(x)) == Ru(x) for every case: applying the
// rules again must not change an already-typeset string.
func TestRuIdempotent(t *testing.T) {
	for _, c := range cases {
		once := Ru(c.in)
		if twice := Ru(once); twice != once {
			t.Errorf("%s: not idempotent: Ru(Ru(%q)) = %q, want %q", c.name, c.in, twice, once)
		}
	}
}
