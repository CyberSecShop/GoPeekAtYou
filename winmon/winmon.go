package winmon

import (
	"encoding/json"
	"errors"
	"fmt"
	log "github.com/sirupsen/logrus"
	"golang.org/x/sys/windows/svc/eventlog"
	"strconv"
)

func Start(name string) (int, error) {
	log.Info("Starting monitoring of Windows Event logs.")

	handle, err := GetHandle(name)

	json, fetchErr := FetchInfo(handle)
	if fetchErr != nil {
		log.Error("Fetching info failed! Because " + fetchErr.Error())
		return -3, fetchErr
	} else {
		log.Debug("Info retrieved:\n%s", string(json)) //TODO testing!
	}

	// Before we Go, let's clean up behind ourselves.
	defer func() {
		log.Debug("Closing handler(s).")
		err = eventlog.Remove(name)
		if err != nil {
			log.Warn("Closing handle failed! Because " + err.Error())
		} else {
			log.Debug("Handler(s) closed.")
		}
	}()

	return 0, nil // All good, let's go home
}

func GetHandle(name string) (*eventlog.Log, error) {

	// Need to install the monitoring agent first
	installReturnCode, err := Install(name)
	if installReturnCode == 0 {
		log.Debug("Installation completed.")
	} else {
		log.Warning("Installation failed! Because " + err.Error() +
			". (Return Code: " + strconv.Itoa(installReturnCode) + ")")
	}

	// Let's start monitoring...
	handle, err := Open(name)
	if err != nil {
		log.Error("Failed to get handle! Because " + err.Error())
		return nil, err
	} else {
		handleStr := fmt.Sprint(uintptr(handle.Handle))
		log.Debug("Obtained Event log handler: " + handleStr)
	}

	return handle, nil
}

func Install(name string) (int, error) {
	log.Debug("Installing winmon agent...")

	const supports = eventlog.Error | eventlog.Warning | eventlog.Info

	err := eventlog.InstallAsEventCreate(name, supports)

	if err != nil {
		log.Warn("Blast!! Something went wrong during installation!")
		return -1, errors.New(err.Error())
	}

	log.Debug("Installation successful")
	return 0, nil // All good, return error code 0
}

func Open(name string) (*eventlog.Log, error) {
	log.Debug("Opening Event Log handler for " + name)

	handle, err := eventlog.Open(name)
	if err != nil || handle == nil {
		log.Warn("Failed to start monitoring for " + name)
		return handle, errors.New(err.Error())
	}

	log.Debug("Opened Event Log handler for " + name)
	return handle, nil // All good, return error code 0
}

func FetchInfo(handle *eventlog.Log) ([]byte, error) {
	log.Debug("Fetching info for Event Log handler " + fmt.Sprint(uintptr(handle.Handle)))

	//ComputerName := handle.Get

	// TODO testing
	raw := map[string]interface{}{}

	json, err := json.Marshal(raw)
	if err != nil {
		log.Warn("could not marshal json: %s\n", err)
		return nil, err
	}

	log.Debug("Done fetching info for Event Log handler " + fmt.Sprint(uintptr(handle.Handle)))
	return json, nil // All good, return error code 0
}
