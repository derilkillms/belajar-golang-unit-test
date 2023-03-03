package helper

import "testing"

func TestHelloWorld(t *testing.T) {
	result := HelloWorld("Deril")
	if result != "Hello Deril" {
		panic("Result is not Hello Deril")
	}
}

func TestHelloWorldDeril(t *testing.T) {
	result := HelloWorld("MDeril")
	if result != "Hello MDeril" {
		panic("Result is not Hello Deril")
	}
}
