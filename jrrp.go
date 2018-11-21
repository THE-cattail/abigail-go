package main

import (
	"time"

	"github.com/catsworld/api"
	"github.com/catsworld/botmaid"
)

func init() {
	bm.AddCommand(botmaid.Command{
		Do:       switchJrrp,
		Priority: 5,
		Master:   true,
	})
	bm.AddTimer(botmaid.Timer{
		Do:        jrrp,
		Time:      time.Date(2018, 10, 9, 0, 0, 5, 0, loc),
		Frequency: "daily",
	})
}

func switchJrrp(e *api.Event, b *botmaid.Bot) bool {
	args := botmaid.SplitCommand(e.Message.Text)
	if b.IsCommand(e, "jrrp") && len(args) == 1 {
		bm.SwitchBroadcast("jrrp", e.Place, b)
		return true
	}

	return false
}

func jrrp() {
	bm.Broadcast("jrrp", &api.Message{
		Text: ".jrrp",
	})
}
