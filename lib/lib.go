package lib

import "os"

func CWD() string {
	cwd, cwderr := os.Getwd()
	if cwderr != nil {
		panic("An error occurred on getting cwd")
	} else {
		return cwd
	}
}