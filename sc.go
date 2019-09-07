package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/catsworld/abigail/coc"
	"github.com/catsworld/botmaid"
	"github.com/catsworld/nyamath"
	"github.com/catsworld/random"
)

type scType struct {
	Status bool
	a, b   string
	Result coc.CheckResult
}

var (
	scMap = make(map[int64]map[int64]*scType)
)

func init() {
	bm.AddCommand(botmaid.Command{
		Do: func(u *botmaid.Update, b *botmaid.Bot) bool {
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
				return true
			}
			_, err := nyamath.New(scMap[u.Chat.ID][u.User.ID].a)
			if err != nil {
				return true
			}
			_, err = nyamath.New(scMap[u.Chat.ID][u.User.ID].b)
			if err != nil {
				return true
			}
			message := botmaid.WordSlice{
				botmaid.Word{
					Word:   "",
					Weight: 99,
				},
				botmaid.Word{
					Word:   "你将会遇见一只小阿比(*^▽^*)~\n",
					Weight: 1,
				},
			}.Random() + fmt.Sprintf(random.String([]string{
				"%s，请进行一次意志检定。",
			}), u.User.NickName)
			send(b, botmaid.Update{
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

func scResp(u *botmaid.Update, b *botmaid.Bot) {
	time.Sleep(time.Second * 2)
	ee := nyamath.Expression{}
	if scMap[u.Chat.ID][u.User.ID].Result.BigSuccess() {
		ee, _ = nyamath.New(strings.Replace(scMap[u.Chat.ID][u.User.ID].a, "d", "+0*", -1))
	} else if scMap[u.Chat.ID][u.User.ID].Result.BigFailure() {
		ee, _ = nyamath.New(strings.Replace(scMap[u.Chat.ID][u.User.ID].b, "d", "*", -1))
	} else if scMap[u.Chat.ID][u.User.ID].Result.Success() {
		ee, _ = nyamath.New(scMap[u.Chat.ID][u.User.ID].a)
	} else {
		ee, _ = nyamath.New(scMap[u.Chat.ID][u.User.ID].b)
	}
	message := fmt.Sprintf(random.String([]string{
		"%s，汝的理智损失了 %d 点。",
	}), u.User.NickName, ee.Result())
	scMap[u.Chat.ID][u.User.ID] = &scType{
		Status: false,
	}
	send(b, botmaid.Update{
		Message: &botmaid.Message{
			Text: message,
		},
		Chat: u.Chat,
	}, false, u)
}
