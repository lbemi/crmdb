package util

import (
	"github.com/lbemi/lbemi/pkg/global"
	"strconv"
)

// ParseInt64 将字符串转换为 int64
func ParseInt64(s string) uint64 {
	res, err := strconv.ParseUint(s, 10, 64)
	if err != nil {
		global.Logger.Error(err)
	}
	return res
}
