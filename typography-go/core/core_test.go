package core

import (
	"regexp"
	"testing"
)

func TestApplyConstAndGroups(t *testing.T) {
	rules := []Rule{
		Const(regexp.MustCompile(`a`), "A"),
		{
			Pattern: regexp.MustCompile(`(\d)x`),
			Replace: func(g []string) string { return g[1] + "X" },
		},
	}
	if got, want := Apply("a 1x a", rules), "A 1X A"; got != want {
		t.Errorf("Apply = %q, want %q", got, want)
	}
}

func TestApplyEmpty(t *testing.T) {
	if got := Apply("", nil); got != "" {
		t.Errorf("Apply(\"\") = %q, want empty", got)
	}
}

func TestNbspConstants(t *testing.T) {
	if Nbsp != "\u00A0" || NarrowNbsp != "\u202F" {
		t.Fatalf("NBSP constants drifted: %q %q", Nbsp, NarrowNbsp)
	}
}
