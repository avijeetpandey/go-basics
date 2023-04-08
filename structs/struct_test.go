package main

import "testing"

func TestPersonWithContactInfoCreation(t *testing.T) {
	jim := personWithContact{
		firstName: "Jim",
		lastName:  "Kick",
		contact: contactInfo{
			email: "test@g.cc",
			zip:   122002,
		},
	}

	if jim.firstName != "Jim" {
		t.Errorf("Error initialising firstName")
	}

	if jim.lastName != "Kick" {
		t.Errorf("Error initialising lastName")
	}

	if jim.contact.email != "test@g.cc" {
		t.Errorf("Error initialising contactInfo")
	}

	if jim.contact.zip != 122002 {
		t.Errorf("Error initialising zip")
	}
}
