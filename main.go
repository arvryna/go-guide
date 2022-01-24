package main

import (
	"fmt"

	"github.com/arvryna/go-guide/src/jsman"
)

type Contact struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Pin       int    `json:"pin"`
}

func main() {
	var j jsman.JsonParser = jsman.JsonParser{
		JsonStr: `{"firstName":"arv", "lastName": "ryna", "pin": 17643}`,
	}

	c := Contact{}
	j.Unpack(&c)

	fmt.Printf("Unmarshalled data: %#v", c)
}
