// Inject HTML
document.getElementById("app").innerHTML = '<div id="chart" style="height:550px;width:100wh;margin:0 auto"></div>'

const gr = graph.data

// Plot according to the graph's type
switch (gr.type) {
case "line":
	plotLines(gr)
	break
case "matrix":
	plotMatrix(gr)
	break
}
