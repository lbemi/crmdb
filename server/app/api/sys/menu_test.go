package sys

import (
	"github.com/gin-gonic/gin"
	"testing"
)

func TestGetMenuList(t *testing.T) {
	type args struct {
		c *gin.Context
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			GetMenuList(tt.args.c)
		})
	}
}
