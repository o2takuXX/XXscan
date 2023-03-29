package urllist

import (
	"bufio"
	"flag"
	"fmt"
	"os"
)

var File = flag.String("file", "", "input path to your url.txt") //定义传参file

func List() (lines []string) {

	file, err := os.Open(*File) //打开文件
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file) //读文件并返回数据
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines
}
