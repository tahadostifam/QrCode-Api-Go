package lib

import (
	"fmt"
	"os"
)

func CWD() string {
	cwd, cwderr := os.Getwd()
	if cwderr != nil {
		fmt.Println("An error occurred on getting cwd")
		os.Exit(1)
		return ""
	} else {
		return cwd
	}
}
