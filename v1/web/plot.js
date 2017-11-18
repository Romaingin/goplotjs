var xhttp = new XMLHttpRequest();
xhttp.open("GET", "/data", false);
xhttp.setRequestHeader("Content-type", "application/json");
xhttp.send();
res = JSON.parse(xhttp.responseText);

switch (res.type) {
	case "matrix":
	plotMatrix(res)
	break
	case "line":
	plotLines(res)
	break
	case "map":
	plotMap(res)
	break
	default:

}

function plotMatrix(res) {
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
}

function plotLines(res) {
	var data = [];
	for (var i = 0; i < res.x.length; i++) {
		var trace = {
			x: res.x[i],
			y: res.y[i],
			name: res.names[i],
			type: 'scatter'
		}

		data.push(trace)
	}

	var layout = {
		title: res.title,
	}

	Plotly.newPlot('chart', data, layout);
}

function plotMap(res) {
	var lon = [4.413841];
	var lat = [50.860589];
	// var symbols = ['monument', 'harbor', 'music'];


	var result = {
		type: 'scattermapbox',
		lon: [4.413841, 4.433841],
		lat: [50.860589, 50.870589],
		mode: 'lines',
		line: {
			width: 2,
			color: 'red'
		},
	};
	data = [result]
	//
	// var data = [result, {
	// 	type: 'scattermapbox',
	// 	mode: 'markers',
	// 	lon: lon,
	// 	lat: lat,
	// 	marker: {
	// 		size: 10,
	// 		color: 'rgb(250,68,0)'
	// 		// symbol: symbols
	// 	}
	// }]

	var layout = {
		title: res.title,
		mapbox: {
			style: 'mapbox://styles/mapbox/satellite-streets-v9',
			domain: {
				x: [0, 1],
				y: [0, 1]
			},
			center: {
				lon: 4.413841,
				lat: 50.860589,
			},
			zoom: 10.8,
			pitch: 0,
			bearing: 0
		},
		showlegend: false
	}

	Plotly.setPlotConfig({
		// mapboxAccessToken: 'pk.eyJ1Ijoicm9tYWluZ2luIiwiYSI6ImNqYTRneTdyZTc4eWsycXA3OTEyYWNxaXAifQ.tLd68m7Re2Ztu9rTki081w'
		mapboxAccessToken: 'pk.eyJ1IjoiZXRwaW5hcmQiLCJhIjoiY2luMHIzdHE0MGFxNXVubTRxczZ2YmUxaCJ9.hwWZful0U2CQxit4ItNsiQ'
	})

	Plotly.newPlot('chart', data, layout, {});
}
