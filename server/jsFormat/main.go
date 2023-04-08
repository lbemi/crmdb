package main

import (
	"encoding/json"
	"fmt"
	"github.com/lbemi/lbemi/jsFormat/model"
	"os"
)

func main() {
	bytes, err := os.ReadFile("./test.json")
	if err != nil {
		panic(err)
	}
	var oldJson []model.OldJs
	err = json.Unmarshal(bytes, &oldJson)
	if err != nil {
		panic(err)
	}
	//fmt.Println(oldJson)
	//fmt.Println("----------------")
	var newJsons []model.NewJS
	for _, v := range oldJson {
		var newJS model.NewJS
		newJS.Id = v.Field
		newJS.Name = v.Title
		newJS.Align = v.Align
		if v.Formatter == "moneyFormate" {
			newJS.Type = "money"
		} else if v.Formatter == "timeStampFormat" {
			newJS.Type = "time"
		} else if v.Formatter == "" {
			newJS.Type = ""
		} else {
			newJS.Type = "customs"
		}
		newJS.Custom = make([]string, 0)
		newJS.Width = "50"
		newJsons = append(newJsons, newJS)
	}
	//fmt.Println(newJsons)

	//fmt.Println("----------------")
	data, err := json.Marshal(&newJsons)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(data))
}
