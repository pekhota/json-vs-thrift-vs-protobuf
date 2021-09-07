package main

import (
	"encoding/json"
	"github.com/bxcodec/faker/v3"
	"github.com/montanaflynn/stats"
	"log"
	json_example "speed-test/pkg/json-example"
	"time"
)

var jsonData [][]byte

func init()  {
	jsonInit()
}

func jsonInit()  {
	jsonData = make([][]byte, 0)
}

func testJsonEncode() float64  {
	data := make([]float64, 0)

	makeJsonPeople := func(n int) []json_example.JsonPerson {
		l := make([]json_example.JsonPerson, 0)

		for i := 0; i < n; i++ {
			t := time.Now()
			l = append(l, json_example.JsonPerson{
				Id: int32(i),
				Name: faker.Name(),
				Email: faker.Email(),
				Phones: []json_example.JsonPhoneNumber{
					{
						Number: faker.Phonenumber(), PhoneType: json_example.HOME,
					},
				},
				LastUpdated: json_example.Timestamp{
					Seconds: t.Unix(),
					Nanos:   int32(t.Nanosecond()),
				},
			})
		}
		return l
	}

	for i := 0; i < InternalIterations; i++ {
		book := json_example.JsonAddressBook{}

		book.People = append(book.People, makeJsonPeople(NumberOfPersons)...)

		out, err := json.Marshal(book)
		if err != nil {
			log.Fatalln(err)
		}

		jsonData = append(jsonData, out)
		data = append(data, float64(len(out)))
	}

	m, err := stats.Median(data)
	if err != nil {
		log.Fatalln(err)
	}

	return m
	//log.Println("Median Json size (in bytes): ", m)
}

func testJsonDecode() float64 {
	for i := 0; i < InternalIterations; i++ {
		jsonBytes := jsonData[i]

		book := json_example.JsonAddressBook{}

		err := json.Unmarshal(jsonBytes, &book)
		if err != nil {
			log.Fatalln(err)
		}
	}
	return 0
}
