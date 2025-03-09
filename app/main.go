package main

import (
	"fmt"

	tools "github.com/qaiswardag/goweb-toolkit/toolkit"
)

func main() {
	s := tools.Tools{}

	fmt.Println(s.RandomString(20))
}
