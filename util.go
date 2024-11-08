package main

import (
	"github.com/gen2brain/beeep"
	"log/slog"
	"os"
	"path/filepath"
)

func SendNotification(content string) {
	icon := filepath.Join(os.TempDir(), "Alarm.png")
	if _, err := os.Stat(icon); err != nil {
		err = os.WriteFile(icon, alarmAsset, 0644)
		if err != nil {
			icon = ""
		}
	}
	slog.Info("SendNotification", "content", content)
	err := beeep.Notify("My Desktop Notifier", content, icon)
	if err != nil {
		slog.Error("Error sending notification", "err", err)
	}
}
