package util

import (
	"errors"
	"github.com/gin-gonic/gin"
	"strconv"
)

// GetQueryToStrE
func GetQueryToStrE(c *gin.Context, key string) (string, error) {
	str := c.Param(key)
	if str == "" {
		return "", errors.New("没有这个值传入")
	}

	return str, nil
}

// GetQueryToStr
func GetQueryToStr(c *gin.Context, key string, defaultValues ...string) string {
	var defaultValue string
	if len(defaultValues) > 0 {
		defaultValue = defaultValues[0]
	}
	str, err := GetQueryToStrE(c, key)
	if str == "" || err != nil {
		return defaultValue
	}
	return str
}

// GetQueryToUintE
func GetQueryToUintE(c *gin.Context, key string) (uint64, error) {
	str, err := GetQueryToStrE(c, key)
	if err != nil {
		return 0, err
	}
	intNum, _ := strconv.ParseUint(str, 10, 32)

	return intNum, nil
}

// QueryToUint
func GetQueryToUint(c *gin.Context, key string, defaultValues ...uint64) uint64 {
	var defaultValue uint64
	if len(defaultValues) > 0 {
		defaultValue = defaultValues[0]
	}
	val, err := GetQueryToUintE(c, key)
	if err != nil {
		return defaultValue
	}
	return val
}

// QueryToUintE
func getQueryToUint64E(c *gin.Context, key string) (uint64, error) {
	str, err := GetQueryToStrE(c, key)
	if err != nil {
		return 0, err
	}
	return strconv.ParseUint(str, 10, 64)
}

// GetQueryToUint64 获取参数并转换成uint64
func GetQueryToUint64(c *gin.Context, key string, defaultValues ...uint64) uint64 {
	var defaultValue uint64
	if len(defaultValues) > 0 {
		defaultValue = defaultValues[0]
	}
	val, err := getQueryToUint64E(c, key)
	if err != nil {
		return defaultValue
	}
	return val
}

// GetQueryToInt64 获取参数并转换成uint64
func GetQueryToInt64(c *gin.Context, key string, defaultValues ...uint64) uint64 {
	var defaultValue uint64
	if len(defaultValues) > 0 {
		defaultValue = defaultValues[0]
	}
	val, err := getQueryToUint64E(c, key)
	if err != nil {
		return defaultValue
	}
	return val
}
