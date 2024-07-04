package tests

import "testing"

func AssertStrings(t *testing.T, expected, actual string) {
	if expected != actual {
		t.Fatalf("Expected: %s, actual: %s", expected, actual)
	}
}

func AssertInt(t *testing.T, expected, actual int) {
	if expected != actual {
		t.Fatalf("Expected: %d, actual: %d", expected, actual)
	}
}

func AssertNoError(t *testing.T, err error) {
	if err != nil {
		t.Fatal(err)
	}
}
