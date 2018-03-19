var xhttp = new XMLHttpRequest()
xhttp.open("GET", "/allgraphs", false)
xhttp.setRequestHeader("Content-type", "application/json")
xhttp.send()
graphs = JSON.parse(xhttp.responseText)

// Display all the available charts
for (var g in graphs) {
	var createTR = document.createElement('tr')
	var createTDtitle = document.createElement('td')
	var createTDA = document.createElement('a')
	var createTitle = document.createTextNode(graphs[g].title)
	var createTDtype = document.createElement('td')
	var createType = document.createTextNode(graphs[g].type)
	createTDA.setAttribute('href', '/graph.html?id=' + g)
	createTDA.setAttribute('target', '_blank')

	createTR.appendChild(createTDtitle)
	createTDtitle.appendChild(createTDA)
	createTDA.appendChild(createTitle)
	createTR.appendChild(createTDtype)
	createTDtype.appendChild(createType)
	document.getElementById("table").appendChild(createTR)
}
