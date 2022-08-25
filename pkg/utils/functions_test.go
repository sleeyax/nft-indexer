package utils

import "testing"

func TestOrString(t *testing.T) {
	if OrString("", "", "", "abc", "") != "abc" {
		t.Fail()
	}
}

func TestTernary(t *testing.T) {
	if Ternary(true, "foo", "bar") != "foo" {
		t.Fail()
	}

	if Ternary(false, "foo", "bar") != "bar" {
		t.Fail()
	}
}
