package main

import (
	"fmt"

	tools "github.com/qaiswardag/goweb-toolkit/toolkit"
)

func main() {
	tools := tools.Tools{}
	fmt.Println("Hello")
	tools.EnsureDirCreated("qw")
}
