package main

import (
	"fmt"
	"github.com/cybersecshop/gopeekatyou/winmon"
	log "github.com/sirupsen/logrus"
)

func main() {

	log.SetLevel(log.DebugLevel)

	// Get a greeting message and print it.
	message := winmon.Hello("Mrs. Robinson")
	fmt.Println(message)

	log.Info("Let's install the monitoring agent...")
	installReturnCode, err := winmon.Install()
	if installReturnCode == 0 {
		log.Info("Installation completed.")
	} else {
		log.Error("Installation failed! Because " + err.Error())
	}

}
