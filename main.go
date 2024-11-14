package main

import (
	_ "embed"
	"github.com/samber/lo"
	"io"
	"log/slog"
	"net/http"
	"os"
	"strconv"
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
	go func() {
		http.ListenAndServe(":"+strconv.Itoa(lo.Ternary(config.Port == 0, 7888, config.Port)),
			http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
				if request.Method != http.MethodPost {
					writer.WriteHeader(400)
					return
				}
				cfg, err := ReadConfig()
				if err != nil {
					writer.WriteHeader(400)
					writer.Write([]byte(err.Error()))
					return
				}
				config = cfg
				scheduler.UpdateSchedules(config.Schedules)
			}))
	}()
	scheduler.Run()
}
