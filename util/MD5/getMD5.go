package MD5

import (
	"crypto/md5"
	"encoding/hex"
)

func CalcMD5(content []byte) string {
	ret := md5.New()
	ret.Write(content)
	return hex.EncodeToString(ret.Sum(nil))
}
