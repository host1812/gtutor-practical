package main

import "testing"

func BenchmarkRemoveDuplicateChars(b *testing.B) {
	// run the Fib function b.N times
	for n := 0; n < b.N; n++ {
		RemoveDuplicateChars("Hello World! 😊. I don't know you 😊.")
	}
}

func BenchmarkRemoveDuplicateChars2(b *testing.B) {
	// run the Fib function b.N times
	for n := 0; n < b.N; n++ {
		RemoveDuplicateChars2("Hello World! 😊. I don't know you 😊.")
	}
}
