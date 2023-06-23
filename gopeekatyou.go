package main

import (
	"github.com/cybersecshop/gopeekatyou/winmon"
	log "github.com/sirupsen/logrus"
	"strconv"
)

func main() {

	log.SetLevel(log.DebugLevel)

	const name = "mylog"

	ReturnCode, err := winmon.Start(name)
	if ReturnCode == 0 {
		log.Info("completed.")
	} else {
		log.Error("Failed! Because " + err.Error() +
			". (Return Code: " + strconv.Itoa(ReturnCode) + ")")
	}

}
