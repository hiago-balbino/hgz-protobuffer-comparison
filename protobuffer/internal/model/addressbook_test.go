package model

import (
	"os"
	"testing"

	"google.golang.org/protobuf/proto"
)

func BenchmarkMarshalProtobuf(b *testing.B) {
	for i := 0; i < b.N; i++ {
		addressBook := CreateAddressBook(int32(i))

		_, err := proto.Marshal(addressBook)
		if err != nil {
			b.Fatal("error to marshal proto file")
		}
	}
}

func BenchmarkUnmarshalProtobuf(b *testing.B) {
	data, err := os.ReadFile("../../../files/protofile.txt")
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
