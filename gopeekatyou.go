package main

import (
	"github.com/cybersecshop/gopeekatyou/winmon"
	log "github.com/sirupsen/logrus"
)

const name = "mylog"

func main() {
	log.SetLevel(log.DebugLevel)

	FetchInfo()

	// Before we Go, let's clean up behind ourselves.
	defer func() {
		log.Debug("Tearing down...")

		log.Debug("Shut down gracefully.")
	}()
}

func FetchInfo() {
	log.Debugf("Fetching info for Event Log handler ")

	err := DescribeLogs("*sec*")

	if err != nil {
		log.Errorf("Failed to describe logs!")
		panic(err.Error())
	}

	log.Debug("Done fetching info for Event Log handler")
}

func DescribeLogs(filter string) error {
	log.Debugf("Describing logs for %s", filter)

	logs, err := winmon.ListLogs(filter)

	if err != nil {
		log.Errorf("Failed to list logs!")
		return err
	} else {
		log.Debugf("Number of logs: %d", len(logs))

		for index, element := range logs {
			log.Debugf("%d) Log name: %s", index, element["LogName"])
		}

		log.Tracef("Logs: %s", logs)
		return nil
	}
}

func GetLogEvents() (map[string]interface{}, error) {
	log.Debugf("Fetching Windows Log events")

	log.Debug("Done fetching info for Event Log handler")
	return nil, nil // All good, return error code 0
}
