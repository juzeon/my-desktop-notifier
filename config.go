package main

import (
	"bytes"
	"github.com/samber/lo"
	"gopkg.in/yaml.v3"
	"os"
	"time"
)

type Cfg struct {
	Schedules []Schedule `yaml:"schedules"`
}
type Schedule struct {
	Week    time.Weekday `yaml:"week"`
	Time    string       `yaml:"time"`
	Content string       `yaml:"content"`
}

func ReadConfig() Cfg {
	v := lo.Must(os.ReadFile("config.yml"))
	var cfg Cfg
	decoder := yaml.NewDecoder(bytes.NewReader(v))
	decoder.KnownFields(true)
	lo.Must0(decoder.Decode(&cfg))
	for _, schedule := range cfg.Schedules {
		MustParseTime(schedule.Time) // validate all time in advance
	}
	return cfg
}
