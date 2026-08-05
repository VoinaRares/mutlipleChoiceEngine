package dialog

import (
	"awesomeProject/traversel"
	"bufio"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
)

const DIALOG_PATH = "dialog.json"

func resolveDialogPath() string {
	_, callerFile, _, ok := runtime.Caller(0)
	if !ok {
		return DIALOG_PATH
	}

	candidate := filepath.Clean(filepath.Join(filepath.Dir(callerFile), "..", DIALOG_PATH))
	if _, err := os.Stat(candidate); err == nil {
		return candidate
	}

	return DIALOG_PATH
}

func Run() {

	err := traversel.Initialize(resolveDialogPath())
	if err != nil {
		fmt.Print("Error intiliazing Dialog Tree")
		return
	}
	reader := bufio.NewReader(os.Stdin)

	for {
		currentOption := traversel.GetCurrent()
		options := traversel.GetOptions(currentOption.Id)
		displayOptions(options)
		newOptionIndexStr, err := reader.ReadString('\n')
		if err != nil {
			fmt.Print("Error reading options")
			return
		}

		newOptionIndexStr = strings.TrimSpace(newOptionIndexStr)
		newOptionIndex, err := strconv.Atoi(newOptionIndexStr)
		if err != nil {
			fmt.Print("Error parsing string to option id")
			continue
		}
		err = validateOption(newOptionIndex, options)
		if err != nil {
			fmt.Println("Invalid option")
			continue
		}
		newOptionIndex-- //Decrement to use as index in slice
		newOptionId := options[newOptionIndex].Id
		err = traversel.ChooseOption(newOptionId)
	}

}

func displayOptions(options []traversel.Option) {
	fmt.Println("Choosse an Option: ")
	for i, option := range options {
		fmt.Printf("%d. "+option.Text+"\n", i+1)
	}
}

func validateOption(option int, options []traversel.Option) error {
	if option < 1 || option > len(options) {
		return errors.New("Invalid option")
	}
	return nil
}
