package tests

import "testing"

func AssertStrings(t *testing.T, expected, actual string) {
	if expected != actual {
		t.Fatalf("Expected: %s, actual: %s", expected, actual)
	}
}
