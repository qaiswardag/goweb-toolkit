package toolkit

import "testing"

func TestToolsRandomWords(t *testing.T) {
	testTools := Tools{}

	rw := testTools.RandomWords(20)

	if len(rw) != 20 {
		t.Error("Wrong length random words returned.")
	}
}
