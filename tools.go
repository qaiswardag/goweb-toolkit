package tools

import "fmt"

func AddHelper(a, b int) int {
	fmt.Println("This is running from tools package. Result is:", a+b)
	return a + b
}
