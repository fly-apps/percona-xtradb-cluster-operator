package main

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os/exec"
)

func main() {
	http.HandleFunc("/probe", probeHandler)

	log.Println("Starting liveness/readiness probes HTTP server on port 8090")
	if err := http.ListenAndServe(":8090", nil); err != nil {
		log.Fatalf("Failed to start HTTP server: %v", err)
	}
}

func probeHandler(w http.ResponseWriter, r *http.Request) {
	scriptPath := r.URL.Query().Get("script")
	if scriptPath == "" {
		http.Error(w, "Missing 'script' query parameter", http.StatusBadRequest)
		return
	}

	// Decode the URL-encoded script path
	decodedPath, err := url.QueryUnescape(scriptPath)
	if err != nil {
		http.Error(w, "Failed to decode 'script' query parameter", http.StatusBadRequest)
		return
	}

	cmd := exec.Command("/bin/bash", decodedPath)

	// Run the command and capture the exit code
	if err := cmd.Run(); err != nil {
		if exitError, ok := err.(*exec.ExitError); ok {
			exitCode := exitError.ExitCode()
			http.Error(w, fmt.Sprintf("Script failed with exit code %d", exitCode), http.StatusInternalServerError)
			return
		}
		// Handle other errors (e.g., command not found)
		http.Error(w, fmt.Sprintf("Failed to run script: %v", err), http.StatusInternalServerError)
		return
	}

	// If the script ran successfully
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Probe check passed"))
}
