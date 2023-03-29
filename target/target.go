package target

import (
	"fmt"
	"io"
	"net/http"
	"regexp"
	"strings"
)

func Requset(url string) *http.Response { //发送请求获取响应

	response, err1 := http.Get(url) //get请求
	if err1 != nil {
		fmt.Println(err1.Error())
	}

	return response
}

func Have_Body(data *http.Response) []byte { //获取响应正文

	response := data
	Response_data, err := io.ReadAll(response.Body) //获取解析响应数据
	if err != nil {
		fmt.Println(err.Error())
	}

	return Response_data
}

func Have_Server(data *http.Response) string { //获取header中的server

	response := data
	server := response.Header["Server"][0]

	return server
}

func Have_Title(data []byte) string { //在响应正文中用正则匹配获取title

	response := data
	re := regexp.MustCompile("<title.*>(.*?)</title>")
	title := re.FindStringSubmatch(string(response))[1]

	return title
}

func Have_Banner(data []byte) [][]string { //在正文中用正则匹配获取banner

	response := data
	re := regexp.MustCompile(`(?im)<\s*banner.*>(.*?)<\s*/\s*banner>`)
	banner := re.FindAllStringSubmatch(string(response), -1) //这里抄的Ch1nfo师傅,因为banner不知道怎么测试

	return banner
}

func Have_Header(response *http.Response) http.Header { //获取响应中的header内容
	data := response.Header

	return data
}

func Find_Body(data []byte, rule string) bool { //匹配响应正文中对应的特征

	find := strings.Contains(string(data), rule)

	return find
}

func Find_Server(data string, rule string) bool { //匹配Server中对应的特征

	if data == rule {
		return true
	} else {
		return false
	}

}

func Find_Title(data string, rule string) bool { //匹配Title中对应的特征

	if data == rule {
		return true
	} else {
		return false
	}

}

func Find_Banner(data [][]string, rule string) bool { //匹配Banner中对应的特征,抄的ch1nfo师傅

	flag := false
	for i := 0; i < len(data); i++ {
		if rule == data[i][1] {
			flag = true
		}
	}
	return flag
}

func Find_Header(data http.Header, rule string) bool { //匹配Header中是否存在对应特征
	x := ""
	for key, value := range data {
		x = x + " " + key
		for i := 0; i < len(value); i++ {
			x = x + " " + value[i]
		}
	}
	find := strings.Contains(x, rule)

	return find
}

func Target(url string, match string, rule string, response *http.Response, body []byte) bool { //匹配对应特征
	flag := false
	if match == "body_contains" {
		if Find_Body(body, rule) {
			flag = true
			return flag
		}
	} else if match == "server_contains" || match == "server" {
		if Find_Server(Have_Server(response), rule) {
			flag = true
			return flag
		}
	} else if match == "title_contains" {
		if Find_Title(Have_Title(body), rule) {
			flag = true
			return flag
		}
	} else if match == "banner_contains" {
		if Find_Banner(Have_Banner(body), rule) {
			flag = true
			return flag
		}
	} else if match == "header_contains" {
		if Find_Header(Have_Header(response), rule) {
			flag = true
			return flag
		}
	}

	return flag
}
