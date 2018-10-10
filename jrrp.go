package main

import (
	"time"

	"github.com/catsworld/api"
	"github.com/catsworld/botmaid"
)

func init() {
	botmaid.AddTimer(&timers, jrrp, time.Date(2018, 10, 9, 0, 0, 5, 0, loc), "daily")
}

func jrrp() {
	for _, v := range botMaid.Bots {
		if v.Self.ID == 1261413197 {
			v.API.Push(&api.Event{
				Message: &api.Message{
					Text: ".jrrp",
				},
				Place: &api.Place{
					Type: "group",
					ID:   738979059,
				},
			})
			v.API.Push(&api.Event{
				Message: &api.Message{
					Text: ".jrrp",
				},
				Place: &api.Place{
					Type: "group",
					ID:   773745852,
				},
			})
			return
		}
	}
}
