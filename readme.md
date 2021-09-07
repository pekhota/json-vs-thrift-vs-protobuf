# json-vs-thrift-vs-protobuf

Comparison between json, thrift and protobuf protocols 

## Results 

| Protocol | Encoding time | Decoding time | Payload size (bytes) | Test duration |
|----------|---------------|---------------|----------------------|---------------|
| Json     | 216.74875ms   | 210.512125ms  | 17715                | 2.199822167s  |
| Protobuf | 198.921792ms  | 32.473791ms   | 7811                 | 1.315487459s  |
| Thrift   | 224.548ms     | 46.453167ms   | 7853                 | 1.438012417s  |

## Summary 

1. Encoding time +- the same
2. Decoding time (on the server): proto 7x faster then json
3. Decoding time (on the server): thrift 5x faster then json
4. Payload size: 2 times smaller

## Test scenario

1. Create golang object with faked data
2. Marshal it via json, thrift, proto encoder 1000 times
3. Unmarshal it 1000 times
4. Repeat 2-3 5 times
5. Calculate median for each metric via stats 

### Example object to marshal/unmarshal

```go
type phoneType int

const (
	MOBILE phoneType = iota;
	HOME;
	WORK;
)

type JsonPhoneNumber struct {
	Number    string    `json:"number"`
	PhoneType phoneType `json:"phone_type"`
}

type Timestamp struct {
	Seconds int64 `json:"seconds"`
	Nanos int32 `json:"nanos"`
}

type JsonPerson struct {
	Name string `json:"name"`
	Id int32 `json:"id"`
	Email string `json:"email"`

	Phones      []JsonPhoneNumber `json:"phones"`
	LastUpdated Timestamp         `json:"last_updated"`
}

type JsonAddressBook struct {
	People []JsonPerson `json:"people"`
}
```

Number of person generated per each address book: 100 


| Operation system | macOS Big Sur |
|------------------|---------------|
| CPU              | apple m1      |
| Memory           | 16 gb         |
| Cores            | 8             |



