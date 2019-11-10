package main

import (
	"fmt"

	"github.com/catsworld/botmaid"
	"github.com/spf13/pflag"
)

func init() {
	bm.AddCommand(&botmaid.Command{
		Do: func(u *botmaid.Update, f *pflag.FlagSet) bool {
			del, _ := f.GetBool("del")
			if del {
				bm.Redis.HDel("record", fmt.Sprintf("%v_%v_%v", u.Bot.ID, u.Chat.ID, f.Args()[1]))
				bm.Reply(u, f.Args()[1]+"已被删除。")
				return true
			}

			if len(f.Args()) == 2 {
				s := bm.Redis.HGet("record", fmt.Sprintf("%v_%v_%v", u.Bot.ID, u.Chat.ID, f.Args()[1])).Val()

				if s == "" {
					bm.Reply(u, fmt.Sprintf("条目“%v”未被记录。", f.Args()[1]))
					return true
				}

				bm.Reply(u, f.Args()[1]+"：\n"+s)
				return true
			}

			if len(f.Args()) == 3 {
				bm.Redis.HSet("record", fmt.Sprintf("%v_%v_%v", u.Bot.ID, u.Chat.ID, f.Args()[1]), f.Args()[2])
				bm.Reply(u, f.Args()[1]+"已被记录。")
				return true
			}

			return false
		},
		Help: &botmaid.Help{
			Menu:  "record",
			Help:  "记录功能",
			Names: []string{"record", "rec"},
			Usage: "使用方法：record 条目 [内容]",
			SetFlag: func(f *pflag.FlagSet) {
				f.BoolP("del", "d", false, "删除指定条目")
			},
		},
	})
}
