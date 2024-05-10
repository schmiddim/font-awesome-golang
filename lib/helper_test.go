package lib

import "testing"

func TestIdempotent(t *testing.T) {
	one := GetIconForString("fooobar")
	two := GetIconForString("fooobar")
	three := GetIconForString("fooobar")

	if one != two && two != three {
		t.Error("Not idempotent")
	}

}
