// @Time    : 2022/5/19 9:09 下午
// @Author  : HuYuan
// @File    : rand_test.go
// @Email   : huyuan@virtaitech.com

package rand

import "testing"

func TestRandom(t *testing.T) {
	for i := 0; i < 100; i++ {
		t.Log(Random(10))
	}
}

func TestGeneratorMacAddr(t *testing.T) {
	for i := 0; i < 100; i++ {
		t.Log(GeneratorMacAddr())
	}
}
