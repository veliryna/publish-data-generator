package handlers

import (
	"data-generator/reports"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
	"google.golang.org/grpc"
)

var upgrader = websocket.Upgrader{}
var timeConn *websocket.Conn
var eventConn *websocket.Conn
var broadcastTimeBasedReports = make(chan *reports.TimeBasedReport, 10000)
var broadcastEventBasedReports = make(chan *reports.EventBasedReport, 10000)

func CreatePublishingServer() {
	http.HandleFunc("/time_based_reports", handleWSTime)
	http.HandleFunc("/event_based_reports", handleWSEvents)
	log.Println("WebSocket server listening on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

// Handle WebSocket for time-based reports
func handleWSTime(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Print("WebSocket upgrade error (time):", err)
		return
	}
	timeConn = conn
	go timeWritePump()
}

// Handle WebSocket for event-based reports
func handleWSEvents(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Print("WebSocket upgrade error (event):", err)
		return
	}
	eventConn = conn
	go eventWritePump()
}

func timeWritePump() {
	for report := range broadcastTimeBasedReports {
		if timeConn == nil {
			continue
		}
		if err := timeConn.WriteJSON(report); err != nil {
			log.Printf("WebSocket write error (time): %v", err)
			timeConn.Close()
			timeConn = nil
			return
		}
	}
}

func eventWritePump() {
	for report := range broadcastEventBasedReports {
		if eventConn == nil {
			continue
		}
		if err := eventConn.WriteJSON(report); err != nil {
			log.Printf("WebSocket write error (event): %v", err)
			eventConn.Close()
			eventConn = nil
			return
		}
	}
}

func SendTimeBasedReports(timeStream grpc.ServerStreamingClient[reports.TimeBasedReport]) {
	for {
		report, err := timeStream.Recv()
		if err != nil {
			log.Printf("Time-based stream closed: %v", err)
			return
		}
		broadcastTimeBasedReports <- report
	}
}

func SendEventBasedReports(eventStream grpc.ServerStreamingClient[reports.EventBasedReport]) {
	for {
		report, err := eventStream.Recv()
		if err != nil {
			log.Printf("Event-based stream closed: %v", err)
			return
		}
		broadcastEventBasedReports <- report
	}
}
