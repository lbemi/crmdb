package util

import (
	"github.com/lbemi/lbemi/pkg/restfulx"
	"strconv"
	"strings"
)

func ConvertSliceStrToInt(str string) []uint64 {
	if str == "" {
		return []uint64{}
	}
	strSlice := strings.Split(str, ",") // 按逗号分割字符串得到切片
	var intSlice []uint64

	for _, s := range strSlice {
		i, err := strconv.ParseUint(s, 10, 64)
		if err != nil {
			// 处理转换错误
			restfulx.ErrNotNilDebug(err, restfulx.ParamErr)
			continue
		}
		intSlice = append(intSlice, i)
	}
	return intSlice
}
