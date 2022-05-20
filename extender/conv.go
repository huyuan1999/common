// @Time    : 2022/5/19 8:34 下午
// @Author  : HuYuan
// @File    : conv.go
// @Email   : huyuan@virtaitech.com

package extender

import (
	"fmt"
	"strconv"
	"strings"
)

func StorageUnitConvToBytes(size string) (int64, error) {
	SIZE := strings.ToUpper(size)
	switch {
	case strings.HasSuffix(SIZE, "GB"), strings.HasSuffix(SIZE, "G"):
		s, err := strconv.ParseFloat(strings.Trim(strings.Trim(SIZE, "GB"), "G"), 64)
		if err != nil {
			return 0, err
		}
		v := s * 1024 * 1024 * 1024
		return int64(v), nil
	case strings.HasSuffix(SIZE, "MB"), strings.HasSuffix(SIZE, "M"):
		s, err := strconv.ParseFloat(strings.Trim(strings.Trim(SIZE, "MB"), "M"), 64)
		if err != nil {
			return 0, err
		}
		v := s * 1024 * 1024
		return int64(v), nil
	case strings.HasSuffix(SIZE, "KB"), strings.HasSuffix(SIZE, "K"):
		s, err := strconv.ParseFloat(strings.Trim(strings.Trim(SIZE, "KB"), "K"), 64)
		if err != nil {
			return 0, err
		}
		v := s * 1024
		return int64(v), nil
	}
	return 0, fmt.Errorf("unsupported unit")
}

func StorageUnitConvToGB(size string) (int64, error) {
	SIZE := strings.ToUpper(size)
	switch {
	case strings.HasSuffix(SIZE, "MB"), strings.HasSuffix(SIZE, "M"):
		s, err := strconv.ParseFloat(strings.Trim(strings.Trim(SIZE, "MB"), "M"), 64)
		if err != nil {
			return 0, err
		}
		v := s / 1024
		return int64(v), nil
	case strings.HasSuffix(SIZE, "KB"), strings.HasSuffix(SIZE, "K"):
		s, err := strconv.ParseFloat(strings.Trim(strings.Trim(SIZE, "KB"), "K"), 64)
		if err != nil {
			return 0, err
		}
		v := s / 1024 / 1024
		return int64(v), nil
	case strings.HasSuffix(SIZE, "B"):
		s, err := strconv.ParseFloat(strings.Trim(SIZE, "B"), 64)
		if err != nil {
			return 0, err
		}
		v := s / 1024 / 1024 / 1024
		return int64(v), nil
	}
	return 0, fmt.Errorf("unsupported unit")
}
