package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/catsworld/api"
	"github.com/catsworld/botmaid"
	"github.com/catsworld/coc"
	"github.com/catsworld/nyamath/expression"
	"github.com/catsworld/random"
)

type scType struct {
	Status bool
	a, b   string
	Result coc.CheckResult
}

var (
	scMap = make(map[int64]map[int64]*scType)

	wordSCEgg = botmaid.WordSlice{
		botmaid.Word{
			Word:   "",
			Weight: 99,
		},
		botmaid.Word{
			Word:   "你将会遇见一只小阿比(*^▽^*)~\n",
			Weight: 1,
		},
	}

	formatSC = []string{
		"%s，请进行一次意志检定。",
	}
	formatSCResult = []string{
		"%s，汝的理智损失了 %d 点。",
	}
)

func init() {
	bm.AddCommand(botmaid.Command{
		Do:       sc,
		Priority: 5,
		Menu:     "sc",
		Names:    []string{"sc", "sancheck"},
		Help:     " <SANCheck公式> - 进行一次SAN Check",
	})
}

func sc(e *api.Event, b *botmaid.Bot) bool {
	args := botmaid.SplitCommand(e.Message.Text)
	if b.IsCommand(e, "sc", "sancheck") && len(args) > 1 {
		if scMap[e.Place.ID] == nil {
			scMap[e.Place.ID] = map[int64]*scType{}
		}
		scMap[e.Place.ID][e.Sender.ID] = nil
		now := 0
		for i, v := range args[1] {
			if v == '(' {
				now++
			}
			if v == ')' {
				now--
			}
			if v == '/' && now == 0 {
				scMap[e.Place.ID][e.Sender.ID] = &scType{
					Status: true,
					a:      args[1][0:i],
					b:      args[1][i+1 : len(args[1])],
				}
			}
		}
		if scMap[e.Place.ID][e.Sender.ID] == nil {
			return true
		}
		_, err := expression.New(scMap[e.Place.ID][e.Sender.ID].a)
		if err != nil {
			return true
		}
		_, err = expression.New(scMap[e.Place.ID][e.Sender.ID].b)
		if err != nil {
			return true
		}
		message := wordSCEgg.Random() + fmt.Sprintf(random.String(formatSC), e.Sender.NickName)
		send(api.Event{
			Message: &api.Message{
				Text: message,
			},
			Place: e.Place,
		}, b, false)
		return true
	}
	return false
}

func scResp(e *api.Event, b *botmaid.Bot) {
	time.Sleep(time.Second * 2)
	ee := expression.Expression{}
	if scMap[e.Place.ID][e.Sender.ID].Result.BigSuccess() {
		ee, _ = expression.New(strings.Replace(scMap[e.Place.ID][e.Sender.ID].a, "d", "+0*", -1))
	} else if scMap[e.Place.ID][e.Sender.ID].Result.BigFailure() {
		ee, _ = expression.New(strings.Replace(scMap[e.Place.ID][e.Sender.ID].b, "d", "*", -1))
	} else if scMap[e.Place.ID][e.Sender.ID].Result.Success() {
		ee, _ = expression.New(scMap[e.Place.ID][e.Sender.ID].a)
	} else {
		ee, _ = expression.New(scMap[e.Place.ID][e.Sender.ID].b)
	}
	message := fmt.Sprintf(random.String(formatSCResult), e.Sender.NickName, ee.Result())
	scMap[e.Place.ID][e.Sender.ID] = &scType{
		Status: false,
	}
	send(api.Event{
		Message: &api.Message{
			Text: message,
		},
		Place: e.Place,
	}, b, false)
}
