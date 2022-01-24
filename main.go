package main

import "fmt"

type JsonParser struct {
	data           []byte
	useDefaultData bool
}

// working with json
func (t JsonParser) defaultTestData() []byte {
	return []byte(`{ "fname": "Arv", "lname" : "pyrna", "pin" : "187763"}`)
}

// If you use type of pointer like below
func (t *JsonParser) SetJSON(jsonStr string) {
	t.data = []byte(jsonStr)
}

func (t JsonParser) printData() {
	fmt.Println(t.data)
}

func main() {
	var j JsonParser = JsonParser{
		data: []byte(`{"uname":"arv"}`),
	}
	j.SetJSON(`{"zoo1111111111":"mkv"}`)
	j.printData()
}
