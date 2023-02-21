package model

import (
	"encoding/json"
	"os"
	"testing"
)

func BenchmarkMarshalJSON(b *testing.B) {
	for i := 0; i < b.N; i++ {
		addressBook := CreateAddressBook(int32(i))

		_, err := json.Marshal(addressBook)
		if err != nil {
			b.Fatal("error to marshal json file")
		}
	}
}

func BenchmarkUnmarshalJSON(b *testing.B) {
	data, err := os.ReadFile("../../../files/jsonfile.txt")
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
