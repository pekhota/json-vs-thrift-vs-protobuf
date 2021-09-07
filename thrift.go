package main

import (
	"bytes"
	"context"
	"github.com/apache/thrift/lib/go/thrift"
	"github.com/bxcodec/faker/v3"
	"github.com/montanaflynn/stats"
	"github.com/pekhota/json-vs-thrift-vs-protobuf/internal/pkg/thriftexample"
	"github.com/pekhota/json-vs-thrift-vs-protobuf/internal/pkg/thrifttimestamp"
	"log"
	"time"
)

var thriftData [][]byte

func init() {
	thriftInit()
}

func thriftInit() {
	thriftData = make([][]byte, 0)
}

func testThriftEncode() float64 {
	data := make([]float64, 0)

	makePersons := func(n int) []*thriftexample.Person {
		l := make([]*thriftexample.Person, 0)

		for i := 0; i < n; i++ {
			t := time.Now()
			ts := thrifttimestamp.Timestamp{
				Seconds: t.Unix(),
				Nanos:   int32(t.Nanosecond()),
			}
			p := thriftexample.Person{
				ID:    int32(i),
				Name:  faker.Name(),
				Email: faker.Email(),
				Phones: []*thriftexample.PhoneNumber{
					{
						Number: faker.Phonenumber(),
						Type:   thriftexample.PhoneType_HOME,
					},
				},
				LastUpdated: &ts,
			}
			l = append(l, &p)
		}

		return l
	}

	for i := 0; i < InternalIterations; i++ {
		book := thriftexample.AddressBook{}
		book.People = append(book.People, makePersons(NumberOfPersons)...)

		var data1 []byte

		transport := &thrift.TMemoryBuffer{Buffer: bytes.NewBuffer(data1)}
		proto := thrift.NewTCompactProtocolConf(transport, nil)
		err := book.Write(context.Background(), proto)
		if err != nil {
			log.Fatalln(err)
		}
		out := transport.Bytes()

		data = append(data, float64(len(out)))
		thriftData = append(thriftData, out)
	}

	m, err := stats.Median(data)
	if err != nil {
		log.Fatalln(err)
	}

	return m
}

func testThriftDecode() float64 {
	for i := 0; i < InternalIterations; i++ {
		thriftBytes := thriftData[i]
		t := &thrift.TMemoryBuffer{Buffer: bytes.NewBuffer(thriftBytes)}
		protocol := thrift.NewTCompactProtocolConf(t, nil)
		b := thriftexample.AddressBook{}
		err := b.Read(context.Background(), protocol)
		if err != nil {
			log.Fatalln(err)
		}
	}

	return 0
}
