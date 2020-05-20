package main

import (
	"flag"
	"fmt"
	"github.com/fjlyx97/golang-stress-tool/stress_core"
	"runtime"
)

var (
	configPath string
	requestNum int64
	concurrencyNum int
)

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	flag.StringVar(&configPath,"f","","Select json config path")
	flag.Int64Var(&requestNum,"n",100,"Set number of requests")
	flag.IntVar(&concurrencyNum,"c",5,"Set the number of concurrent connections")
}

func main() {
	flag.Parse()
	if len(configPath) == 0 {
		fmt.Printf("You must pecify a json file path.\n")
		fmt.Printf("eg. go run main.go -n 1000 -c 100 -f your_json_file.json\n")
		return
	}

	var req = stress_core.CreateHttpRequest(configPath)
	c := stress_core.Client{}
	c.Run(requestNum,concurrencyNum,req)
	c.GetResult()
}