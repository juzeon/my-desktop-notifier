package main

import (
	"github.com/samber/lo"
	"slices"
	"strconv"
	"time"
)

type Scheduler struct {
	Schedules               []*Schedule
	RemainingSchedulesToday []*Schedule
	DateToday               string
}

func NewScheduler(schedules []*Schedule) *Scheduler {
	return &Scheduler{
		Schedules:               schedules,
		RemainingSchedulesToday: nil,
		DateToday:               "0000-00-00",
	}
}
func (o *Scheduler) UpdateSchedules(schedules []*Schedule) {
	o.Schedules = schedules
	o.scheduleForToday()
}
func (o *Scheduler) scheduleForToday() {
	timeNow := time.Now()
	week := timeNow.Weekday()
	o.RemainingSchedulesToday = nil
	for _, schedule := range o.Schedules {
		if schedule.Week == week && schedule.MustGetTime().After(timeNow) {
			o.RemainingSchedulesToday = append(o.RemainingSchedulesToday, schedule)
		}
	}
	o.DateToday = timeNow.Format("2006-01-02")
	SendNotification("Scheduled " + strconv.Itoa(len(o.RemainingSchedulesToday)) + " tasks today.")
}
func (o *Scheduler) Run() {
	for {
		timeNow := time.Now()
		if time.Now().Format("2006-01-02") != o.DateToday {
			o.scheduleForToday()
		}
		var removingSchedules []*Schedule
		for _, schedule := range o.RemainingSchedulesToday {
			if schedule.MustGetTime().Before(timeNow) {
				SendNotification(schedule.Content)
				removingSchedules = append(removingSchedules, schedule)
			}
		}
		if len(removingSchedules) != 0 {
			o.RemainingSchedulesToday = lo.Filter(o.RemainingSchedulesToday, func(item *Schedule, index int) bool {
				return !slices.Contains(removingSchedules, item)
			})
		}
		time.Sleep(30 * time.Second)
	}
}
