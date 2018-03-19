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
	var m string
	if len(mode) == 0 {
		m = "browser"
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

// showWindow open a webview with the first graph
func showWindow() {
	if graphs.len() == 0 {
		println("Nothing to plot...")
		return
	}

	w := webview.New(webview.Settings{
		Title:  "GoPlotJS " + graphs.get(0).Title,
		Debug:  true,
		Width:  800,
		Height: 600,
	})
	defer w.Exit()

	w.Dispatch(func() {
		// Inject graph
		w.Bind("graph", graphs.get(0))

		// Inject CSS and JS
		// Fetch info about module directory
		_, filename, _, _ := runtime.Caller(0)

		// Open assets
		dat, _ := ioutil.ReadFile(path.Dir(filename) + "/assets/common/plotly.js")
		w.Eval(string(dat))

		dat, _ = ioutil.ReadFile(path.Dir(filename) + "/assets/common/plot.js")
		w.Eval(string(dat))

		dat, _ = ioutil.ReadFile(path.Dir(filename) + "/assets/webview/app.js")
		w.Eval(string(dat))
	})

	// Run the webview
	w.Run()
}

// showBrowser runs a local HTTP server on port 8080 to display charts
func showBrowser(blocking ...bool) {
	_, runContext, _, ok := runtime.Caller(0)
	if !ok {
		panic("No caller information")
	}

	http.Handle("/", http.FileServer(http.Dir(path.Dir(runContext)+"/assets")))
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

	fmt.Println("\x1b[34mVisit localhost:8080 to see the graphs\x1b[0m")

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
