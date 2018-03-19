var url_string = window.location.href
var url = new URL(url_string)
var id = url.searchParams.get("id")

var xhttp = new XMLHttpRequest()
xhttp.open("GET", "/graph?id=" + id, true)

xhttp.onload = function (e) {
	if (xhttp.status === 200) {
		graph = JSON.parse(xhttp.responseText)

		document.title = graph.title
		document.getElementById('chart').style.width = '100vh'

		switch (graph.type) {
		case "matrix":
			plotMatrix(graph)
			break
		case "line":
			plotLines(graph)
			break
		case "map":
			plotMap(graph)
			break
		default:
			console.log("Bad Type")
		}
	}
}

xhttp.send()
