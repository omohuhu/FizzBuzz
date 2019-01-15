package main

import "testing"

func TestFizzBuzz(t *testing.T) {
	t.Run("fizzBuzz1", func(t *testing.T) {
		get := fizzBuzz(1)
		want := "1"
		if want != get {
			t.Errorf("Error %s %s", want, get)
		}
	})

	t.Run("fizzBuzz2", func(t *testing.T) {
		get := fizzBuzz(2)
		want := "2"
		if want != get {
			t.Errorf("Error %s %s", want, get)
		}
	})

	t.Run("fizzBuzz3", func(t *testing.T) {
		get := fizzBuzz(3)
		want := "Fizz"
		if want != get {
			t.Errorf("Error %s %s", want, get)
		}
	})

	t.Run("fizzBuzz4", func(t *testing.T) {
		get := fizzBuzz(4)
		want := "4"
		if want != get {
			t.Errorf("Error %s %s", want, get)
		}
	})

}