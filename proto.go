package main

import (
	"github.com/bxcodec/faker/v3"
	"github.com/golang/protobuf/proto"
	"github.com/montanaflynn/stats"
	pb "github.com/pekhota/json-vs-thrift-vs-protobuf/internal/pkg/protobuf-example"
	"google.golang.org/protobuf/types/known/timestamppb"
	"log"
)

var protoData [][]byte

func init() {
	protoInit()
}

func protoInit() {
	protoData = make([][]byte, 0)
}

func testProtoEncode() float64 {
	data := make([]float64, 0)

	makePersons := func(n int) []*pb.Person {
		l := make([]*pb.Person, 0)

		for i := 0; i < n; i++ {
			p := pb.Person{
				Id:    int32(i),
				Name:  faker.Name(),
				Email: faker.Email(),
				Phones: []*pb.PhoneNumber{
					{
						Number: faker.Phonenumber(), Type: pb.PhoneType_HOME,
					},
				},
				LastUpdated: timestamppb.Now(),
			}
			l = append(l, &p)
		}

		return l
	}

	for i := 0; i < InternalIterations; i++ {
		book := &pb.AddressBook{}
		book.People = append(book.People, makePersons(NumberOfPersons)...)
		out, err := proto.Marshal(book)
		if err != nil {
			log.Fatalln("Failed to encode address book:", err)
		}
		data = append(data, float64(len(out)))
		protoData = append(protoData, out)
	}

	m, err := stats.Median(data)
	if err != nil {
		log.Fatalln(err)
	}

	return m
	//log.Println("Median proto size (in bytes): ", m)
}

func testProtoDecode() float64 {
	for i := 0; i < InternalIterations; i++ {
		protoBytes := protoData[i]

		book := &pb.AddressBook{}
		err := proto.Unmarshal(protoBytes, book)
		if err != nil {
			log.Fatalln(err)
		}
	}
	return 0
}
