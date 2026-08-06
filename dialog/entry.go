package dialog

import (
	"bufio"
	"errors"
	"fmt"
	"learningProject/traversal"
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

	err := traversal.Initialize(resolveDialogPath())
	if err != nil {
		fmt.Print("Error initializing Dialog Tree: ", err)
		return
	}

	for {
		currentOption := traversal.GetCurrent()
		options := traversal.GetOptions(currentOption.Id)
		displayPreviews(options)

		newOptionIndex, err := readOption()
		if err != nil {
			fmt.Println("error:", err)
			continue
		}
		err = validateOption(newOptionIndex, options)
		if err != nil {
			fmt.Println("Invalid option", err)
			continue
		}
		newOptionIndex-- //Decrement to use as index in slice
		newOptionId := options[newOptionIndex].Id
		err = traversal.ChooseOption(newOptionId)
	}

}

func readOption() (int, error) {
	reader := bufio.NewReader(os.Stdin)

	newOptionIndexStr, err := reader.ReadString('\n')
	if err != nil {
		return 0, errors.New("error reading options")
	}
	newOptionIndexStr = strings.TrimSpace(newOptionIndexStr)
	newOptionIndex, err := strconv.Atoi(newOptionIndexStr)
	if err != nil {
		return 0, errors.New("error parsing string to option id")
	}

	return newOptionIndex, nil

}
func displayPreviews(options []traversal.Option) {
	fmt.Println("Choose an Option: ")
	for i, option := range options {
		fmt.Printf("%d. "+option.PreviewText+"\n", i+1)
	}
}

func validateOption(option int, options []traversal.Option) error {
	if option < 1 || option > len(options) {
		return errors.New("option is out of range")
	}
	return nil
}
