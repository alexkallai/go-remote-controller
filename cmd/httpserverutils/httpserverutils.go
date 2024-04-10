package httpserverutils

import (
	"encoding/json"
	"fmt"
	"go-remote-controller/cmd/inpututils"
	"log"
	"net/http"
	"time"
)

type Event struct {
	EventType            string    `json:"eventType"`
	XCoordinateArray     []float64 `json:"xCoordinateArray"`
	YCoordinateArray     []float64 `json:"yCoordinateArray"`
	CoordinateIdentifier []int32   `json:"coordinateIdentifier"`
}

func SetupHttpServerEndpoints(PORT *int) {
	// Set up listener endpoints
	http.HandleFunc("/", handleRootEndpoint)
	http.HandleFunc("/api", handleApiEndpoint)
	http.HandleFunc("/favicon.ico", doNothing)

	// Set up http server
	port_string := fmt.Sprintf(":%d", *PORT)
	err := http.ListenAndServe(port_string, nil)
	if err != nil {
		log.Fatalf("Failed to serve - %s", err)
	}
}

func handleRootEndpoint(w http.ResponseWriter, r *http.Request) {
	log.Println("Main page served!")
	fmt.Fprint(w, string(htmlFileTemplate))

}

func handleApiEndpoint(w http.ResponseWriter, r *http.Request) {
	start := time.Now()

	var e Event
	err := json.NewDecoder(r.Body).Decode(&e)
	if err != nil {
		log.Printf("Error reading body: %v", err)
	}
	if e.EventType == "touchmove" || e.EventType == "touchstart" || e.EventType == "touchend" || e.EventType == "touchcancel" {
		log.Printf("Handling touch event for %v", e)
		HandleTouchEvent(e)
	} else if e.EventType == "rightclick" || e.EventType == "midclick" || e.EventType == "leftclick" {
		log.Printf("Handling single click event: %v", e.EventType)
		inpututils.SendMouseInput(inpututils.EventNameToMouseInputMap[e.EventType])
	} else {
		log.Printf("Handling single keyboard input event: %v", e.EventType)
		inpututils.SendKeys(inpututils.EventNameToKbdInputMap[e.EventType])
	}

	elapsed := time.Since(start)
	log.Println("API request served in", elapsed)
}

func doNothing(w http.ResponseWriter, r *http.Request) {}
