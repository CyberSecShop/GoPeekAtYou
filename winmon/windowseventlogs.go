package winmon

import (
	"errors"
	"fmt"
	"github.com/abdfnx/gosh"
	log "github.com/sirupsen/logrus"
	"os/exec"
)

func Exec(command string) ([]byte, error) {

	validationFailure := ValidateCommand(command)

	if validationFailure != nil {
		log.Panicf("Command failed validation!! %s", validationFailure)
		return nil, validationFailure
	}

	cmd := exec.Command(command)

	stdin, err := cmd.StdinPipe()
	if err != nil {
		log.Fatal(err)
	}
	go func() {
		defer stdin.Close()
		log.Warnf("Failed to exit command gracefully.")
	}()
	out, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatalf("Failed to output command: %s", err)
		return nil, err
	}

	return out, nil
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

func ListLogs() string {

	//out, _ := Exec("Get-WinEvent -ListLog *")

	err, out, errout := gosh.PowershellOutput(`Get-WinEvent -LogName 'Application' -MaxEvents 10`)

	if err != nil {
		log.Printf("error: %v\n", err)
		fmt.Print(errout)
	}

	log.Debugf(out)

	return fmt.Sprintf("%s", out)
}
