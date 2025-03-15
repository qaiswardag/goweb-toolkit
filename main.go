package main

import (
	"fmt"

	tools "github.com/qaiswardag/goweb-toolkit/toolkit"
)

func main() {
	tools := tools.Tools{}

	rw := tools.RandomWords(12)
	fmt.Println(rw)
}
