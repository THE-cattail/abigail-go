package main

import (
	"fmt"

	"github.com/catsworld/botmaid"
	"github.com/spf13/pflag"
)

func init() {
	bm.AddCommand(&botmaid.Command{
		Do: func(u *botmaid.Update, f *pflag.FlagSet) bool {
			get, _ := f.GetString("get")

			if get != "" {
				s := bm.Redis.HGet("record", fmt.Sprintf("%v_%v_%v", u.Bot.ID, u.Chat.ID, get)).Val()

				if s == "" {
					botmaid.Reply(u, fmt.Sprintf("条目“%v”未被记录。", get))
					return true
				}

				botmaid.Reply(u, get+"：\n"+s)
			}

			if len(f.Args()) == 2 {
				bm.Redis.HSet("record", fmt.Sprintf("%v_%v_%v", u.Bot.ID, u.Chat.ID, f.Args()[1]), f.Args()[2])
				botmaid.Reply(u, f.Args()[1]+"已被记录。")
				return true
			}

			return false
		},
		Help: &botmaid.Help{
			Menu:  "record",
			Help:  "记录功能",
			Names: []string{"record", "rec"},
			Full: `使用方法：record [选项] 条目 内容

%v`,
			SetFlag: func(f *pflag.FlagSet) {
				f.String("get", "", "获得条目的内容")
			},
		},
	})
}
