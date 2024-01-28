package main

import "fmt"

func main() {
	const name string = "nico" // go lang은 type language

	var name2 string = "irene"
	name2 = "nico"
	fmt.Print(name2)

	name3:="irene" // 축약: 자동 타입 찾기, func 내에서만 가능
	fmt.Print(name3)
}