package jsonxx

import (
	"fmt"
	"io"
	"os"

	jsoniter "github.com/json-iterator/go" //解析json 比官方的"encoding/json"更快
)

type Fofa_json []struct { //fofa.json对应的结构体 用https://mholt.github.io/json-to-go/一键生成
	RuleID         string `json:"rule_id"`
	Level          string `json:"level"`
	Softhard       string `json:"softhard"`
	Product        string `json:"product"`
	Company        string `json:"company"`
	Category       string `json:"category"`
	ParentCategory string `json:"parent_category"`
	Rules          [][]struct {
		Match   string `json:"match"`
		Content string `json:"content"`
	} `json:"rules"`
}

var Fofa Fofa_json                                      //定义结构体用于解析json
var json = jsoniter.ConfigCompatibleWithStandardLibrary //实例化工具类

func Read() Fofa_json { //Read方法: 读取指纹json文件并解析,返回解析后的json数据
	jsonfile, err1 := os.Open(`./fofa.json`) //打开json文件并进行错误处理
	if err1 != nil {
		fmt.Println(err1)
	} else {
		fmt.Println("Successfully Opened fofa.json")
	}

	defer jsonfile.Close()

	byteValue, _ := io.ReadAll(jsonfile)     //读取json数据
	err2 := json.Unmarshal(byteValue, &Fofa) //解析json数据并错误处理
	if err2 != nil {
		fmt.Println(err2.Error())
	}

	return Fofa
}
