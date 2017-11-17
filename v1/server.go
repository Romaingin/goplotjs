package goplotjs

import (
	"encoding/json"
	"fmt"
	"net/http"
	"path"
	"runtime"
)

var (
	serverInstances int
	visualizations []VisualizeData
)

func routeData(w http.ResponseWriter, req *http.Request) {
	if req.Body == nil {
		http.Error(w, "Please send a request body", 400)
		return
	}

	fmt.Print(req.URL.Query().Get("id"))

	// Send response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(visualizations[0])
}

func Show(blocking bool) {
	_, runContext, _, ok := runtime.Caller(0)
	if !ok {
		panic("No caller information")
	}
	http.Handle("/", http.FileServer(http.Dir(path.Dir(runContext) + "/web")))
	http.HandleFunc("/data", routeData)
	port := "8080"

	// Lauch HTTP Server in concurrent routine
	if serverInstances == 0 {
		serverInstances += 1
		go http.ListenAndServe("localhost:" + port, nil)
	} else {
		panic("Visualization already in progress")
	}

	if blocking {
		var input string
		fmt.Println("Press Ctrl-C to stop...")
		fmt.Scanln(&input)
	}
}

func init() {
	fmt.Println("Using goplotjs...")
	serverInstances = 0
}
