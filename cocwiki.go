package main

import (
	"fmt"
	"strings"

	"github.com/catsworld/botmaid"
	"github.com/catsworld/random"
)

func init() {
	bm.AddCommand(&botmaid.Command{
		Do: func(u *botmaid.Update, b *botmaid.Bot) bool {
			p := u.Message.Args[1]
			if !bm.Redis.HExists("wiki", p).Val() && !bm.Redis.HExists("wikiSynonym", p).Val() {
				b.Reply(u, fmt.Sprintf(random.String([]string{
					"%v是什么……？规则书上没有关于它的描述呀……",
				}), p))
				return true
			}

			if bm.Redis.HExists("wikiSynonym", p).Val() {
				p = bm.Redis.HGet("wikiSynonym", p).Val()
			}

			s := bm.Redis.HGet("wiki", p).Val()

			if strings.HasPrefix(s, "image|") {
				b.Reply(u, rootDir+"/images/wiki/"+strings.Replace(s, "image|", "", 1), "Image")
				return true
			}

			b.Reply(u, p+"：\n"+s)

			return true
		},
		Menu:       "wiki",
		MenuText:   "CoC 百科",
		Names:      []string{"wiki"},
		ArgsMinLen: 2,
		ArgsMaxLen: 2,
		Help:       " <词条> - 查询 CoC 百科",
	})
}
