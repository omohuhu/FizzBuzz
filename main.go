package main

import (
	"fmt"
	"strconv"
)

func main() {
	data := 1
	fmt.Println(fizzBuzz(data))
}

func fizzBuzz(data int) string {
	if data % 3 == 0 && data % 5 == 0 {
		return "FizzBuzz"
	} else if data % 3 == 0 {
        return "Fizz"
    } else if data % 5 == 0 {
        return "Buzz"
    } else {
		return strconv.Itoa(data)
	}
}