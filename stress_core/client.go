package stress_core

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

var countDownLatch = sync.WaitGroup{}
var waitGoroutine = sync.WaitGroup{}
var count int64


type Client struct {
	requestNum int64
	concurrency int
	requestType interface{}
	resultChan chan OutputResult
	runTime time.Duration
}

func (c *Client)Run(requestNum int64 , concurrency int , requestType interface{}) {
	c.requestNum = requestNum
	c.concurrency = concurrency
	c.requestType = requestType
	c.resultChan = make(chan OutputResult,requestNum)

	//使所有线程同步开始
	waitGoroutine.Add(concurrency)
	countDownLatch.Add(concurrency)
	count = requestNum

	startTime := time.Now()
	for i := 0 ; i < concurrency ; i++ {
		countDownLatch.Done()
		go c.createGoroutine()
	}
	waitGoroutine.Wait()
	c.runTime = time.Now().Sub(startTime)
}

func (c *Client)createGoroutine() {

	for {
		ok := atomic.CompareAndSwapInt64(&count,count,count-1)
		if !ok {
			continue
		}
		if atomic.LoadInt64(&count) < 0 {
			break
		}

		switch c.requestType.(type) {
		case HttpRequest:
			countDownLatch.Wait()
			sendHttpRequest(c.requestType.(HttpRequest),c.resultChan)
		}
	}
	waitGoroutine.Done()
}

func (c *Client)GetResult() {
	switch c.requestType.(type) {
	case HttpRequest:
		result := ProcessOutput(c.resultChan)
		fmt.Printf("Used for : %f second \n",c.runTime.Seconds())
		fmt.Printf("Request url : %s\n",c.requestType.(HttpRequest).url)
		fmt.Printf("Concurrency : %d \n",c.concurrency)
		fmt.Printf("Requests : %d \n",c.requestNum)
		fmt.Printf("Fastest request : %dms \n",result["fastestReq"])
		fmt.Printf("Slowest request : %dms \n",result["slowestReq"])
		fmt.Printf("Avg request : %fms \n",result["avgReq"])
		fmt.Printf("Query per second : %f \n",float64(result["totalOutputs"].(int))/c.runTime.Seconds())
		fmt.Printf("Success request : %d \n",result["successReq"])
		fmt.Printf("Error request : %d \n",result["errorReq"])
		fmt.Printf("Total results : %d \n",result["totalOutputs"])
	}
}
