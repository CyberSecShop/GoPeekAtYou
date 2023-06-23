package winmon

import (
	"encoding/json"
	"errors"
	"fmt"
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
	handle, err := Open(name)
	if err != nil {
		log.Error("Monitoring start failed! Because " + err.Error())
		return -2, err
	} else {
		handleStr := fmt.Sprint(uintptr(handle))
		log.Info("Monitoring started. Event log handle: " + handleStr)

		json, fetchErr := FetchInfo(handle)
		if fetchErr != nil {
			log.Error("Fetching info failed! Because " + fetchErr.Error())
			return -3, fetchErr
		} else {
			log.Debug("Info retrieved:\n%s", json) //TODO testing!
		}
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

	log.Debug("Installation successful")
	return 0, nil // All good, return error code 0
}

func Open(name string) (windows.Handle, error) {
	log.Debug("Opening Event Log handler for " + name)

	handle, err := eventlog.Open(name)
	if err != nil || handle == nil {
		log.Warn("Failed to start monitoring for " + name)
		return handle.Handle, errors.New(err.Error())
	}

	log.Debug("Opened Event Log handler for " + name)
	return handle.Handle, nil // All good, return error code 0
}

func FetchInfo(handle windows.Handle) ([]byte, error) {
	log.Debug("Fetching info for Event Log handler " + fmt.Sprint(uintptr(handle)))

	// TODO testing
	raw := map[string]interface{}{
		"intValue":    1234,
		"boolValue":   true,
		"stringValue": "hello!",
		"objectValue": map[string]interface{}{
			"arrayValue": []int{1, 2, 3, 4},
		},
	}

	json, err := json.Marshal(raw)
	if err != nil {
		log.Warn("could not marshal json: %s\n", err)
		return nil, err
	}

	log.Debug("Done fetching info for Event Log handler " + fmt.Sprint(uintptr(handle)))
	return json, nil // All good, return error code 0
}
