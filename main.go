package main // for compile

import (
	"fmt" // package for format

	"example.com/m/something"
)

func main(){ // crucial for go codes - entry point
	fmt.Print("Hello World!")
	something.SayHello() // 대문자 함수 - 다른 패키지에서 import 해옴
	// ctrl (command) + click -> 함수 조사 가능
}