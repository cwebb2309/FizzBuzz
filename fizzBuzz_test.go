package main

import (
	"testing"
)

type fizzBuzzTest struct {
	n        int
	expected string
}

var fizzBuzzTests = []fizzBuzzTest{
	{99, "fizz"}, {25, "buzz"}, {60, "fizzbuzz"}, {8, "8"},
}

func TestCheckFizzBuzz(t *testing.T) {
	for _, tt := range fizzBuzzTests {
		actual := checkFizzBuzz(tt.n)
		if actual != tt.expected {
			t.Errorf("Fib(%d): expected %s, actual %s", tt.n, tt.expected, actual)
		}
	}
}
