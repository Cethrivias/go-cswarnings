package main

import (
	"cs-warnings/charts"
	"cs-warnings/parser"
	"fmt"
	"io"
	"os"

	"github.com/go-echarts/go-echarts/v2/components"
)

func main() {
	warnings := parser.ParseWarnings(os.Stdin)
	fmt.Printf("Warnings count: %v\n", len(warnings))

	page := components.NewPage().AddCharts(
		charts.RenderPie(warnings),
		charts.RenderTreeMap(warnings),
	)

	f, err := os.Create("charts.html")
	if err != nil {
		panic(err)
	}
	page.Render(io.MultiWriter(f))
}
