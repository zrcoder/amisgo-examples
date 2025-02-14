package util

import "os"

func IsDemo() bool {
	return os.Getenv("DEMO") != ""
}

func IsDev() bool {
	return os.Getenv("DEV") != ""
}
