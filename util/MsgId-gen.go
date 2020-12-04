package util

import (
	"fmt"
	"time"
)

func bindStringInt(s string,v int64) string {
	return s + fmt.Sprintf("%063b", v)
}

func MsgIdGen(UID string, registerTime, lastMsgTime, tip int64) (nowTime,nowTip int64, res string) {
	nowTime = time.Now().Unix()
	if nowTime == lastMsgTime {
		nowTip = tip + 1
	} else {
		nowTip = 0
	}
	res = bindStringInt(UID, (nowTime - registerTime) << 12 | tip)
	return nowTime, nowTip, res
}
