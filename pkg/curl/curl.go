package curl

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func POST(url string, data interface{}, header map[string]string) []byte {
	bs, err := json.Marshal(data)

	reader := bytes.NewReader(bs)

	request, err := http.NewRequest("POST", url, reader)
	if err != nil{
		fmt.Println("请求server端失败...")
	}

	// 携带头部
	request.Header.Set("Content-Type", "application/json;charset=UTF-8")

	// fmt.Printf("mapNum:%d \n", len(header))

	if len(header) > 0 {
		for key,value := range header {
			request.Header.Set(key, value)
		}
	}

	client := http.Client{}
	// 返回服务端的响应数据
	resp, _ := client.Do(request)
	if err != nil {
		fmt.Println("请求获取响应失败")
	}
	body, _ := ioutil.ReadAll(resp.Body)

	return body
}
