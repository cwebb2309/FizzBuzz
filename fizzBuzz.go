package main

import (
	"fmt"
	"strconv"
)

func main() {

	var fizzBuzz string

	for i := 0; i <= 100; i++ {
		fizzBuzz = checkFizzBuzz(i)
		fmt.Println(fizzBuzz)
	}
}

func checkFizzBuzz(number int) string {

	var retvalue string

	if number%5 == 0 && number%3 == 0 {
		retvalue = "fizzbuzz"
	} else if number%5 == 0 {
		retvalue = "buzz"
	} else if number%3 == 0 {
		retvalue = "fizz"
	} else {
		retvalue = strconv.Itoa(number)
	}

	return retvalue
}
