package main

import (
	"fmt"
	"strings"

	"github.com/catsworld/botmaid"
	"github.com/catsworld/botmaid/random"
)

type callType struct {
	Status bool

	List, Resped []int64

	Sponsor *botmaid.User
}

var (
	callMap = map[int64]*callType{}
)

func init() {
	bm.AddCommand(&botmaid.Command{
		SetFlag: func(u *botmaid.Update) {
			u.Message.Flag.BoolP("cancel", "c", false, "")
		},
		Do: func(u *botmaid.Update) bool {
			cancel, _ := u.Message.Flag.GetBool("cancel")
			if cancel {
				if callMap[u.Chat.ID] != nil && callMap[u.Chat.ID].Sponsor.ID == u.User.ID {
					callMap[u.Chat.ID] = nil
					botmaid.Reply(u, "点名已取消。")
				}
				return true
			}

			return false
		},
		Priority: 1000,
	})

	bm.AddCommand(&botmaid.Command{
		Do: func(u *botmaid.Update) bool {
			if callMap[u.Chat.ID] == nil {
				return false
			}

			if botmaid.In(u.User.ID, callMap[u.Chat.ID].List) && !botmaid.In(u.User.ID, callMap[u.Chat.ID].Resped) {
				callMap[u.Chat.ID].Resped = append(callMap[u.Chat.ID].Resped, u.User.ID)
			}

			if len(callMap[u.Chat.ID].List) == len(callMap[u.Chat.ID].Resped) {
				callMap[u.Chat.ID] = nil
				botmaid.Reply(u, "点名完成。")
			}
			return false
		},
		Priority: 1000,
	})

	bm.AddCommand(&botmaid.Command{
		SetFlag: func(u *botmaid.Update) {
			u.Message.Flag.BoolP("status", "s", false, "")
		},
		Do: func(u *botmaid.Update) bool {
			status, _ := u.Message.Flag.GetBool("status")

			if status {
				if callMap[u.Chat.ID] == nil {
					botmaid.Reply(u, random.String([]string{
						"没有正在进行的点名。",
					}))
					return true
				}

				gu := ""
				for _, user := range callMap[u.Chat.ID].List {
					if !botmaid.In(user, callMap[u.Chat.ID].Resped) {
						gu += fmt.Sprintf("%v ", user)
					}
				}
				gu = strings.TrimSpace(gu)

				botmaid.Reply(u, fmt.Sprintf("未到名单：%v", gu))
				return true
			}

			return false
		},
		Menu:  "call",
		Names: []string{"call"},
		Help:  " status - 查看当前点名情况",
	})
	bm.AddCommand(&botmaid.Command{
		Do: func(u *botmaid.Update) bool {
			callMap[u.Chat.ID] = &callType{
				Status: true,
				List:   []int64{},
				Resped: []int64{},
			}

			for i := 1; i < len(u.Message.Args); i++ {
				callMap[u.Chat.ID].List = append(callMap[u.Chat.ID].List, u.Message.Args[i])
			}

			botmaid.Reply(u, random.String([]string{
				"呵呵呵，看来调查员的召集开始了。",
				"你们又要演出新的戏剧了吗？阿比也来看吧w",
			}))
			return true
		},
		Menu:       "call",
		MenuText:   "点名",
		Names:      []string{"call"},
		ArgsMinLen: 2,
		Help:       " <@其他人> - 进行一次点名",
	})
}
