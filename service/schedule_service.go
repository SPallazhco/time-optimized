package service

import (
	"time"

	"timeslot-optimizer/model"
)

func OptimizeSchedule(tasks []model.Task) []model.ScheduleItem {
	startTime := time.Now().Round(30 * time.Minute)

	var schedule []model.ScheduleItem
	current := startTime

	for _, task := range tasks {
		end := current.Add(time.Duration(task.Duration) * time.Minute)

		schedule = append(schedule, model.ScheduleItem{
			Start: current.Format("15:04"),
			End:   end.Format("15:04"),
			Task:  task.Name,
		})

		current = end
	}

	return schedule
}
