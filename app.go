package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type person struct {
	First string
	Last  string
}

func main() {
	// data_marshal()
	// data_unmarshal()
	http_services()
}

func http_services() {
	// declare service endpoint
	http.HandleFunc("/encode", encode_func)
	http.HandleFunc("/decode", decode_func)
	http.ListenAndServe(":8080", nil)
}

// curl command consumming the encode service
// curl localhost:8080/encode
func encode_func(w http.ResponseWriter, r *http.Request) {
	p1 := person{
		First: "Andrew",
		Last:  "Zou",
	}

	err := json.NewEncoder((w)).Encode(p1)
	if err != nil {
		log.Println("Encoded bad data!", err)
	}
}

// curl reqest consumming the decode service
// curl -XGET -H "Content-type: application/json" -d '{"First": "Andrew", "Last": "ZOU"}' 'localhost:8080/decode'
func decode_func(w http.ResponseWriter, r *http.Request) {
	var p1 person
	err := json.NewDecoder(r.Body).Decode(&p1)
	if err != nil {
		log.Println("Decode bad request data: ", err)
	}
	log.Println("decode request body: ", p1)
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
