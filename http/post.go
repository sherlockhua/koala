package http

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func PostJson(url string, data interface{}) (res []byte, err error) {

	fmt.Println(url)
	jsonData, err := json.Marshal(data)
	if err != nil {
		fmt.Println("Error marshaling JSON:", err)
		return
	}

	header := make(map[string]string)
	header["Content-Type"] = "application/json"
	res, err = Post(url, jsonData, header)
	return
}

func Post(url string, data []byte, header map[string]string) (res []byte, err error) {

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(data))
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}

	// 设置请求头，表明发送的数据是JSON格式
	//req.Header.Set("Content-Type", "application/json")
	for k, v := range header {
		req.Header.Set(k, v)
	}

	// 使用默认的HTTP客户端发送请求
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error making POST request:", err)
		return
	}
	defer resp.Body.Close()

	// 读取响应体
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return
	}

	return body, nil
}
