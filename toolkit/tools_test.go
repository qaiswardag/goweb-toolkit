package toolkit

import "testing"

func TestTools_RandomString(t *testing.T) {
	testTools := Tools{}

	s := testTools.RandomString(20)

	if len(s) != 20 {
		t.Error("Wrong length random string returned.")
	}
}
