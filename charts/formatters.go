package charts

var ToolTipFormatter = `
function (info) {
	return info.name.slice(0, 7) + '<div class="tooltip-title">' + info.name.slice(8) + '</div>'; 
}
`

var LableFormatter = `
function (info) {
	return info.name.slice(0, 6) + ' (' + info.value + ')'; 
}
`
