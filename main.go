package main

import (
	_ "embed"
	"github.com/samber/lo"
	"io"
	"log/slog"
	"os"
)

//go:embed assets/Alarm.png
var alarmAsset []byte // from https://github.com/Semporia/Hand-Painted-icon

func main() {
	logFile := lo.Must(os.OpenFile("log.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644))
	logWriter := io.MultiWriter(os.Stderr, logFile)
	logHandler := slog.New(slog.NewTextHandler(logWriter, &slog.HandlerOptions{Level: slog.LevelInfo}))
	slog.SetDefault(logHandler)
	config, err := ReadConfig()
	if err != nil {
		SendNotification(err.Error())
		os.Exit(1)
	}
	scheduler := NewScheduler(config.Schedules)
	scheduler.Run()
}
