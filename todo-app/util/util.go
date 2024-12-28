package util

import "os"

func ReadOnly() bool {
	return os.Getenv("ALLOW_WRITE") == ""
}
