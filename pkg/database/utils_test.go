package database

import "testing"

func TestNormalize(t *testing.T) {
	if normalized := Normalize("0xaBA7161A7fb69c88e16ED9f455CE62B791EE4D03"); normalized != "0xaba7161a7fb69c88e16ed9f455ce62b791ee4d03" {
		t.Fail()
	}
}
