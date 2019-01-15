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

	t.Run("fizzBuzz5", func(t *testing.T) {
		get := fizzBuzz(5)
		want := "Buzz"
		if want != get {
			t.Errorf("Error %s %s", want, get)
		}
	})

	t.Run("fizzBuzz6", func(t *testing.T) {
		get := fizzBuzz(6)
		want := "Fizz"
		if want != get {
			t.Errorf("Error %s %s", want, get)
		}
	})

	t.Run("fizzBuzz7", func(t *testing.T) {
		get := fizzBuzz(7)
		want := "7"
		if want != get {
			t.Errorf("Error %s %s", want, get)
		}
	})

	t.Run("fizzBuzz8", func(t *testing.T) {
		get := fizzBuzz(8)
		want := "8"
		if want != get {
			t.Errorf("Error %s %s", want, get)
		}
	})

	t.Run("fizzBuzz9", func(t *testing.T) {
		get := fizzBuzz(9)
		want := "Fizz"
		if want != get {
			t.Errorf("Error %s %s", want, get)
		}
	})

	t.Run("fizzBuzz10", func(t *testing.T) {
		get := fizzBuzz(10)
		want := "Buzz"
		if want != get {
			t.Errorf("Error %s %s", want, get)
		}
	})

	t.Run("fizzBuzz11", func(t *testing.T) {
		get := fizzBuzz(11)
		want := "11"
		if want != get {
			t.Errorf("Error %s %s", want, get)
		}
	})

	t.Run("fizzBuzz12", func(t *testing.T) {
		get := fizzBuzz(12)
		want := "Fizz"
		if want != get {
			t.Errorf("Error %s %s", want, get)
		}
	})

	t.Run("fizzBuzz13", func(t *testing.T) {
		get := fizzBuzz(13)
		want := "13"
		if want != get {
			t.Errorf("Error %s %s", want, get)
		}
	})

	t.Run("fizzBuzz14", func(t *testing.T) {
		get := fizzBuzz(14)
		want := "14"
		if want != get {
			t.Errorf("Error %s %s", want, get)
		}
	})

	t.Run("fizzBuzz15", func(t *testing.T) {
		get := fizzBuzz(15)
		want := "FizzBuzz"
		if want != get {
			t.Errorf("Error %s %s", want, get)
		}
	})

}