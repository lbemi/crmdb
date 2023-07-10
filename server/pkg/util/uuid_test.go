package util

import (
	"fmt"
	"testing"
)

func TestGetUUID(t *testing.T) {
	id := GetUUID()
	fmt.Println(id)
}
