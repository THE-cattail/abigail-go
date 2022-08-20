package main

import (
	"fmt"

	"github.com/the-cattail/abigail-gococ"
	"github.com/the-cattail/botmaid"
	"github.com/spf13/pflag"
)

type pkRollResult struct {
	User   *botmaid.User
	Result coc.CheckResult
}

type pkType struct {
	Results []pkRollResult
}

var (
	pkMap = map[int64]*pkType{}
)

func init() {
	bm.AddCommand(&botmaid.Command{
		Do: func(u *botmaid.Update, f *pflag.FlagSet) bool {
			pkMap[u.Chat.ID] = &pkType{}

			reply(u, "进行对抗检定，请其中一方掷骰。")
			return true
		},
		Help: &botmaid.Help{
			Menu:  "pk",
			Help:  "对抗检定功能",
			Names: []string{"pk"},
			Usage: "使用方法：pk",
		},
	})
}

func pkResp(u *botmaid.Update) {
	if pkMap[u.Chat.ID] == nil {
		return
	}

	if len(pkMap[u.Chat.ID].Results) == 1 {
		reply(u, "请另一方掷骰。")
		return
	}

	r := coc.PK(pkMap[u.Chat.ID].Results[0].Result, pkMap[u.Chat.ID].Results[1].Result)

	if r == coc.PKDraw {
		reply(u, "对抗检定的结果为平局。")
		return
	}

	v := bm.At(pkMap[u.Chat.ID].Results[0].User)
	if r == coc.PKBWin {
		v = bm.At(pkMap[u.Chat.ID].Results[1].User)
	}

	pkMap[u.Chat.ID] = nil
	reply(u, fmt.Sprintf("对抗检定的结果为%v胜利。", v))
}
