package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("saying hello to people", func(t *testing.T){
		got := Hello("world")
		want := "Hello, world"
		
		assertCorrectMessage(t, got, want)
	})
	
	t.Run("saying hello world when empty string is passed", func(t *testing.T){
		got := Hello("")
		want := "Hello, world"
		
		assertCorrectMessage(t,got,want)
	})
}

func assertCorrectMessage(t testing.TB, got string, want string) {
	t.Helper()
	
	if got != want { 
		t.Errorf("got %s, want %s", got, want)
	}
}