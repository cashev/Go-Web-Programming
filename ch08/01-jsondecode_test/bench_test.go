package main

import (
	"testing"
)

// benchmark decoding
func BenchmarkDecode(b *testing.B) {
	for i := 0; i < b.N; i++ {
		decode("post.json")
	}
}

func BenchmarkUnmarshal(b *testing.B) {
	for i := 0; i < b.N; i++ {
		unmarshal("post.json")
	}
}
