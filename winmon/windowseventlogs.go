package winmon

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/abdfnx/gosh"
	log "github.com/sirupsen/logrus"
)

func Exec(command string) ([]byte, error) {
	log.Debugf("Executing command: \n %v", command)

	validationFailure := ValidateCommand(command)

	if validationFailure != nil {
		log.Panicf("Command failed validation!! %s", validationFailure)
		return nil, validationFailure
	}

	err, out, errout := gosh.PowershellOutput(command)

	if err != nil {
		log.Errorf("command error: %v\n", err)
		log.Errorf("command output error: %v\n", errout)
	}

	log.Debugf("Command execution complete: \n %s", command)
	return []byte(out), nil
}

func ValidateCommand(command string) error {
	log.Debugf("Validating command(s): %s", command)

	Good := true //TODO short-circuited for dev purposes

	if Good {
		log.Debugf("Command(s) Validated: %s", command)
		return nil
	} else {
		return errors.New("hey! What do you think you're doing?!")
	}
}

func ListLogs(filter string) ([]map[string]interface{}, error) {
	log.Debugf("Listing logs. Filter: \n %v", filter)

	cmd := fmt.Sprintf("Get-WinEvent -ListLog %s", filter)

	// Convert to JSON
	cmd = fmt.Sprintf("%v | ConvertTo-Json", cmd)

	out, err := Exec(cmd)

	if err != nil {
		log.Warnf("Failed to list logs: %v\n", err)
		return nil, err
	}

	return ConvertBytesToJson(out)
}

func ConvertBytesToJson(raw []byte) ([]map[string]interface{}, error) {

	//TODO debugging
	//raw = []byte(`{"num":6.13,"strs":["a","b"]}`)
	//raw = []byte(`[{"num":3.14,"strs":["a","b"]}]`)

	var loggers []map[string]interface{}
	err := json.Unmarshal(raw, &loggers)
	if err != nil && err.Error() == "json: cannot unmarshal object into Go value of type []map[string]interface {}" {
		log.Debugf("Seems like JSON is a single object, placing in array.")

		raw = []byte(fmt.Sprintf("[%s]", raw))
		err = json.Unmarshal(raw, &loggers)
	}

	log.Tracef("Marshalling []bytes to JSON map. Input: %s", loggers)

	if err != nil {
		log.Warnf("Could not marshal json:")
		return nil, err
	}

	return loggers, nil

	// https://gobyexample.com/json
}
