package static

import (
	"os"
	"strconv"
)

const SANDBOX_USER = "sandbox"

var SANDBOX_USER_UID = func() int {
	uidStr := os.Getenv("SANDBOX_USER_UID")
	uid, err := strconv.Atoi(uidStr)
	if err != nil {
		return 2000651600 // default value if environment variable is not set
	}
	return uid
}()

var SANDBOX_GROUP_ID = 0
