package main

import (
	"fmt"

	"github.com/the-cattail/botmaid"
	"github.com/spf13/pflag"
)

type callType struct {
	List, Resped []int64
	At           []string

	Sponsor *botmaid.User
}

var (
	callMap = map[int64]*callType{}
)

func init() {
	bm.AddCommand(&botmaid.Command{
		Do: func(u *botmaid.Update, f *pflag.FlagSet) bool {
			cancel, _ := u.Message.Flags["call"].GetBool("cancel")
			if cancel {
				if callMap[u.Chat.ID] != nil && callMap[u.Chat.ID].Sponsor.ID == u.User.ID {
					callMap[u.Chat.ID] = nil
					bm.Reply(u, "点名已取消。")
				}
				return true
			}

			return false
		},
		Priority: 1000,
	})

	bm.AddCommand(&botmaid.Command{
		Do: func(u *botmaid.Update, f *pflag.FlagSet) bool {
			if callMap[u.Chat.ID] == nil {
				return false
			}

			if botmaid.Contains(callMap[u.Chat.ID].List, u.User.ID) && !botmaid.Contains(callMap[u.Chat.ID].Resped, u.User.ID) {
				callMap[u.Chat.ID].Resped = append(callMap[u.Chat.ID].Resped, u.User.ID)
			}

			if len(callMap[u.Chat.ID].List) == len(callMap[u.Chat.ID].Resped) {
				callMap[u.Chat.ID] = nil
				bm.Reply(u, "点名完成。")
			}

			return false
		},
		Priority: 1000,
	})

	bm.AddCommand(&botmaid.Command{
		Do: func(u *botmaid.Update, f *pflag.FlagSet) bool {
			status, _ := f.GetBool("status")

			if status {
				if callMap[u.Chat.ID] == nil {
					bm.Reply(u, "没有正在进行的点名。")
					return true
				}

				l := ""
				for i := range callMap[u.Chat.ID].List {
					if !botmaid.Contains(callMap[u.Chat.ID].Resped, callMap[u.Chat.ID].List[i]) {
						if l != "" {
							l += " "
						}
						l += callMap[u.Chat.ID].At[i]
					}
				}

				bm.Reply(u, fmt.Sprintf("未到名单：%v", l))
				return true
			}

			if len(f.Args()) > 1 {
				callMap[u.Chat.ID] = &callType{
					Sponsor: u.User,
				}

				for i := 1; i < len(f.Args()); i++ {
					id, err := (*u.Bot.API).ParseUserID(u, f.Args()[i])
					if err != nil {
						bm.Reply(u, fmt.Sprintf(bm.Words["invalidUser"], bm.At(u.User), f.Args()[i]))
						callMap[u.Chat.ID] = nil
						return true
					}

					callMap[u.Chat.ID].List = append(callMap[u.Chat.ID].List, id)
					callMap[u.Chat.ID].At = append(callMap[u.Chat.ID].At, f.Args()[i])
				}

				bm.Reply(u, "开始点名。")
				return true
			}

			return false
		},
		Help: &botmaid.Help{
			Menu:  "call",
			Help:  "点名功能",
			Names: []string{"call"},
			Usage: "使用方法：call @用户...",
			SetFlag: func(f *pflag.FlagSet) {
				f.BoolP("cancel", "c", false, "取消当前点名")
				f.BoolP("status", "s", false, "查看当前点名情况")
			},
		},
	})
}
