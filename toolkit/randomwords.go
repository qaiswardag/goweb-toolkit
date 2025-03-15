package toolkit

import (
	"math/rand"
	"strings"
	"time"
)

// Sample words for generating Lorem Ipsum text
var randomWords = []string{
	"amet", "consectetur", "adipiscing", "elit",
	"nunc", "potenti", "felis", "rhoncus", "sodales", "arcu", "eleifend", "ullamcorper",
	"tristique", "nisi", "platea", "pellentesque", "aptent", "risus", "odio", "tincidunt",
	"ante", "sem", "dolor", "parturient", "egestas", "nisl", "purus", "habitasse",
	"magna", "hac", "luctus", "sapien", "turpis", "orci", "pharetra",
	"taciti", "inceptos", "donec", "eros", "imperdiet", "volutpat",
	"morbi", "lectus", "leo", "vestibulum", "posuere", "suspendisse",
	"aurora", "cherish", "crystalline", "dulcet", "enchant", "ephemeral", "ethereal",
	"euphoria", "felicity", "halcyon", "idyllic", "incandescent", "labyrinth", "luminous",
	"opalescent", "petrichor", "sequoia", "serendipity", "solitude", "supine", "volutpat",
	"lacinia", "metus", "proin", "sodales", "pretium", "en",
	"plethora", "vellichor", "blossoming", "panacea", "languor", "limerence",
	"am", "on", "is", "og",
}

// Tools struct
type Tools struct{}

// capitalizeFirstLetter capitalizes the first letter of a string
func capitalizeFirstLetter(s string) string {
	if len(s) == 0 {
		return s
	}
	return strings.ToUpper(string(s[0])) + s[1:]
}

func (t *Tools) RandomWords(n int) string {
	src := rand.NewSource(time.Now().UnixNano())
	r := rand.New(src)

	words := []string{}
	for range make([]int, n) {
		words = append(words, randomWords[r.Intn(len(randomWords))])
	}

	// Capitalize the first character of the first word
	if len(words) > 0 {
		words[0] = capitalizeFirstLetter(words[0])

	}

	// Insert dots randomly after every 10 to 16 words and capitalize the first character after each dot
	for i := 10; i < len(words); i += r.Intn(7) + 10 {
		if i < len(words) {
			words[i] += "."
			if i+1 < len(words) {
				words[i+1] = capitalizeFirstLetter(words[i+1])
			}
		}
	}

	return strings.Join(words, " ") + "."
}
