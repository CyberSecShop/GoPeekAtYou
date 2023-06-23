package winmon

import (
	"errors"
	"fmt"
	log "github.com/sirupsen/logrus"
	"golang.org/x/sys/windows/svc/eventlog"
)

// Hello returns a greeting for the named person.
func Hello(name string) string {
	// Return a greeting that embeds the name in a message.
	message := fmt.Sprintf("Hello, %v. Welcome!", name)
	return message
}

func Install() (int, error) {
	log.Debug("Installing winmon agent...")

	const name = "mylog"
	const supports = eventlog.Error | eventlog.Warning | eventlog.Info

	err := eventlog.InstallAsEventCreate(name, supports)

	if err != nil {
		log.Error("Blast!! Something went wrong during installation!")
		return -1, errors.New("installation failed")
	}

	log.Debug("Installation successful")
	return 0, nil // All good, return error code 0
}
