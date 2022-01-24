package main

import (
	"github.com/arvryna/go-guide/src/jsman"
)

func main() {
	var j jsman.JsonParser = jsman.JsonParser{
		Data: []byte(`{"uname":"arv"}`),
	}
	j.SetJSON(`{"zoo1111111111":"mkv"}`)
	j.PrintData()
}
