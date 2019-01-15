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
}