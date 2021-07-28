package model

import (
	"io/ioutil"
	"testing"

	"google.golang.org/protobuf/proto"
)

func BenchmarkMarshalAddressBook(b *testing.B) {
	for i := 0; i < b.N; i++ {
		addressBook := CreateAddressBook(int32(i))

		_, err := proto.Marshal(addressBook)
		if err != nil {
			b.Fatal("error to marshal proto file")
		}
	}
}

func BenchmarkUnmarshalAddressBook(b *testing.B) {
	data, err := ioutil.ReadFile("../../../files/protofile.txt")
	if err != nil {
		b.Fatal("error to read file")
	}

	for i := 0; i < b.N; i++ {
		addressBook := AddressBook{}
		err := proto.Unmarshal(data, &addressBook)
		if err != nil {
			b.Fatal("error to unmarshal proto file")
		}
	}
}
