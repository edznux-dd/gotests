package foobar

import "testing"

// Run with `go test -fuzz=^FuzzCorge$ -run=^FuzzCorge$`
func FuzzCorge(f *testing.F) {
	// Add "known edge case" values here. It might be a specific payload format in a string,
	// or some other hardcoded values.
	f.Add("hello", 0)
	f.Fuzz(func(t *testing.T, s string, i int) {
		_, _, err := Corge(s, i)
		if err != nil {
			// Ignore handled errors
			return
		}
		return
	})
}
