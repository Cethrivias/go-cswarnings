package charts

import (
	"cs-warnings/types"
	"io"
	"os"

	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/opts"
)

var ToolTipFormatter = `
function (info) {
	return info.name.slice(0, 7) + '<div class="tooltip-title">' + info.name.slice(8) + '</div>'; 
}
`

func Render(warnings map[string]types.Warning) {
	warningCnts := make(map[string]WarningCount)
	for _, warning := range warnings {
		count := warningCnts[warning.Code]
		count.count++
		if count.title == "" {
			count.title = warning.Description
		}
		warningCnts[warning.Code] = count
	}

	pieData := make([]opts.PieData, 0)

	for key, count := range warningCnts {
		pieData = append(pieData, opts.PieData{Name: key + ": " + count.title, Value: count.count})

	}

	pie := charts.NewPie()
	pie.SetGlobalOptions(
		charts.WithLegendOpts(opts.Legend{Show: false, Orient: "vertical", Left: "0"}),
		charts.WithTooltipOpts(opts.Tooltip{Show: true, Formatter: opts.FuncOpts(ToolTipFormatter), Enterable: true, TriggerOn: "click"}),
	)
	pie.AddSeries("pie", pieData).SetSeriesOptions(
		charts.WithLabelOpts(opts.Label{Show: true, Formatter: "{c} ({d}%)"}),
	)
	f, err := os.Create("pie.html")
	if err != nil {
		panic(err)
	}
	pie.Render(io.MultiWriter(f))

}

type WarningCount struct {
	count int
	title string
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
