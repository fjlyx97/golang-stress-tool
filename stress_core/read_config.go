package stress_core

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"time"
)

type HttpRequest struct {
	method string
	url string
	contentType string
	timeout int
	keepAlive bool
	postData interface{}
}

type OutputResult struct {
	statusCode int
	duration time.Duration
}

//解析json配置文件
func readJsonConfig(path string) map[string]interface{}{
	data , err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Println("Read json config error")
		os.Exit(-1)
	}
	var allData map[string]interface{}
	err = json.Unmarshal(data,&allData)
	if err != nil {
		fmt.Println("Unmarshal json data error")
		os.Exit(-1)
	}
	return allData
}

//通过解析后的文件构造Http请求
func CreateHttpRequest(path string) HttpRequest {
	allData := readJsonConfig(path)
	//构建返回的请求（可以使用反射动态创建）
	var req HttpRequest
	if allData["url"] != nil {
		req.url = allData["url"].(string)
	}

	if allData["content_type"] != nil {
		req.contentType = allData["content_type"].(string)
	}

	if allData["method"] != nil {
		req.method = allData["method"].(string)
	}

	if allData["post_data"] != nil {
		req.postData = allData["post_data"]
	}

	if allData["keep_alive"] != nil {
		req.keepAlive = allData["keep_alive"].(bool)
	} else {
		req.keepAlive = false
	}

	if allData["timeout"] != nil {
		req.timeout = int(allData["timeout"].(float64))
	} else {
		req.timeout = 30
	}


	return req
}
