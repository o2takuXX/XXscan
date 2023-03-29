package main

import (
	"XXscan/jsonxx"
	"XXscan/scanxx"
	"XXscan/urllist"
	"flag"
	"fmt"
	"os"
)

func main() {

	var url = flag.String("url", "", "input your url") //定义传参url

	o2takuXX() //输出相关信息

	if !flag.Parsed() { //解析传参
		flag.Parse()
	}

	fmt.Println("----------------------------------------------")

	finger_data := jsonxx.Read() //解析fofa.json

	if *url != "" { //单个url
		scanxx.Scanurl(*url, finger_data)

	}
	if *urllist.File != "" { //url.txt文本
		list := urllist.List()
		for _, url := range list {
			scanxx.Scanurl(url, finger_data)
		}
	}
	if *url == "" && *urllist.File == "" {
		fmt.Println("please input your url or file")
		os.Exit(0)
	}
	fmt.Println("----------------------------------------------")
	fmt.Println("Done!")

}

func o2takuXX() {
	o2takuXX := `
___________________________________________
|             ___   ___      _      _     |
| \  / \  /  (   ) (   )    / |    /|   / |
|  \/   \/    \    |       /__|   / |  /  |
|  /\   /\     \   |      /   |  /  | /   |
| /  \ /  \ (___)  (___) /    | /   |/    |
|_________________________________________|
|              by o2takuXX                |
|______________version 1.0________________|
	`
	fmt.Println(o2takuXX)
}
