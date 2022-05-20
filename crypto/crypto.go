// @Time    : 2022/5/19 8:14 下午
// @Author  : HuYuan
// @File    : crypto.go
// @Email   : huyuan@virtaitech.com

package crypto

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io"
	"os"
)

func Md5sumWithFile(filename string) (string, error) {
	f, err := os.Open(filename)
	if err != nil {
		return "", err
	}
	defer f.Close()
	md5hash := md5.New()
	if _, err := io.Copy(md5hash, f); err != nil {
		return "", err
	}
	return fmt.Sprintf("%x", md5hash.Sum(nil)), nil
}

func Md5sumWithString(s string) string {
	ret := md5.Sum([]byte(s))
	return hex.EncodeToString(ret[:])
}
