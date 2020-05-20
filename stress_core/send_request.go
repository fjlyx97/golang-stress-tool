package stress_core

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"time"
)

func sendHttpRequest(request HttpRequest,resChan chan OutputResult) {
	var result OutputResult
	postData , err := json.Marshal(request.postData)
	if err != nil {
		panic("Convert postData failed")
	}
	body := bytes.NewReader(postData)
	req , err := http.NewRequest(request.method,request.url,body)
	defer req.Body.Close()
	if err != nil {
		fmt.Println(err.Error())
		//result.statusCode = -1
		//resChan <- result
		return
	}
	//设置请求头
	if len(request.contentType) != 0 {
		req.Header.Set("Content-Type",request.contentType)
	}

	tran := &http.Transport{
		DisableKeepAlives: !request.keepAlive,
		ResponseHeaderTimeout: time.Duration(request.timeout) * time.Second,
	}

	cli := http.Client{
		Transport: tran,
	}
	startTime := time.Now()
	response , err := cli.Do(req)
	if err != nil {
		//fmt.Println(err.Error())
		result.statusCode = -1
		resChan <- result
		return
	}
	_ , _ = io.Copy(ioutil.Discard,response.Body)
	result.duration = time.Now().Sub(startTime)
	result.statusCode = response.StatusCode
	resChan <- result
}
