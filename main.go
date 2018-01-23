package main

import (
	"fmt"
)


func main() {
	var my_token Token
	var one, two, first_channel chan Token
	var i int
	my_token = Token{4, 8}
	first_channel = make(chan Token)
	one = first_channel
	two = make(chan Token)

	for i = 1; i < 50; i++ {
		var data Data
		data = Data{i, one, two}
		go test(data)


		one = two
		two = make(chan Token)
	}

	first_channel <- my_token
}


type Token struct {
	RecipientName int
	TTL       int
}

type Data struct {
	name int
	input       chan Token
	output       chan Token
}

func test(data Data) {
	t := <- data.input

	if data.name == t.RecipientName {
		fmt.Println(data.name)
		return
	}

	t.TTL = t.TTL - 1

	if t.TTL > 0 {
		data.output <- t
	}
}
