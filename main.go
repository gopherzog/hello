package main

import "fmt"
import myx "github.com/mahler6/mystringutils"

func main() {
	msg := "Hello, New World"
	fmt.Println(myx.Upper(msg))
	fmt.Println(msg)
}
