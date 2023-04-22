package charts

import (
	localtypes "cs-warnings/types"

	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/opts"
	"github.com/go-echarts/go-echarts/v2/types"
)

func RenderPie(warnings map[string]localtypes.Warning) *charts.Pie {
	warningCnts := make(map[string]WarningCounter)
	for _, warning := range warnings {
		counter := warningCnts[warning.Code]
		counter.Value++
		if counter.Name == "" {
			counter.Name = warning.Description
		}
		warningCnts[warning.Code] = counter
	}

	pieData := make([]opts.PieData, 0)

	for code, counter := range warningCnts {
		pieData = append(pieData, opts.PieData{Name: code + ": " + counter.Name, Value: counter.Value})
	}

	pie := charts.NewPie()
	pie.SetGlobalOptions(
		charts.WithInitializationOpts(opts.Initialization{Theme: types.ThemeMacarons}),
		charts.WithTitleOpts(opts.Title{
			Title: "All warnings across all projects",
			Left:  "center",
		}),
		charts.WithLegendOpts(opts.Legend{Show: false}),
		charts.WithTooltipOpts(opts.Tooltip{
			Show:      true,
			Formatter: opts.FuncOpts(ToolTipFormatter),
			Enterable: true,
			TriggerOn: "click",
		}),
	)
	pie.AddSeries("pie", pieData).SetSeriesOptions(
		charts.WithLabelOpts(opts.Label{Show: true, Formatter: opts.FuncOpts(LableFormatter)}),
	)

	return pie
}

type WarningCounter struct {
	Value int
	Name  string
}

// return JSON.stringify(Object.keys(info))

// var formatUtil = echarts.format;
// var value = info.value;
// var treePathInfo = info.treePathInfo;
// var treePath = [];
// for (var i = 1; i < treePathInfo.length; i++) {
// 	treePath.push(treePathInfo[i].name);
// }
// return ['<div class="tooltip-title">' + formatUtil.encodeHTML(treePath.join('/')) + '</div>',
// 	'Disk Usage: ' + formatUtil.addCommas(value) + ' KB',
// 	].join('');
