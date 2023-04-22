package main

import (
	"cs-warnings/charts"
	"cs-warnings/parser"
	"fmt"
	"os"
)

func main() {
	warnings := parser.ParseWarnings(os.Stdin)
	fmt.Printf("Warnings count: %v\n", len(warnings))

	charts.Render(warnings)
}
