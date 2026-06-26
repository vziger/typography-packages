package en

import "testing"

func TestEnIsIdentity(t *testing.T) {
	cases := []string{"", "hello world", "3 items", "Box Score"}
	for _, in := range cases {
		if got := En(in); got != in {
			t.Errorf("En(%q) = %q, want %q", in, got, in)
		}
	}
}
