function plotLines (g) {
	var data = []
	for (var l = 0; l < g.data.length; l++) {
		var trace = {
			x: g.data[l].x,
			y: g.data[l].y,
			name: g.data[l].name,
			type: 'scatter'
		}

		data.push(trace)
	}

	var layout = {
		title: g.title
	}

	Plotly.newPlot('chart', data, layout)
}

function plotMatrix (g) {
	var visualize = [{
		z: g.data[0].matrix,
		x: g.data[0].x,
		y: g.data[0].y,
		zsmooth: 'best',
		type: 'heatmap'
	}]

	var layout = {
		title: g.title,
		showscale: true
	}

	Plotly.newPlot('chart', visualize, layout)
}

// TODO
function plotMap (g) {
	document.getElementById("chart").style.width = '100vw'

	var lon = [4.413841]
	var lat = [50.860589]
	// var symbols = ['monument', 'harbor', 'music'];

	var result = {
		type: 'scattermapbox',
		lon: [4.413841, 4.433841],
		lat: [50.860589, 50.870589],
		mode: 'lines',
		line: {
			width: 2,
			color: 'red'
		}
	}
	data = [result]

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
				lat: 50.860589
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

	Plotly.newPlot('chart', data, layout, {})
}
