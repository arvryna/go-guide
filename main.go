package main

import (
	"fmt"

	"github.com/arvryna/go-guide/src/jsman"
)

type Contact struct {
	FirstName string `json:"FirstName"`
	LastName  string `json:"lastName"`
	Pin       int    `json:"pin"`
}

func main() {
	var j jsman.JsonParser = jsman.JsonParser{
		JsonStr: `{"firstName":"arv", "lastName": "ryna", "pin": 17643}`,
	}

	// Unpacking
	c := Contact{}
	j.Unpack(&c)
	fmt.Printf("Unmarshalled data: %#v \n\n", c)

	// Packing
	c.FirstName = "Ash"
	fmt.Printf("Marshalled data: ")
	fmt.Println(j.Pack(&c))
}
