var xhttp = new XMLHttpRequest()
xhttp.open("GET", "/visualizations", false);
xhttp.setRequestHeader("Content-type", "application/json")
xhttp.send()
res = JSON.parse(xhttp.responseText)

// Display all the available charts
for (var t in res.titles) {
	var createTR = document.createElement('tr')
	var createTDtitle = document.createElement('td')
	var createTDA = document.createElement('a')
	var createTitle = document.createTextNode(res.titles[t])
	var createTDtype = document.createElement('td')
	var createType = document.createTextNode(res.types[t])
	createTDA.setAttribute('href', '/plot.html?id=' + t)
	createTDA.setAttribute('target', '_blank')

	createTR.appendChild(createTDtitle)
	createTDtitle.appendChild(createTDA)
	createTDA.appendChild(createTitle)
	createTR.appendChild(createTDtype)
	createTDtype.appendChild(createType)
	document.getElementById("table").appendChild(createTR)
}
