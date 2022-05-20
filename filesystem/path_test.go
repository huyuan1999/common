// @Time    : 2022/5/19 5:46 下午
// @Author  : HuYuan
// @File    : path_test.go.go
// @Email   : huyuan@virtaitech.com

package filesystem

import "testing"

func TestExists(t *testing.T) {
	t.Log("console", Exists("/dev/console"))
	t.Log("proc", Exists("/proc"))
	t.Log("device: ", Exists("/dev/vda1"))
	t.Log("soft link: ", Exists("/etc/localtime"))     // soft link
	t.Log("invalid soft link: ", Exists("/root/pppd")) // invalid soft link
	t.Log("not exists dir: ", Exists("/isnotexists"))
}

func TestIsDir(t *testing.T) {
	t.Log("console", IsDir("/dev/console"))
	t.Log("proc", IsDir("/proc"))
	t.Log("device: ", IsDir("/dev/vda1"))
	t.Log("soft link: ", IsDir("/etc/localtime"))     // soft link
	t.Log("invalid soft link: ", IsDir("/root/pppd")) // invalid soft link
	t.Log("etc: ", IsDir("/etc"))
	t.Log("not exists dir: ", IsDir("/isnotexists"))
}

func TestIsFile(t *testing.T) {
	t.Log("console", IsFile("/dev/console"))
	t.Log("proc", IsFile("/proc"))
	t.Log("device: ", IsFile("/dev/vda1"))
	t.Log("soft link: ", IsFile("/etc/localtime"))     // soft link
	t.Log("invalid soft link: ", IsFile("/root/pppd")) // invalid soft link
	t.Log("etc: ", IsFile("/etc"))
	t.Log("not exists dir: ", IsFile("/isnotexists"))
}
