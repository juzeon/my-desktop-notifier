package main

import (
	"github.com/gen2brain/beeep"
	"github.com/samber/lo"
	"os"
	"path/filepath"
	"time"
)

func MustParseTime(t string) time.Time {
	return lo.Must(time.Parse("15:04", t))
}
func MustSendNotification(content string) {
	icon := filepath.Join(os.TempDir(), "Alarm.png")
	if _, err := os.Stat(icon); err != nil {
		err = os.WriteFile(icon, alarmAsset, 0644)
		if err != nil {
			icon = ""
		}
	}
	lo.Must0(beeep.Notify("My Desktop Notifier", content, icon))
}
