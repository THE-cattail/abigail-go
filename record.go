package main

import "github.com/catsworld/botmaid"

func init() {
	bm.AddCommand(&botmaid.Command{
		Do: func(u *botmaid.Update) bool {
			return true
		},
		Menu:       "record",
		MenuText:   "记录条目",
		Names:      []string{"record", "rec"},
		ArgsMinLen: 3,
		ArgsMaxLen: 3,
		Help:       " <条目> <内容> - 记录条目内容",
	})
}
