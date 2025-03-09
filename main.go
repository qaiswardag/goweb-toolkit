package main

import (
	"fmt"

	tools "github.com/qaiswardag/goweb-toolkit/toolkit"
)

func main() {
	w := tools.Tools{}

	fmt.Println(w.RandomWords(2))
}
