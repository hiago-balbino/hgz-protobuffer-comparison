package model

// CreateAddressBook is a constructor to create a new default instance of AddressBook
func CreateAddressBook(id int32) *AddressBook {
	return &AddressBook{
		People: []*Person{
			{
				Name:  "Some Name",
				Id:    id,
				Email: "email@abcd.com",
				Phones: []*Person_PhoneNumber{
					{
						Number: "1231231231",
						Type:   Person_MOBILE,
					},
				},
			},
		},
	}
}
