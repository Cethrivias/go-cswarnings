package charts

var ToolTipFormatter = `
function (info) {
	return info.name.slice(0, 7) + '<div class="tooltip-title">' + info.name.slice(8) + '</div>'; 
}
`

var LableFormatter = `
function (info) {
	let idx = info.name.indexOf(':');
	if (idx >= 0) {
	  return info.name.slice(0, idx) + ' (' + info.value + ')'; 
	}
	return '(' + info.value + ') ' + info.name 
}
`
