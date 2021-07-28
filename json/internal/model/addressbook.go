package model

// AddressBook is a struct to represent a book with people who have written
type AddressBook struct {
	People []Person `json:"people"`
}

// Person is a struct to represent each person that have written a book
type Person struct {
	Name   string  `json:"name"`
	Id     int32   `json:"id"`
	Email  string  `json:"email"`
	Phones []Phone `json:"phones"`
}

// Phone is a struct to represent a phone number for each person
type Phone struct {
	Number string `json:"number"`
	Type   int32  `json:"type"`
}

// CreateAddressBook is a constructor to create a new default instance of AddressBook
func CreateAddressBook(id int32) AddressBook {
	return AddressBook{
		People: []Person{
			{
				Name:  "Some Name",
				Id:    id,
				Email: "email@abcd.com",
				Phones: []Phone{
					{
						Number: "1231231231",
						Type:   1,
					},
				},
			},
		},
	}
}
