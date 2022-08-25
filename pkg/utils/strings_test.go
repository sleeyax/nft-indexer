package utils

import "testing"

func TestToSearchFriendly(t *testing.T) {
	if ToSearchFriendly("abc-def_hello world") != "abcdefhelloworld" {
		t.Fail()
	}
}
