// @Time    : 2022/5/19 8:18 下午
// @Author  : HuYuan
// @File    : crypto_test.go
// @Email   : huyuan@virtaitech.com

package crypto

import "testing"

func TestMd5sumWithFile(t *testing.T) {
	t.Log(Md5sumWithFile("/etc/passwd"))
	t.Log(Md5sumWithFile("/etc/shadow"))
	t.Log(Md5sumWithFile("/etc/fstab"))
}

func TestMd5sumWithString(t *testing.T) {
	t.Log(Md5sumWithString("123"))
	t.Log(Md5sumWithString("123"))
	t.Log(Md5sumWithString("1234"))
	t.Log(Md5sumWithString("1234"))
	t.Log(Md5sumWithString("1235"))
	t.Log(Md5sumWithString("1235"))
}
