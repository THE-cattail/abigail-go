package main

import (
	"fmt"
	"strings"

	"github.com/catsworld/botmaid"
	"github.com/spf13/pflag"
)

func init() {
	bm.AddCommand(&botmaid.Command{
		Do: func(u *botmaid.Update, f *pflag.FlagSet) bool {
			if len(f.Args()) == 2 {
				p := f.Args()[1]
				if !bm.Redis.HExists("wiki", p).Val() && !bm.Redis.HExists("wikiSynonym", p).Val() {
					bm.Reply(u, fmt.Sprintf("条目“%v”不存在。", p))
					return true
				}

				if bm.Redis.HExists("wikiSynonym", p).Val() {
					p = bm.Redis.HGet("wikiSynonym", p).Val()
				}

				s := bm.Redis.HGet("wiki", p).Val()

				if strings.HasPrefix(s, "image|") {
					bm.ReplyType(u, rootDir+"/images/wiki/"+strings.Replace(s, "image|", "", 1), "Image")
					return true
				}

				bm.Reply(u, p+"：\n"+s)
				return true
			}

			return false
		},
		Help: &botmaid.Help{
			Menu:  "wiki",
			Help:  "CoC 百科功能",
			Names: []string{"wiki"},
			Full: `使用方法：wiki 条目

%v`,
		},
	})
}
