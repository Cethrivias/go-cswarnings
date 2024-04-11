package charts

import (
	localtypes "cs-warnings/types"
	"strings"

	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/opts"
	"github.com/go-echarts/go-echarts/v2/types"
)

func RenderTreeMap(warnings map[string]localtypes.Warning) *charts.TreeMap {
	treemap := charts.NewTreeMap()

	var processedWarnings = make(map[string]map[string]WarningCounter)

	for _, warning := range warnings {
		path := strings.Split(warning.Project, "/")
		project := path[len(path)-1]
		if processedWarnings[project] == nil {
			processedWarnings[project] = make(map[string]WarningCounter)
		}
		counter := processedWarnings[project][warning.Code]
		counter.Value++
		if counter.Name == "" {
			counter.Name = warning.Code + ": " + warning.Description
		}
		processedWarnings[project][warning.Code] = counter
	}

	nodes := []opts.TreeMapNode{}
	for project, counters := range processedWarnings {
		children := []opts.TreeMapNode{}
		for _, counter := range counters {
			children = append(children, opts.TreeMapNode{Name: counter.Name, Value: counter.Value})
		}
		nodes = append(nodes, opts.TreeMapNode{Name: project, Children: children})
	}

	treemap.SetGlobalOptions(
		charts.WithInitializationOpts(opts.Initialization{Theme: types.ThemeMacarons}),
		charts.WithLegendOpts(opts.Legend{Show: false}),
		charts.WithTitleOpts(opts.Title{
			Title: "Warnings per project",
			Left:  "center",
		}),
		charts.WithTooltipOpts(opts.Tooltip{
			Show:      true,
			Formatter: opts.FuncOpts(ToolTipFormatter),
			Enterable: true,
			TriggerOn: "click",
		}),
		charts.WithToolboxOpts(opts.Toolbox{
			Show:   true,
			Orient: "horizontal",
			Left:   "right",
			Feature: &opts.ToolBoxFeature{
				SaveAsImage: &opts.ToolBoxFeatureSaveAsImage{
					Show: true, Title: "Save as image"},
				Restore: &opts.ToolBoxFeatureRestore{
					Show: true, Title: "Reset"},
			}}),
	)

	treemap.AddSeries("Solution", nodes).
		SetSeriesOptions(
			charts.WithTreeMapOpts(
				opts.TreeMapChart{
					Animation:  true,
					Roam:       true,
					UpperLabel: &opts.UpperLabel{Show: true},
					Levels: &[]opts.TreeMapLevel{
						{ // Series
							ItemStyle: &opts.ItemStyle{
								BorderColor: "#777",
								BorderWidth: 2,
								GapWidth:    1},
							UpperLabel: &opts.UpperLabel{Show: true},
						},
						{ // Level
							ItemStyle: &opts.ItemStyle{
								BorderColor: "#666",
								BorderWidth: 2,
								GapWidth:    1},
							Emphasis: &opts.Emphasis{
								ItemStyle: &opts.ItemStyle{BorderColor: "#555"},
							},
						},
						{ // Node
							ColorSaturation: []float32{0.35, 0.5},
							ItemStyle: &opts.ItemStyle{
								GapWidth:              1,
								BorderWidth:           0,
								BorderColorSaturation: 0.6,
							},
						},
					},
				},
			),
			charts.WithItemStyleOpts(opts.ItemStyle{BorderColor: "#fff"}),
			charts.WithLabelOpts(opts.Label{
				Show:      true,
				Position:  "inside",
				Color:     "White",
				Formatter: opts.FuncOpts(LableFormatter),
			}),
		)

	return treemap
}
