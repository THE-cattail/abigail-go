package main

import (
	"fmt"

	"github.com/catsworld/botmaid"
	"github.com/spf13/pflag"

	"github.com/Pallinder/go-randomdata"
	"github.com/goroom/rand"
	"github.com/mattn/go-gimei"
)

func init() {
	bm.AddCommand(&botmaid.Command{
		Do: func(u *botmaid.Update, f *pflag.FlagSet) bool {
			lang, _ := f.GetString("lang")

			if botmaid.In(lang, "cn", "中", "中文") {
				botmaid.Reply(u, rand.GetRand().ChineseName())
				return true
			}

			if botmaid.In(lang, "en", "英", "英文") {
				botmaid.Reply(u, randomdata.FullName(randomdata.RandomGender))
				return true
			}

			if botmaid.In(lang, "jp", "日", "日文") {
				botmaid.Reply(u, gimei.NewName().Kanji())
				return true
			}

			botmaid.Reply(u, fmt.Sprintf("语言“%v”未收录", lang))
			return true
		},
		Help: &botmaid.Help{
			Menu:  "name",
			Help:  "随机生成角色姓名",
			Names: []string{"name"},
			Full: `使用方法：name [选项]

%v`,
			SetFlag: func(f *pflag.FlagSet) {
				f.String("lang", "cn", "指定生成姓名的语言（cn、en 或 jp）")
			},
		},
	})
}
