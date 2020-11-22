/*
@Time : 2020/11/22 18:05
@Author : Firbath
@File : demo.go
@Software: GoLand
*/
package main

import (
	"encoding/json"
	"fmt"
	"github.com/mitchellh/mapstructure"
	"io/ioutil"
)

type Cls struct {
	Action string      `json:"action"`
	Data   interface{} `json:"data"`
}

type Sub struct {
	Ab []string
	Cd string
}

func GoDemo() {
	data, err := ioutil.ReadFile("demo.json")
	if err != nil {
		fmt.Println("read fail", err)
	}
	fmt.Println(data)
	jsonData := make(map[string]interface{})
	clsData := &Cls{}
	//subData := &Sub{}
	if err := json.Unmarshal(data, clsData); err != nil {
		fmt.Println(err)
	}
	if err := json.Unmarshal(data, &jsonData); err != nil {
		fmt.Println(err)
	}
	fmt.Println(*clsData)
	fmt.Println(jsonData)
	fmt.Println(jsonData["action1"])
	datas := jsonData["data"].(map[string]interface{})
	fmt.Println(datas)
	fmt.Println(datas["Cd"])
	var sub Sub
	//将 map 转换为指定的结构体
	if err := mapstructure.Decode(datas, &sub); err != nil {
		fmt.Println(err)
	}
	fmt.Println(sub)

}

func main() {
	GoDemo()
}
