package toolkit

import "testing"

func TestToolsRandomWords(t *testing.T) {
	testTools := Tools{}

	w := testTools.RandomWords(20)

	if len(w) != 20 {
		t.Error("Wrong length random words returned.")
	}
}
