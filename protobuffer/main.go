package main

import (
	"fmt"
	"io/ioutil"
	"log"

	"github.com/hiago-balbino/hgz-protobuffer-comparison/protobuffer/model"
	"google.golang.org/protobuf/proto"
)

const filename = "./output.txt"

func main() {
	writeOutput()
	readOutput()
	fmt.Println("done!")
}

func readOutput() {
	in, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatalf("error while reading output file: %s", err.Error())
	}

	book := &model.AddressBook{}
	if err := proto.Unmarshal(in, book); err != nil {
		log.Fatalf("error while performing proto unmarshal: %s", err.Error())
	}

	fmt.Println(book)
}

func writeOutput() {
	book := &model.AddressBook{
		People: []*model.Person{
			{
				Name:  "Some Name",
				Id:    1,
				Email: "email@abcd.com",
				Phones: []*model.Person_PhoneNumber{
					{
						Number: "1231231231",
						Type:   model.Person_MOBILE,
					},
				},
			},
		},
	}

	out, err := proto.Marshal(book)
	if err != nil {
		log.Fatalf("error while performing proto marshal: %s", err.Error())
	}

	if err := ioutil.WriteFile(filename, out, 0644); err != nil {
		log.Fatalf("error while writting output: %s", err.Error())
	}
}
