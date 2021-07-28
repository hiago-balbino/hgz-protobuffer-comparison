package model

import (
	"encoding/json"
	"io/ioutil"
	"testing"
)

func BenchmarkMarshalAddressBook(b *testing.B) {
	for i := 0; i < b.N; i++ {
		addressBook := CreateAddressBook(int32(i))

		_, err := json.Marshal(addressBook)
		if err != nil {
			b.Fatal("error to marshal json file")
		}
	}
}

func BenchmarkUnmarshalAddressBook(b *testing.B) {
	data, err := ioutil.ReadFile("../../../files/jsonfile.txt")
	if err != nil {
		b.Fatal("error to read file")
	}

	for i := 0; i < b.N; i++ {
		addressBooks := []AddressBook{}
		err := json.Unmarshal(data, &addressBooks)
		if err != nil {
			b.Fatal("error to unmarshal proto file")
		}
	}
}
