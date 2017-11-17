var xhttp = new XMLHttpRequest();
xhttp.open("GET", "/data", false);
xhttp.setRequestHeader("Content-type", "application/json");
xhttp.send();
res = JSON.parse(xhttp.responseText);

var visualize = [{
	z: res.data,
	x: res.x,
	y: res.y,
	type: 'contour',
	contours: {
		coloring: 'heatmap'
	}
}];

var layout = {
	title: res.title,
	showscale: true,
};

Plotly.newPlot('chart', visualize, layout);
