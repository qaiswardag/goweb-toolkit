package tools

import "fmt"

func AddHelper(a, b int) int {
	fmt.Println("Updated package. This is running from tools package. Result is:", a+b)
	return a + b
}
