package json_example

type phoneType int

const (
	MOBILE phoneType = iota;
	HOME;
	WORK;
)

type JsonPhoneNumber struct {
	Number string `json:"number"`
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

	Phones []JsonPhoneNumber `json:"phones"`
	LastUpdated Timestamp `json:"last_updated"`
}

type JsonAddressBook struct {
	People []JsonPerson `json:"people"`
}
