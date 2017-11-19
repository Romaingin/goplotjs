package goplotjs

import (
	"encoding/json"
	"fmt"
	"net/http"
	"path"
	"runtime"
	"strconv"
)

var (
	serverInstances int
)

type Visualizations struct {
	Titles []string `json:"titles"`
	Types []string `json:"types"`
}

// > routeData respond with dat required for a particular id
func routeData(w http.ResponseWriter, req *http.Request) {
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
	json.NewEncoder(w).Encode(visualizations[id])
}

// > routeVisualizations respond with all available visualizations
func routeVisualizations(w http.ResponseWriter, req *http.Request) {
	if req.Body == nil {
		http.Error(w, "Please send a request body", 400)
		return
	}

	var viz Visualizations
	for _, v := range visualizations {
		viz.Titles = append(viz.Titles, v.getTitle())
		viz.Types = append(viz.Types, v.getType())
	}

	// Send response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(viz)
}

// > Show runs a local HTTP server on port 8080 to display charts
func Show(blocking bool) {
	_, runContext, _, ok := runtime.Caller(0)
	if !ok {
		panic("No caller information")
	}
	http.Handle("/", http.FileServer(http.Dir(path.Dir(runContext) + "/web")))
	http.HandleFunc("/visualizations", routeVisualizations)
	http.HandleFunc("/data", routeData)
	port := "8080"

	// Lauch HTTP Server in concurrent routine
	if serverInstances == 0 {
		serverInstances += 1
		go http.ListenAndServe("localhost:" + port, nil)
	} else {
		panic("Visualization already in progress")
	}

	fmt.Println("\x1b[34mVisit localhost:8080 to see the graphs\x1b[0m")

	if blocking {
		var input string
		fmt.Println("Press Ctrl-C to stop...")
		fmt.Scanln(&input)
	}
}

func init() {
	serverInstances = 0
}
