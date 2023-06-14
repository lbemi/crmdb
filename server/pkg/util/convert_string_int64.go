package util

import (
	"github.com/lbemi/lbemi/pkg/restfulx"
	"strconv"
	"strings"
)

func ParseStrInt64(value string) []uint64 {
	vs := make([]uint64, 0)
	for _, v := range strings.Split(value, ",") {
		i, err := strconv.ParseUint(v, 10, 64)
		if err != nil {
			restfulx.ErrNotNilDebug(err, restfulx.ParamErr)
		}
		vs = append(vs, i)
	}
	return vs
}
