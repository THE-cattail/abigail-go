package main

import (
	"fmt"

	"github.com/catsworld/api"
	"github.com/catsworld/botmaid"
	"github.com/catsworld/random"
	"github.com/catsworld/slices"
)

type callType struct {
	Status bool
	List   map[string]bool
	Resped map[string]bool
	Total  int
	Get    int
}

var (
	callMap       = make(map[int64]*callType)
	wordCallStart = []string{
		"呵呵呵，看来调查员的召集开始了。",
	}
	wordCallNotStart = []string{
		"好像有什么很热闹的样子呢……我也好想被叫去参加呢。",
	}
	wordCallNobody = []string{
		"还没有谁在哦，果然大家都是鸽子吧（生气）。",
	}
	wordCallComplete = []string{
		"调查员都已经聚集好了哦。呵呵，是不是又有什么事情要发生了。",
	}
	formatCallList = []string{
		"已经有%d名伙伴出现啦~\n鸽子名单：\n",
	}
	formatGule = []string{
		"看啊看啊！%s这家伙咕了哦！",
	}
)

func init() {
	botmaid.AddCommand(&commands, gule, 10)
	botmaid.AddCommand(&commands, callResp, 10)
	botmaid.AddCommand(&commands, callStatus, 5)
	botmaid.AddCommand(&commands, callGugugu, 5)
	botmaid.AddCommand(&commands, call, 5)
}

func call(e *api.Event, b *botmaid.Bot) bool {
	if b.IsCommand(e, "call") {
		args := botmaid.SplitCommand(e.Message.Text)
		callMap[e.Place.ID] = &callType{
			Status: true,
			Total:  len(args) - 1,
			Get:    0,
			List:   make(map[string]bool),
			Resped: make(map[string]bool),
		}
		for i := 1; i < len(args); i++ {
			callMap[e.Place.ID].List[args[i]] = true
		}
		send(&api.Event{
			Message: &api.Message{
				Text: random.String(wordCallStart),
			},
			Place: e.Place,
		}, b, false)
		return true
	}
	return false
}

func callGugugu(e *api.Event, b *botmaid.Bot) bool {
	args := botmaid.SplitCommand(e.Message.Text)
	if b.IsCommand(e, "call") && len(args) > 2 && slices.In(args[1], "-gugugu") {
		theGugugu := dbAbiGugugu{}
		err := db.QueryRow("SELECT * FROM abi_gugugu WHERE chat_id = $1 AND name = $2", e.Place.ID, args[2]).Scan(&theGugugu.ID, &theGugugu.PlaceID, &theGugugu.Name, &theGugugu.Members, &theGugugu.At, &theGugugu.Status)
		if err != nil || theGugugu.At == "" {
			return true
		}
		ee := e
		ee.Message.Text = "/call " + theGugugu.At
		call(ee, b)
		send(&api.Event{
			Message: &api.Message{
				Text: theGugugu.At,
			},
			Place: e.Place,
		}, b, false)
		return true
	}
	return false
}

func callStatus(e *api.Event, b *botmaid.Bot) bool {
	args := botmaid.SplitCommand(e.Message.Text)
	if b.IsCommand(e, "call") && len(args) > 1 && slices.In(args[1], "-status", "-s") {
		if _, ok := callMap[e.Place.ID]; !ok || !callMap[e.Place.ID].Status {
			send(&api.Event{
				Message: &api.Message{
					Text: random.String(wordCallNotStart),
				},
				Place: e.Place,
			}, b, false)
			return true
		}
		if callMap[e.Place.ID].Get == 0 {
			send(&api.Event{
				Message: &api.Message{
					Text: random.String(wordCallNobody),
				},
				Place: e.Place,
			}, b, false)
			return true
		}
		message := fmt.Sprintf(random.String(formatCallList), callMap[e.Place.ID].Get)
		for key := range callMap[e.Place.ID].List {
			if !callMap[e.Place.ID].Resped[key] {
				message += key + " "
			}
		}
		send(&api.Event{
			Message: &api.Message{
				Text: message[:len(message)-1],
			},
			Place: e.Place,
		}, b, false)
		return true
	}
	return false
}

func gule(e *api.Event, b *botmaid.Bot) bool {
	if b.IsCommand(e, "咕咕", "咕了") {
		if _, ok := callMap[e.Place.ID]; ok && callMap[e.Place.ID].Status && callMap[e.Place.ID].List[b.At(e.Sender)] {
			callMap[e.Place.ID].Status = false
			callMap[e.Place.ID].List = make(map[string]bool)
			send(&api.Event{
				Message: &api.Message{
					Text: fmt.Sprintf(random.String(formatGule), e.Sender.NickName),
				},
				Place: e.Place,
			}, b, false)
		}
		return false
	}
	return false
}

func callResp(e *api.Event, b *botmaid.Bot) bool {
	if _, ok := callMap[e.Place.ID]; !ok || !callMap[e.Place.ID].Status {
		return false
	}
	if callMap[e.Place.ID].List[b.At(e.Sender)] && !callMap[e.Place.ID].Resped[b.At(e.Sender)] {
		callMap[e.Place.ID].Resped[b.At(e.Sender)] = true
		callMap[e.Place.ID].Get++
	}
	if callMap[e.Place.ID].Get == callMap[e.Place.ID].Total {
		callMap[e.Place.ID].Status = false
		send(&api.Event{
			Message: &api.Message{
				Text: random.String(wordCallComplete),
			},
			Place: e.Place,
		}, b, false)
	}
	return false
}
