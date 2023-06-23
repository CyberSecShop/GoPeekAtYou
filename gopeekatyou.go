package main

import (
	"fmt"
	"github.com/cybersecshop/gopeekatyou/winmon"
)

func main() {
	// Get a greeting message and print it.
	message := winmon.Hello("Mrs. Robinson")
	fmt.Println(message)
}
