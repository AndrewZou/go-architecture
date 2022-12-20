package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type person struct {
	First string
	Last  string
}

func main() {
	data_marshal()
	data_unmarshal()
}

func data_marshal() {
	p1 := person{
		First: "Andrew",
		Last:  "Zou",
	}

	p2 := person{
		First: "Wendy",
		Last:  "Zou",
	}

	// create a slice/array of objects
	xp := []person{p1, p2}

	// Call json Marshal
	bs, err := json.Marshal(xp)
	if err != nil {
		log.Panic(err)
	}
	fmt.Println("Marshaled data: ", string(bs))
}

func data_unmarshal() {
	p1 := person{
		First: "Andrew",
		Last:  "Zou",
	}

	p2 := person{
		First: "Wendy",
		Last:  "Zou",
	}

	// create a slice/array of objects
	xp := []person{p1, p2}

	// Call json Marshal
	bs, err := json.Marshal(xp)
	if err != nil {
		log.Panic(err)
	}

	// unMarshal the marshaled data
	xp2 := []person{}

	err2 := json.Unmarshal(bs, &xp2)
	if err2 != nil {
		log.Panic(err2)
	}
	fmt.Println("UnMarshal data: ", xp2)
}
