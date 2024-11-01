package main

import (
	_ "embed"
	"github.com/go-co-op/gocron/v2"
	"github.com/samber/lo"
	"io"
	"log/slog"
	"os"
	"strconv"
	"time"
)

//go:embed assets/Alarm.png
var alarmAsset []byte // from https://github.com/Semporia/Hand-Painted-icon

func main() {
	logFile := lo.Must(os.OpenFile("log.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644))
	logWriter := io.MultiWriter(os.Stderr, logFile)
	logHandler := slog.New(slog.NewTextHandler(logWriter, &slog.HandlerOptions{Level: slog.LevelInfo}))
	slog.SetDefault(logHandler)
	config := ReadConfig()
	scheduler := lo.Must(gocron.NewScheduler())
	ScheduleToday(config, scheduler)
	lo.Must(scheduler.NewJob(
		gocron.DailyJob(1, gocron.NewAtTimes(gocron.NewAtTime(0, 0, 1))),
		gocron.NewTask(func() {
			slog.Info("Scheduling for a new day", "time", time.Now())
			ScheduleToday(config, scheduler)
		}),
	))
	scheduler.Start()
	select {}
}
func ScheduleToday(config Cfg, scheduler gocron.Scheduler) []gocron.Job {
	timeNow := time.Now()
	var jobs []gocron.Job
	for _, schedule := range config.Schedules {
		schedule := schedule
		if schedule.Week != timeNow.Weekday() {
			continue
		}
		slog.Info("Scheduling", "v", schedule)
		scheduleTimeParsed := MustParseTime(schedule.Time)
		startTime := time.Date(timeNow.Year(), timeNow.Month(), timeNow.Day(),
			scheduleTimeParsed.Hour(), scheduleTimeParsed.Minute(), 0, 0, time.Local)
		j, err := scheduler.NewJob(gocron.OneTimeJob(gocron.OneTimeJobStartDateTime(startTime)),
			gocron.NewTask(func() {
				slog.Info("Schedule reached", "v", schedule)
				MustSendNotification(schedule.Content)
			}),
		)
		if err != nil {
			slog.Warn("Schedule failed", "err", err, "schedule", schedule)
			continue
		}
		jobs = append(jobs, j)
	}
	MustSendNotification("Scheduled " + strconv.Itoa(len(jobs)) + " notifications today")
	return jobs
}
