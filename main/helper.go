package main

import "fmt"

func ClearLine() {
	fmt.Printf("\r\033[K")
}
