package main

import (
	"bytes"
	"github.com/samber/lo"
	"gopkg.in/yaml.v3"
	"os"
	"time"
)

type Cfg struct {
	Schedules []*Schedule `yaml:"schedules"`
}
type Schedule struct {
	Week    time.Weekday `yaml:"week"`
	Time    string       `yaml:"time"`
	Content string       `yaml:"content"`
}

func (o *Schedule) GetTime() (time.Time, error) {
	t, err := time.Parse("15:04", o.Time)
	if err != nil {
		return time.Time{}, err
	}
	timeNow := time.Now()
	return time.Date(timeNow.Year(), timeNow.Month(), timeNow.Day(),
		t.Hour(), t.Minute(), 0, 0, time.Local), nil
}
func (o *Schedule) MustGetTime() time.Time {
	return lo.Must(o.GetTime())
}

func ReadConfig() (Cfg, error) {
	v := lo.Must(os.ReadFile("config.yml"))
	var cfg Cfg
	decoder := yaml.NewDecoder(bytes.NewReader(v))
	decoder.KnownFields(true)
	lo.Must0(decoder.Decode(&cfg))
	for _, schedule := range cfg.Schedules {
		_, err := schedule.GetTime() // validate all time in advance
		if err != nil {
			return cfg, err
		}
	}
	return cfg, nil
}
