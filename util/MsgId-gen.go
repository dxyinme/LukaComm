package util

import (
	"fmt"
	"time"
)

func bindStringInt(s string,v int64) string {
	return s + fmt.Sprintf("%063b", v)
}

func MsgIdGen(UID string, registerTime, lastMsgTime, tip int) (nowTime,nowTip int, res string) {
	nowTime = time.Now().Second()
	if nowTime == lastMsgTime {
		nowTip = tip + 1
	} else {
		nowTip = 0
	}
	res = bindStringInt(UID, int64((nowTime - registerTime) << 12 | tip))
	return nowTime, nowTip, res
}
