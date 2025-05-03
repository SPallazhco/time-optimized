package controller

import (
	"encoding/json"
	"net/http"

	"timeslot-optimizer/model"
	"timeslot-optimizer/service"
)

func OptimizeScheduleHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Only POST method is allowed", http.StatusMethodNotAllowed)
		return
	}

	var req struct {
		Tasks []model.Task `json:"tasks"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	schedule := service.OptimizeSchedule(req.Tasks)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"schedule": schedule,
	})
}
