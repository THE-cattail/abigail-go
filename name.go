package main

import (
	"github.com/catsworld/botmaid"

	"github.com/Pallinder/go-randomdata"
	"github.com/goroom/rand"
	"github.com/mattn/go-gimei"
)

func init() {
	bm.AddCommand(&botmaid.Command{
		Do: func(u *botmaid.Update) bool {
			p := "cn"
			if len(u.Message.Args) > 1 {
				p = u.Message.Args[1]
			}
			if botmaid.In(p, "cn", "中", "中文") {
				botmaid.Reply(u, rand.GetRand().ChineseName())
			} else if botmaid.In(p, "en", "英", "英文") {
				botmaid.Reply(u, randomdata.FullName(randomdata.RandomGender))
			} else if botmaid.In(p, "jp", "日", "日文") {
				name := gimei.NewName()
				botmaid.Reply(u, name.Kanji())
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
