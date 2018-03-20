package goplotjs

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"path"
	"runtime"
	"strconv"

	"github.com/zserge/webview"
)

// Graph instance modified by the user
var graphs Graphs
var browserInstance = 0

// SetTitle set the title of the last graph
func SetTitle(t string) {
	graphs.setTitle(t)
}

// Show displays the graphs in a window or in a browser
func Show(mode ...string) {
	if graphs.len() == 0 {
		println("Nothing to plot...")
		return
	}

	var m string
	if len(mode) == 0 {
		m = "window"
	} else {
		m = mode[0]
	}

	switch m {
	case "window":
		showWindow()
	case "browser":
		showBrowser()
	}
}

// showWindow open a webview with the graphs
func showWindow() {
	// Fetch info about module directory
	_, filename, _, _ := runtime.Caller(0)

	// Open assets
	datPlotly, _ := ioutil.ReadFile(path.Dir(filename) + "/assets/common/plotly.js")
	datPlot, _ := ioutil.ReadFile(path.Dir(filename) + "/assets/common/plot.js")
	datApp, _ := ioutil.ReadFile(path.Dir(filename) + "/assets/webview/app.js")

	webviews := make([]webview.WebView, graphs.len())

	for i := 0; i < graphs.len(); i++ {
		// Setup a new window
		webviews[i] = webview.New(webview.Settings{
			Title:  "GoPlotJS " + graphs.get(i).Title,
			Debug:  true,
			Width:  800,
			Height: 600,
		})
		defer webviews[i].Exit()

		// Inject graph object
		webviews[i].Dispatch(func() {
			webviews[i].Bind("graph", graphs.get(i))
		})

		// Inject CSS and JS
		webviews[i].Eval(string(datPlotly))
		webviews[i].Eval(string(datPlot))
		webviews[i].Eval(string(datApp))

		// Run the windows
		if i == graphs.len()-1 {
			webviews[0].Run()
		}
	}
}

// showBrowser runs a local HTTP server on port 8080 to display charts
func showBrowser(blocking ...bool) {
	_, filepath, _, _ := runtime.Caller(0)

	http.Handle("/", http.FileServer(http.Dir(path.Dir(filepath)+"/assets")))
	http.HandleFunc("/allgraphs", routeAllGraphs)
	http.HandleFunc("/graph", routeGraphByID)
	port := "8080"

	// Lauch HTTP Server in concurrent routine
	if browserInstance == 0 {
		browserInstance++
		go http.ListenAndServe("localhost:"+port, nil)
	} else {
		panic("Visualization already in progress")
	}

	fmt.Println("\x1b[34mVisit http://localhost:8080 to see the graphs\x1b[0m")

	if len(blocking) == 0 || blocking[0] {
		var input string
		fmt.Println("Press Ctrl-C to stop...")
		fmt.Scanln(&input)
	}
}

// routeGraphById respond with the graph at the given id
func routeGraphByID(w http.ResponseWriter, req *http.Request) {
	if req.Body == nil {
		http.Error(w, "Please send a request body", 400)
		return
	}

	id, err := strconv.Atoi(req.URL.Query().Get("id"))
	if err != nil {
		http.Error(w, "Invalid id", 400)
		return
	}

	// Send response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(graphs.get(id))
}

// routeVisualizations respond with all available graphs
func routeAllGraphs(w http.ResponseWriter, req *http.Request) {
	if req.Body == nil {
		http.Error(w, "Please send a request body", 400)
		return
	}

	// Send response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(graphs)
}
