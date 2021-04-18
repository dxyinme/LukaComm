package MD5

import (
	"bufio"
	"crypto/md5"
	"encoding/hex"
	"io"
	"os"
)

func CalcMD5ToByte(content []byte) []byte {
	ret := md5.New()
	ret.Write(content)
	return ret.Sum(nil)
}

func CalcMD5(content []byte) string {
	ret := md5.New()
	ret.Write(content)
	return hex.EncodeToString(ret.Sum(nil))
}

const bufferSize = 1 << 16


func CalcMD5File(filename string) (string,error) {
	if info, err := os.Stat(filename); err != nil {
		return "", err
	} else if info.IsDir() {
		return "", nil
	}

	file, err := os.Open(filename)
	if err != nil {
		return "", err
	}
	defer file.Close()

	hash := md5.New()
	for buf, reader := make([]byte, bufferSize), bufio.NewReader(file); ; {
		n, err := reader.Read(buf)
		if err != nil {
			if err == io.EOF {
				break
			}
			return "", err
		}

		hash.Write(buf[:n])
	}
	return hex.EncodeToString(hash.Sum(nil)), nil
}