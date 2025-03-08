package tools

import "fmt"

const randomStringSource = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789_+"

type Tools struct{}

func (t *Tools) RandomString() {
	fmt.Println("Helooooo. This is a random string.")
}
