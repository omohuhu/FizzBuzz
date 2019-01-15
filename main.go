package main

import (
	"fmt"
)

func main() {
	data := 1
	fmt.Println(fizzBuzz(data))
}

func fizzBuzz(data int) string {
	if data == 1 {
		return "1"
	}
	return "I Don't Know"
}