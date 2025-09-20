package main

import "testing"

func TestGreetingSpecificJohn(t *testing.T) {
	greeting := CreateGreeting("John")
	if greeting != "Welcome, John\n" {
		t.Errorf("Greeting was incorrect, got: %s, want: %s.", greeting, "Hello, John\n")
	}
}

func TestGreetingSpecificDemo(t *testing.T) {
	greeting := CreateGreeting("Demo")
	if greeting != "Welcome, Demo\n" {
		t.Errorf("Greeting was incorrect, got: %s, want: %s.", greeting, "Hello, Demo\n")
	}
}





