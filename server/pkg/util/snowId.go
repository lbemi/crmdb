package util

import (
	"github.com/bwmarrin/snowflake"
	"strconv"
)

func GetSnowID() uint64 {
	node, _ := snowflake.NewNode(1)
	res, _ := strconv.ParseUint(node.Generate().String(), 10, 64)
	return res
}
