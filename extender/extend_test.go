// @Time    : 2022/5/19 8:47 下午
// @Author  : HuYuan
// @File    : conv_test.go
// @Email   : huyuan@virtaitech.com

package extender

import "testing"

func TestStorageUnitConvToBytes(t *testing.T) {
	t.Log(StorageUnitConvToBytes("10g"))
	t.Log(StorageUnitConvToBytes("10Gb"))
	t.Log(StorageUnitConvToBytes("10GB"))
	t.Log(StorageUnitConvToBytes("10gb"))
	t.Log(StorageUnitConvToBytes("10.12gb"))
	t.Log(StorageUnitConvToBytes("10m"))
	t.Log(StorageUnitConvToBytes("10Mb"))
	t.Log(StorageUnitConvToBytes("10MB"))
	t.Log(StorageUnitConvToBytes("10.88MB"))
	t.Log(StorageUnitConvToBytes("10mb"))
	t.Log(StorageUnitConvToBytes("10k"))
	t.Log(StorageUnitConvToBytes("10Kb"))
	t.Log(StorageUnitConvToBytes("10KB"))
	t.Log(StorageUnitConvToBytes("10.46KB"))
	t.Log(StorageUnitConvToBytes("10kb"))
	t.Log(StorageUnitConvToBytes("10ki"))
	t.Log(StorageUnitConvToBytes("10gi"))
}

func TestStorageUnitConvToGB(t *testing.T) {
	t.Log(StorageUnitConvToGB("10737418240b"))
	t.Log(StorageUnitConvToGB("10737418240B"))
	t.Log(StorageUnitConvToGB("10737418240.04B"))
	t.Log(StorageUnitConvToGB("10737418240.48B"))
	t.Log(StorageUnitConvToGB("1024m"))
	t.Log(StorageUnitConvToGB("1024mb"))
	t.Log(StorageUnitConvToGB("1024MB"))
	t.Log(StorageUnitConvToGB("1024Mb"))
	t.Log(StorageUnitConvToGB("1024mB"))
	t.Log(StorageUnitConvToGB("1024mB"))
	t.Log(StorageUnitConvToGB("1024.9318301mB"))
	t.Log(StorageUnitConvToGB("1048576k"))
	t.Log(StorageUnitConvToGB("1048576kb"))
	t.Log(StorageUnitConvToGB("1048576kB"))
	t.Log(StorageUnitConvToGB("1048576Kb"))
	t.Log(StorageUnitConvToGB("1048576KB"))
	t.Log(StorageUnitConvToGB("1048576.0178319KB"))
}

func TestTree(t *testing.T) {
	//t.Log(Tree("/etc/"))
	t.Log(Tree("/etc/network"))
	t.Log(Tree("/var/log/"))
	t.Log(Tree("/isnotexists"))
}
