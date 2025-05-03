package main

import (
	"log"
	"net/http"

	"timeslot-optimizer/controller"
)

func main() {
	http.HandleFunc("/api/optimize-schedule", controller.OptimizeScheduleHandler)
	log.Println("✅ Server running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
