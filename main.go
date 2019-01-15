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
	} else if data == 2 {
        return "2"
    } else {
		return "I Don't Know"
	}
}