package toolkit

import (
	"math/rand"
	"strings"
	"time"
)

// Sample words for generating Lorem Ipsum text
var loremWords = []string{
	"lorem", "ipsum", "dolor", "sit", "amet", "consectetur", "adipiscing", "elit",
	"nunc", "potenti", "felis", "rhoncus", "sodales", "arcu", "eleifend", "ullamcorper",
	"tristique", "nisi", "platea", "pellentesque", "aptent", "risus", "odio", "tincidunt",
	"ante", "sem", "dolor", "parturient", "egestas", "nisl", "purus", "habitasse",
	"magna", "hac", "luctus", "sapien", "turpis", "facilisis", "orci", "pharetra",
	"taciti", "inceptos", "donec", "facilisi", "eros", "imperdiet", "volutpat",
	"morbi", "lectus", "leo", "vestibulum", "posuere", "quam", "suspendisse", "class",
	"lacinia", "metus", "laoreet", "proin", "tempor", "sodales", "pretium",
}

// Tools struct
type Tools struct{}

// RandomWords generates a Lorem Ipsum-like text with n words
func (t *Tools) RandomWords(n int) string {
	src := rand.NewSource(time.Now().UnixNano()) // Create a new random source
	r := rand.New(src)                           // Create a new Rand instance

	var words []string
	for i := 0; i < n; i++ {
		words = append(words, loremWords[r.Intn(len(loremWords))])
	}

	return strings.Join(words, " ")
}
