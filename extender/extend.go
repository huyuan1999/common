// @Time    : 2022/5/19 8:33 下午
// @Author  : HuYuan
// @File    : extend.go
// @Email   : huyuan@virtaitech.com

package extender

import (
	"io/ioutil"
	"path"
)

func Tree(dir string) []string {
	var fileList []string
	fs, err := ioutil.ReadDir(dir)
	if err != nil {
		return nil
	}
	for _, f := range fs {
		if f.IsDir() {
			fileList = append(fileList, Tree(path.Join(dir, f.Name()))...)
			continue
		}
		fileList = append(fileList, path.Join(dir, f.Name()))
	}
	return fileList
}
