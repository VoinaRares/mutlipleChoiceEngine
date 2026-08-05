package dialog

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Entry() {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Println("Welcome to the hotel") // We will need to extract the txt out of this and use some nodes to start representing this
		fmt.Println("Options: ")
		option, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Please use a valid number.")
		}

		option = strings.TrimSpace(option)
		chosen, _ := strconv.ParseInt(option, 10, 64)

		if chosen == 0 {
			break
		}
	}
}
