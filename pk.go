package main

import (
	"fmt"

	"github.com/catsworld/api"
	"github.com/catsworld/botmaid"
	"github.com/catsworld/coc"
	"github.com/catsworld/random"
)

type pkRollResult struct {
	User   *api.User
	Result coc.CheckResult
}

type pkType struct {
	Status  bool
	Results []pkRollResult
}

var (
	pkMap       = make(map[int64]*pkType)
	wordPKStart = []string{
		"开始对抗，请一方roll点。",
	}
	wordPKNext = []string{
		"请另一方roll点。",
	}
	wordPKDraw = []string{
		"对抗结果：平局！",
	}
	formatPKWin = []string{
		"对抗结果：%s胜利！",
	}
)

func init() {
	botmaid.AddCommand(&commands, pk, 5)
}

func pk(e *api.Event, b *botmaid.Bot) bool {
	if b.IsCommand(e, "pk") {
		pkMap[e.Place.ID] = &pkType{
			Status: true,
		}
		send(&api.Event{
			Message: &api.Message{
				Text: random.String(wordPKStart),
			},
			Place: e.Place,
		}, b, false)
		return true
	}
	return false
}

func pkResp(e *api.Event, b *botmaid.Bot) {
	if _, ok := pkMap[e.Place.ID]; !ok || !pkMap[e.Place.ID].Status {
		return
	}
	if len(pkMap[e.Place.ID].Results) == 1 {
		send(&api.Event{
			Message: &api.Message{
				Text: random.String(wordPKNext),
			},
			Place: e.Place,
		}, b, false)
		return
	}
	pkResult := coc.PK(pkMap[e.Place.ID].Results[0].Result, pkMap[e.Place.ID].Results[1].Result)
	message := ""
	if pkResult == coc.PKDraw {
		message = random.String(wordPKDraw)
	} else {
		victor := pkMap[e.Place.ID].Results[0].User.NickName
		if pkResult == coc.PKBWin {
			victor = pkMap[e.Place.ID].Results[1].User.NickName
		}
		message = fmt.Sprintf(random.String(formatPKWin), victor)
	}
	pkMap[e.Place.ID] = &pkType{
		Status: false,
	}
	send(&api.Event{
		Message: &api.Message{
			Text: message,
		},
		Place: e.Place,
	}, b, false)
}
