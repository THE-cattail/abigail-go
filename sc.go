package main

import (
	"fmt"
	"time"

	"github.com/catsworld/abigail/coc"
	"github.com/catsworld/abigail/nyamath"
	"github.com/catsworld/botmaid"
	"github.com/catsworld/botmaid/random"
)

type scType struct {
	Status bool
	a, b   string
	Result coc.CheckResult
}

var (
	scMap = map[int64]map[int64]*scType{}
)

func init() {
	bm.AddCommand(&botmaid.Command{
		Do: func(u *botmaid.Update) bool {
			if scMap[u.Chat.ID] == nil {
				scMap[u.Chat.ID] = map[int64]*scType{}
			}
			scMap[u.Chat.ID][u.User.ID] = nil
			now := 0
			for i, v := range u.Message.Args[1] {
				if v == '(' {
					now++
				}
				if v == ')' {
					now--
				}
				if v == '/' && now == 0 {
					scMap[u.Chat.ID][u.User.ID] = &scType{
						Status: true,
						a:      u.Message.Args[1][0:i],
						b:      u.Message.Args[1][i+1 : len(u.Message.Args[1])],
					}
				}
			}
			if scMap[u.Chat.ID][u.User.ID] == nil {
				return false
			}
			_, err := nyamath.New(scMap[u.Chat.ID][u.User.ID].a)
			if err != nil {
				return false
			}
			_, err = nyamath.New(scMap[u.Chat.ID][u.User.ID].b)
			if err != nil {
				return false
			}
			message := random.WordSlice{
				random.Word{
					Word:   "",
					Weight: 99,
				},
				random.Word{
					Word:   "你将会遇见一只小阿比(*^▽^*)~\n",
					Weight: 1,
				},
			}.Random() + fmt.Sprintf(random.String([]string{
				"%s，请进行一次意志检定。",
			}), u.User.NickName)
			send(&botmaid.Update{
				Message: &botmaid.Message{
					Text: message,
				},
				Chat: u.Chat,
			}, false, u)
			return true
		},
		Menu:       "sc",
		MenuText:   "SAN check",
		Names:      []string{"sc", "sancheck"},
		ArgsMinLen: 2,
		ArgsMaxLen: 2,
		Help:       " <SAN check公式> - 进行一次SAN check",
	})
}

func scResp(u *botmaid.Update) {
	time.Sleep(time.Second * 2)
	ea, _ := nyamath.New(scMap[u.Chat.ID][u.User.ID].a)
	eb, _ := nyamath.New(scMap[u.Chat.ID][u.User.ID].b)
	res := 0
	if scMap[u.Chat.ID][u.User.ID].Result.Great == coc.GreatSucc {
		res = ea.Result.Min
	} else if scMap[u.Chat.ID][u.User.ID].Result.Great == coc.GreatFail {
		res = eb.Result.Max
	} else if scMap[u.Chat.ID][u.User.ID].Result.Succ == coc.Succ {
		res = ea.Result.Value
	} else {
		res = eb.Result.Value
	}
	message := fmt.Sprintf(random.String([]string{
		"%s，汝的理智损失了 %d 点。",
	}), u.User.NickName, res)
	scMap[u.Chat.ID][u.User.ID] = &scType{
		Status: false,
	}
	send(&botmaid.Update{
		Message: &botmaid.Message{
			Text: message,
		},
		Chat: u.Chat,
	}, false, u)
}
