package main

import (
	"fmt"
	"strings"
)

func multiply(a, b int) int {
	return a * b
}

func lenAndUpper(name string) (int, string) {
	return len(name), strings.ToUpper(name)
}

func lenAndUpper2(name string) (length int, uppercase string) { // naked return
	defer fmt.Println("I'm done") // 함수가 리턴할 때 수행할 작업
	length = len(name)
	uppercase = strings.ToUpper(name)
	return
}

func repeatMe(words ...string) {
	fmt.Println(words)
}

func main() {
	fmt.Println(multiply(2, 2))
	len, upperName := lenAndUpper("hyunjin")
	fmt.Println(len, upperName)
	len2, _ := lenAndUpper("hyunjin")
	fmt.Println(len2)

	fmt.Println(lenAndUpper2("kim"))

	repeatMe("abc", "def", "ghi", "jkl")
}