package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

type ipInfo struct {
	Ret  string //正常返回为 ok
	IP   string //查询的ip地址
	Data []string
	/*
		"中国", //返回ip所在国家
		"广西省", //返回ip所在省或者直辖市 ，如果数据不存在返回 空 或者 null
		"南宁市", // 返回ip所在市区或者直辖市，如果数据不存在返回 空 或者 null
		"横县",
		// 返回ip所在县 ，如果数据不存在返回 空 或者 null
		"百合镇",
		//返回ip所在乡、镇 ，如果数据不存在返回 空 或者 null
		"xx小区", // 返回ip所在小区 ，如果数据不存在返回 空 或者 null
		"成记网吧", //返回ip所在公司，比如xx网吧，xx公司（百度，搜狗，搜狐）等等 ，如果数据不存在返回 空 或者 null
		"电信" // 网络运营商 ，如果数据不存在返回 空 或者 null
	*/
}

func GetRealAddressByIP(ip string) string {
	if ip == "127.0.0.1" {
		return "内网IP"
	}

	resp, err := http.Get(fmt.Sprintf("https://www.inte.net/tool/ip/api.ashx?ip=%s&datatype=json&key=12", ip))
	if err != nil {
		return "unknown"
	}
	defer resp.Body.Close()

	ipInfo := &ipInfo{}
	err = json.NewDecoder(resp.Body).Decode(ipInfo)
	if err != nil {
		return "unknown"
	}

	location := strings.Join(ipInfo.Data, "")
	return location
}
func main() {
	// https://www.inte.net/tool/ip/api.ashx?ip=183.136.213.98&datatype=json&key=12
	// http://ip-api.com/json/123.123.123.123?lang=zh-CN
	ip := "127.0.0.1"
	res := GetRealAddressByIP(ip)
	fmt.Println(res)

}
