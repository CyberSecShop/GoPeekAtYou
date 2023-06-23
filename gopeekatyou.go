package main

import (
	"encoding/json"
	"fmt"
	"github.com/cybersecshop/gopeekatyou/winmon"
	log "github.com/sirupsen/logrus"
	"golang.org/x/sys/windows/svc/eventlog"
)

func main() {

	log.SetLevel(log.DebugLevel)

	const name = "mylog"

	//ReturnCode, err := winmon.Start(name)
	//if ReturnCode == 0 {
	//	log.Info("Mission Completed.")
	//} else {
	//	log.Error("Failed! Because " + err.Error() +
	//		". (Return Code: " + strconv.Itoa(ReturnCode) + ")")
	//}

	handle, err := winmon.GetHandle(name)
	json, fetchErr := FetchInfo(handle)
	if fetchErr != nil {
		log.Error("Fetching info failed! Because " + fetchErr.Error())
		panic(fetchErr.Error())
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

}

func FetchInfo(handle *eventlog.Log) ([]byte, error) {
	log.Debug("Fetching info for Event Log handler " + fmt.Sprint(uintptr(handle.Handle)))

	//ComputerName := handle.Get

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

	log.Debug("Done fetching info for Event Log handler " + fmt.Sprint(uintptr(handle.Handle)))
	return json, nil // All good, return error code 0
}
