package util

import (
	"os"
	"sync"
	"time"
)

var mu_ sync.Mutex

func GetSessionIdSuffix() string {
	mu_.Lock()
	defer mu_.Unlock()
	res := ""
	res += string(B64Encode([]byte(time.Now().String()))) + "_"
	osname, err := os.Hostname()
	if err != nil {
		osname = "Luka_linux"
	}
	res += osname
	return res
}
