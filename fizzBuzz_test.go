package main

import (
	"testing"
)

func TestFizz(t *testing.T) {

	var rst string
	rst = checkFizzBuzz(99)

	if rst != "fizz" {
		t.Error("Should return fizz")
	}

}

func TestBuzz(t *testing.T) {

	var rst string
	rst = checkFizzBuzz(25)

	if rst != "buzz" {
		t.Error("Should return buzz")
	}

}

func TestFizzBuzz(t *testing.T) {

	var rst string
	rst = checkFizzBuzz(60)

	if rst != "fizzbuzz" {
		t.Error("Should return fizzbuzz")
	}

}

func TestDigit(t *testing.T) {

	var rst string
	rst = checkFizzBuzz(8)

	if rst != "8" {
		t.Error("Should return 8")
	}

}
