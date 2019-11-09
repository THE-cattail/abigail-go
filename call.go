package main

import (
	"fmt"
	"log"

	"github.com/catsworld/botmaid"
	"github.com/catsworld/botmaid/random"
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
			cancel, _ := bm.Flags["call"].GetBool("cancel")
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
		Do: func(u *botmaid.Update, f *pflag.FlagSet) bool {
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
		Do: func(u *botmaid.Update, f *pflag.FlagSet) bool {
			status, _ := f.GetBool("status")

			if status {
				if callMap[u.Chat.ID] == nil {
					botmaid.Reply(u, random.String([]string{
						"没有正在进行的点名。",
					}))
					return true
				}

				l := ""
				for i := range callMap[u.Chat.ID].List {
					if !botmaid.In(callMap[u.Chat.ID].List[i], callMap[u.Chat.ID].Resped) {
						if l != "" {
							l += " "
						}
						l += callMap[u.Chat.ID].At[i]
					}
				}

				botmaid.Reply(u, fmt.Sprintf("未到名单：%v", l))
				return true
			}

			if len(f.Args()) > 1 {
				callMap[u.Chat.ID] = &callType{
					Sponsor: u.User,
				}

				for i := 1; i < len(f.Args()); i++ {
					id, err := bm.ParseUserID(u, f.Args()[i])
					if err != nil {
						log.Println(err)
						botmaid.Reply(u, fmt.Sprintf(random.String(bm.Words["invalidUser"]), botmaid.At(u.User), f.Args()[i]))
						return true
					}

					callMap[u.Chat.ID].List = append(callMap[u.Chat.ID].List, id)
					callMap[u.Chat.ID].At = append(callMap[u.Chat.ID].At, f.Args()[i])
				}

				botmaid.Reply(u, "开始点名。")
				return true
			}

			return false
		},
		Help: &botmaid.Help{
			Menu:  "call",
			Help:  "点名功能",
			Names: []string{"call"},
			Full: `使用方法：call @用户...

%v`,
			SetFlag: func(f *pflag.FlagSet) {
				f.BoolP("cancel", "c", false, "取消当前点名")
				f.BoolP("status", "s", false, "查看当前点名情况")
			},
		},
	})
}
