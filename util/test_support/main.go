package test_support

import (
	"strings"
	"testing"
)

func ExpectEqual(t *testing.T, expected string, actual string) {
	if strings.TrimSpace(expected) != strings.TrimSpace(actual) {
		t.Errorf("Expected does not match actual, %s != %s", expected, actual)
	}
}

func ExpectEqualInt(t *testing.T, expected int, actual int) {
	if expected != actual {
		t.Errorf("Expected does not match actual, %d != %d", expected, actual)
	}
}

func ExpectNotEmpty(t *testing.T, value *map[string]string, valueName string) {
	if len(*value) == 0 {
		t.Errorf("%s should not be empty", valueName)
	}
}
