package util

import (
	"errors"
	_ "net/http/pprof"
	"strconv"
	"strings"
	"time"
)

type FileItem struct {
	Name        string `json:"name"`
	IsDir       bool   `json:"isDir"`
	FsType      string `json:"fsType"`
	Size        int64  `json:"size"`
	LastModify  string `json:"lastModify"`
	User        string `json:"user"`
	Group       string `json:"group"`
	Permissions string `json:"permissions"`
}

func ParseFsItem(fileString string) (*FileItem, error) {
	fsItem := &FileItem{}

	// 自定义的分割函数，处理不同类型的空白字符
	//splitFunc := func(r rune) bool {
	//	return unicode.IsSpace(r) || r == '\t'
	//}

	// 分割字符串
	parts := strings.Fields(fileString)

	if len(parts) < 9 {
		return fsItem, errors.New("invalid file string")
	}

	// 解析权限模式
	fsItem.IsDir = parts[0][0] == 'd'

	// 解析文件类型
	// if fsItem.IsDir {
	// 	fsItem.FsType = "d"
	// } else {
	// 	fsItem.FsType = "-"
	// }
	fsItem.IsDir = parts[0][0] == 'd'
	fsItem.FsType = string(parts[0][0])
	fsItem.Permissions = parts[0][1:]
	// 解析文件大小
	size, err := strconv.ParseInt(parts[4], 10, 64)
	if err != nil {
		return fsItem, err
	}
	fsItem.Size = size

	// 解析最后修改时间
	// modifyTimeStr := strings.Join(parts[5:8], " ") + " " + strconv.Itoa(time.Now().Year())
	// modifyTime, err := time.Parse("Jan 2 15:04 2006", modifyTimeStr)
	// if err != nil {
	// 	modifyTime, err = time.Parse("Jan 2 2006", modifyTimeStr)
	// 	if err != nil {
	// 		return fsItem, err
	// 	}
	// }
	// fsItem.LastModify = modifyTime.Format("2006-01-02 15:04")

	modifyTimeStr := parts[5] + " " + parts[6] + " " + parts[7]

	if strings.Contains(modifyTimeStr, ":") {
		year := time.Now().Year()
		modifyTimeStr = modifyTimeStr + " " + strconv.Itoa(year)
		modifyTime, err := time.Parse("Jan 2 15:04 2006", modifyTimeStr)
		if err != nil {
			return fsItem, err
		}
		fsItem.LastModify = modifyTime.Format("2006-01-02 15:04")
	} else {
		modifyTime, err := time.Parse("Jan 2 2006", modifyTimeStr)
		if err != nil {
			return fsItem, err
		}
		fsItem.LastModify = modifyTime.Format("2006-01-02 15:04")
	}

	// 解析文件名
	fsItem.Name = parts[8]

	// 解析所有者和所属组
	fsItem.User = parts[2]
	fsItem.Group = parts[3]

	return fsItem, nil
}

func GetDirAndFiles(fileString string) []*FileItem {
	fileStrList := strings.Split(fileString, "\n")

	ret := make([]*FileItem, 0, len(fileStrList))
	for _, item := range fileStrList {
		if fsItem, err := ParseFsItem(item); err == nil {
			ret = append(ret, fsItem)
		}
	}
	return ret
}
