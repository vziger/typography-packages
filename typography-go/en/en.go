// Package en is a minimal starting point for English UI micro-typography.
//
// It currently applies no transformations (En is an identity passthrough) and
// exists as a template for adding a new language: give the package an exported
// entry point that builds on the core engine. See the repository README for the
// "how to add a language" guide.
package en

// En returns s unchanged. It is a placeholder so the package compiles and ships
// a stable public API; real English rules can be added later as a []core.Rule
// fed to core.Apply, mirroring the ru package.
func En(s string) string {
	return s
}
