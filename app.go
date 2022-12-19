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
	fmt.Println(string(bs))
}
