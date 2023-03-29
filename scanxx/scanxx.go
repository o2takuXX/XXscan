package scanxx

import (
	"XXscan/jsonxx"
	"XXscan/target"
	"fmt"
)

func Scanurl(url string, finger_data jsonxx.Fofa_json) {

	fmt.Println("####   " + url + "   ####")
	response := target.Requset(url) //http请求获取相应
	body := target.Have_Body(response)
	for i := 0; i < len(finger_data); i++ {
		Product := finger_data[i].Product                //对应的product
		for j := 0; j < len(finger_data[i].Rules); j++ { //循环遍历fofa.json
			x := 0 //用于判断多个特征点
			for k := 0; k < len(finger_data[i].Rules[j]); k++ {
				if target.Target(url, finger_data[i].Rules[j][k].Match, finger_data[i].Rules[j][k].Content, response, body) {
					x++ //匹配成功一个特征就加一
				}
			}
			if len(finger_data[i].Rules[j]) == x { //特征判断
				fmt.Print("-->  ")
				fmt.Println(Product)
			}
		}
	}
}
