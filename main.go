package main

import (
	"github.com/montanaflynn/stats"
	"log"
	"time"
)

const ExternalIterations = 5
const InternalIterations = 1000
const NumberOfPersons = 100

func main() {
	testProtocol("Json", func() (time.Duration, time.Duration, float64) {
		jsonInit()
		jsonEncodeTime, jsonSize := measure(testJsonEncode)
		jsonDecodeTime, _  		 := measure(testJsonDecode)

		return jsonEncodeTime, jsonDecodeTime, jsonSize
	})

	testProtocol("Protobuf", func() (time.Duration, time.Duration, float64) {
		protoInit()
		jsonEncodeTime, jsonSize := measure(testProtoEncode)
		jsonDecodeTime, _  		 := measure(testProtoDecode)

		return jsonEncodeTime, jsonDecodeTime, jsonSize
	})

	testProtocol("Thrift Compact", func() (time.Duration, time.Duration, float64) {
		thriftInit()
		jsonEncodeTime, jsonSize := measure(testThriftEncode)
		jsonDecodeTime, _  		 := measure(testThriftDecode)

		return jsonEncodeTime, jsonDecodeTime, jsonSize
	})
}

func measure(cb func() float64) (time.Duration, float64) {
	now := time.Now()
	median := cb()
	return time.Since(now), median
}

func testProtocol(name string, cb func() (time.Duration, time.Duration, float64))  {
	var encodeStats []float64
	var decodeStats []float64
	var sizeStats []float64
	for i := 0; i < ExternalIterations; i++ {
		encodeTime, decodeTime, size := cb()
		//log.Println(encodeTime, decodeTime, size)
		encodeStats = append(encodeStats, float64(encodeTime))
		decodeStats = append(decodeStats, float64(decodeTime))
		sizeStats = append(sizeStats, size)
	}

	log.Println(name, ": ")
	PrintMetrics(calcMedians(encodeStats, decodeStats, sizeStats))
}

func calcMedians(encodeStats, decodeStats, sizeStats []float64) (encodeMedian, decodeMedian, sizeMedian float64) {
	var err error

	encodeMedian, err = stats.Median(encodeStats)
	if err != nil {
		return
	}

	decodeMedian, err = stats.Median(decodeStats)
	if err != nil {
		return
	}

	sizeMedian, err = stats.Median(sizeStats)
	if err != nil {
		return
	}

	return
}

func PrintMetrics(encodeMedian, decodeMedian, sizeMedian float64)  {
	log.Println("Encode median time: ", time.Duration(encodeMedian))
	log.Println("Decode median time: ", time.Duration(decodeMedian))
	log.Println("Size median bytes: ", sizeMedian)
}