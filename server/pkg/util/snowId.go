package util

import (
	"github.com/bwmarrin/snowflake"
)

func GetSnowID() int64 {
	node, _ := snowflake.NewNode(1)
	return node.Generate().Int64()
}
