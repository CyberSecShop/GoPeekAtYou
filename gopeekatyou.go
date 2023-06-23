package main

import (
	"fmt"
	"github.com/cybersecshop/gopeekatyou/winmon"
	log "github.com/sirupsen/logrus"
)

func main() {
	// Get a greeting message and print it.
	message := winmon.Hello("Mrs. Robinson")
	fmt.Println(message)

	log.Println("Let's install the monitoring agent...")
	winmon.Install()
}
