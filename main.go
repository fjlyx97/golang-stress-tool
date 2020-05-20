package main

import (
	"golang-stress/stress_core"
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	var req = stress_core.CreateHttpRequest("./config/config.json")
	//stress_core.SendHttpRequest(req)
	c := stress_core.Client{}
	c.Run(100000,10000,req)
	c.GetResult()
}