package stress_core

import (
	"math"
)

func ProcessOutput(outputs chan OutputResult) map[string]interface{}{
	var result = make(map[string]interface{})
	var totalOutputs = len(outputs)
	var fastestReq int64 = math.MaxInt64
	var slowestReq int64 = math.MinInt64
	var avgReq float64 = 0
	var errReq int = 0
	var successReq int = 0

	for i := 0 ; i < totalOutputs ; i++ {
		output := <- outputs
		if output.statusCode != 200 {
			errReq++
			continue
		}
		successReq++
		o := output.duration.Milliseconds()
		if o < fastestReq {
			fastestReq = o
		}
		if o > slowestReq {
			slowestReq = o
		}
		avgReq += float64(o)
	}
	if fastestReq != math.MaxInt64 {
		result["fastestReq"] = fastestReq
	} else {
		result["fastestReq"] = 0
	}
	if slowestReq != math.MinInt64 {
		result["slowestReq"] = slowestReq
	} else {
		result["slowestReq"] = 0
	}

	if successReq != 0 {
		result["avgReq"] = avgReq / float64(successReq)
	} else {
		result["avgReq"] = 0.0
	}

	result["totalOutputs"] = totalOutputs
	result["successReq"] = successReq
	result["errorReq"] = errReq
	return result
}
