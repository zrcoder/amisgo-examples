package util

import "os"

func ReadOnly() bool {
	return os.Getenv("PROD") == ""
}
