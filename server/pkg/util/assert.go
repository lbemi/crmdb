package util

import "fmt"

// IsTrue  判断条件是否为真，否则直接panic
func IsTrue(condition bool, panicMsg string, params ...interface{}) {
	if !condition {
		if len(params) != 0 {
			panic(fmt.Sprintf(panicMsg, params...))
		}
		panic(panicMsg)
	}
}
