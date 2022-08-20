package main

import (
	"fmt"

	"github.com/the-cattail/botmaid"
	"github.com/spf13/pflag"

	"github.com/Pallinder/go-randomdata"
	"github.com/goroom/rand"
	"github.com/mattn/go-gimei"
)

func init() {
	bm.AddCommand(&botmaid.Command{
		Do: func(u *botmaid.Update, f *pflag.FlagSet) bool {
			// lang, _ := f.GetString("lang")
			lang := "cn"
			if len(u.Message.Args) == 2 {
				lang = u.Message.Args[1]
			}

			if lang == "cn" {
				bm.Reply(u, rand.GetRand().ChineseName())
				return true
			}

			if lang == "en" {
				bm.Reply(u, randomdata.FullName(randomdata.RandomGender))
				return true
			}

			if lang == "jp" {
				bm.Reply(u, gimei.NewName().Kanji())
				return true
			}

			bm.Reply(u, fmt.Sprintf("语言“%v”未收录", lang))
			return true
		},
		Help: &botmaid.Help{
			Menu:  "name",
			Help:  "随机生成角色姓名",
			Names: []string{"name"},
			Usage: "使用方法：name [选项]",
			SetFlag: func(f *pflag.FlagSet) {
				f.String("lang", "cn", "指定生成姓名的语言（cn、en 或 jp）")
			},
		},
	})
}
