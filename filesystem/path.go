// @Time    : 2022/5/19 5:44 下午
// @Author  : HuYuan
// @File    : pkg.go
// @Email   : huyuan@virtaitech.com

package filesystem

import "os"

func Exists(path string) bool {
	_, err := os.Stat(path)
	if err != nil {
		if os.IsExist(err) {
			return true
		}
		return false
	}
	return true
}

func IsDir(path string) bool {
	s, err := os.Stat(path)
	if err != nil {
		return false
	}
	return s.IsDir()
}

func IsFile(path string) bool {
	if Exists(path) {
		return !IsDir(path)
	}
	return false
}
