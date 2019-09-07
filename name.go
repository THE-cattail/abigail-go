package main

import (
	"github.com/catsworld/botmaid"

	"github.com/Pallinder/go-randomdata"
	"github.com/goroom/rand"
	"github.com/mattn/go-gimei"
)

func init() {
	bm.AddCommand(botmaid.Command{
		Do: func(u *botmaid.Update, b *botmaid.Bot) bool {
			p := "cn"
			if len(u.Message.Args) > 1 {
				p = u.Message.Args[1]
			}
			if p == "cn" {
				b.Reply(u, rand.GetRand().ChineseName())
			} else if p == "en" {
				b.Reply(u, randomdata.FullName(randomdata.RandomGender))
			} else if p == "jp" {
				name := gimei.NewName()
				b.Reply(u, name.Kanji())
			} else {
				return false
			}
			return true
		},
		Menu:       "name",
		MenuText:   "随机名字",
		Names:      []string{"name"},
		ArgsMinLen: 1,
		ArgsMaxLen: 2,
		Help:       " <cn/en/jp> - 生成一个随机名字",
	})
}
