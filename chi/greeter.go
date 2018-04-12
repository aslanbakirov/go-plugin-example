package main

import "fmt"

type greeting string

func (g greeting) Greet() {
	fmt.Println("Hello World In CHIENESE!!")
}

var Greeter greeting
