package model

type Task struct {
	Name     string `json:"name"`
	Duration int    `json:"duration"`
}
type ScheduleItem struct {
	Start string `json:"start"`
	End   string `json:"end"`
	Task  string `json:"task"`
}
