package main

import (
	"fmt"
	"time"

	"github.com/the-cattail/abigail-go/coc"
	"github.com/the-cattail/abigail-go/cocmath"
	"github.com/the-cattail/botmaid"
	"github.com/the-cattail/botmaid/random"
	"github.com/spf13/pflag"
)

type scType struct {
	a, b   string
	Result coc.CheckResult
}

var (
	scMap = map[int64]map[int64]*scType{}
)

func init() {
	bm.AddCommand(&botmaid.Command{
		Do: func(u *botmaid.Update, f *pflag.FlagSet) bool {
			fmtInvalidSCExp := "%v，“%v”是不合法的 SAN Check 表达式，请查阅规则书中的相关条目。"

			if len(f.Args()) == 2 {
				if scMap[u.Chat.ID] == nil {
					scMap[u.Chat.ID] = map[int64]*scType{}
				}

				now := 0
				for i, v := range f.Args()[1] {
					if v == '(' {
						now++
					}
					if v == ')' {
						now--
					}

					if v == '/' && now == 0 {
						if scMap[u.Chat.ID][u.User.ID] != nil {
							reply(u, fmt.Sprintf(fmtInvalidSCExp, bm.At(u.User), f.Args()[1]))
							return true
						}

						scMap[u.Chat.ID][u.User.ID] = &scType{
							a: f.Args()[1][0:i],
							b: f.Args()[1][i+1 : len(f.Args()[1])],
						}
					}
				}

				if scMap[u.Chat.ID][u.User.ID] == nil {
					reply(u, fmt.Sprintf(fmtInvalidSCExp, bm.At(u.User), f.Args()[1]))
					return true
				}

				_, err := cocmath.New(scMap[u.Chat.ID][u.User.ID].a)
				if err != nil {
					reply(u, fmt.Sprintf(fmtInvalidSCExp, bm.At(u.User), f.Args()[1]))
					return true
				}

				_, err = cocmath.New(scMap[u.Chat.ID][u.User.ID].b)
				if err != nil {
					reply(u, fmt.Sprintf(fmtInvalidSCExp, bm.At(u.User), f.Args()[1]))
					return true
				}

				reply(u, fmt.Sprintf("%v，请进行一次意志检定。", bm.At(u.User)))
				return true
			}

			return false
		},
		Help: &botmaid.Help{
			Menu:  "sc",
			Help:  "SAN check 功能",
			Names: []string{"sc"},
			Usage: "使用方法：sc SANCheck表达式",
		},
	})
}

func scResp(u *botmaid.Update) {
	time.Sleep(time.Second * time.Duration(random.Int(1, 3)))

	a, _ := cocmath.New(scMap[u.Chat.ID][u.User.ID].a)
	b, _ := cocmath.New(scMap[u.Chat.ID][u.User.ID].b)

	r := 0

	if scMap[u.Chat.ID][u.User.ID].Result.Great == coc.GreatSucc {
		r = a.Result.Min
	} else if scMap[u.Chat.ID][u.User.ID].Result.Great == coc.GreatFail {
		r = b.Result.Max
	} else if scMap[u.Chat.ID][u.User.ID].Result.Succ == coc.Succ {
		r = a.Result.Value
	} else {
		r = b.Result.Value
	}

	scMap[u.Chat.ID][u.User.ID] = nil
	reply(u, fmt.Sprintf("%v的理智损失了 %v 点。", bm.At(u.User), r))
}
