package main

import (
	"fmt"
	"strings"
)

func multiply(a int,b int) int { // argument, return value type 지정
	return a*b
}

// go -> multiple return이 가능함
// go의 기본 라이브러리는 다양

func lenAndUpper(name string) (int, string){
	return len(name), strings.ToUpper(name)
}

func repeatMe(words ...string){ // 여러 개의 string 값 받음
	fmt.Println(words) // array 형태로 출력됨
}

func main() {
	fmt.Print(multiply(2,3))
	fmt.Print(lenAndUpper("irene"))

	totalLength, upperName := lenAndUpper("nico") // multiple value 저장
	// totalLength, _ := lenAndUpper("nico") // 2번째 값 저장 하지 않음
	fmt.Print(totalLength, upperName)
	repeatMe("irene", "nico", "esther")
}