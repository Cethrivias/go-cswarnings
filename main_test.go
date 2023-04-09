package main

import (
	"fmt"
	"os"
	"testing"
)

func BenchmarkParseWarnings(b *testing.B) {
	for i := 0; i < b.N; i++ {
		file, err := os.Open("test.txt")
		if err != nil {
			fmt.Println("Could not read the file")

			return
		}
		parseWarnings(file)
	}
}
