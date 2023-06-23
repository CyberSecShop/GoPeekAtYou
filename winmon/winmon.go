package winmon

import (
	"errors"
	log "github.com/sirupsen/logrus"
	"golang.org/x/sys/windows"
	"golang.org/x/sys/windows/svc/eventlog"
	"strconv"
)

func Start(name string) (int, error) {
	log.Info("Starting monitoring of Windows Event logs.")

	// Need to install the monitoring agent first
	installReturnCode, err := Install(name)
	if installReturnCode == 0 {
		log.Info("Installation completed.")
	} else {
		log.Warning("Installation failed! Because " + err.Error() +
			". (Return Code: " + strconv.Itoa(installReturnCode) + ")")
	}

	// Let's start monitoring...
	_, err = Open(name)
	if err != nil {
		log.Info("Monitoring started. Event log handle: ???")
	} else {
		log.Error("Monitoring start failed! Because " + err.Error())
	}

	return 0, nil // All good, return error code 0
}

func Install(name string) (int, error) {
	log.Debug("Installing winmon agent...")

	const supports = eventlog.Error | eventlog.Warning | eventlog.Info

	err := eventlog.InstallAsEventCreate(name, supports)

	if err != nil {
		log.Warn("Blast!! Something went wrong during installation!")
		return -1, errors.New(err.Error())
	}

	log.Info("Installation successful")
	return 0, nil // All good, return error code 0
}

func Open(name string) (windows.Handle, error) {
	log.Debug("Installing winmon agent...")

	handle, err := eventlog.Open(name)
	if err != nil || handle == nil {
		log.Warn("Failed to start monitoring for " + name)
		return handle.Handle, errors.New(err.Error())
	}

	log.Info("Monitoring started for " + name)
	return handle.Handle, nil // All good, return error code 0
}
