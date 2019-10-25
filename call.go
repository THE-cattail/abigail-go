package main

import (
	"fmt"
	"strings"

	"github.com/catsworld/botmaid"
	"github.com/catsworld/random"
)

type callType struct {
	Status bool
	List   map[string]bool
	Resped map[string]bool
	Total  int
	Get    int
}

var (
	callMap = make(map[int64]*callType)
)

func init() {
	bm.AddCommand(&botmaid.Command{
		Do: func(u *botmaid.Update, b *botmaid.Bot) bool {
			if strings.HasPrefix(u.Message.Text, "咕") {
				if _, ok := callMap[u.Chat.ID]; ok && callMap[u.Chat.ID].Status && callMap[u.Chat.ID].List[b.At(u.User)[0]] {
					callMap[u.Chat.ID].Status = false
					callMap[u.Chat.ID].List = map[string]bool{}
					b.Reply(u, fmt.Sprintf(random.String([]string{
						"看啊看啊！%v这家伙咕了哦！",
						"%v说他不在哦_(:з」∠)_今天的玩乐是不是到此为止了呢？",
					}), u.User.NickName))
				}
			}
			return false
		},
		Priority: 1000,
	})
	bm.AddCommand(&botmaid.Command{
		Do: func(u *botmaid.Update, b *botmaid.Bot) bool {
			if _, ok := callMap[u.Chat.ID]; !ok || !callMap[u.Chat.ID].Status {
				return false
			}
			if callMap[u.Chat.ID].List[b.At(u.User)[0]] && !callMap[u.Chat.ID].Resped[b.At(u.User)[0]] {
				callMap[u.Chat.ID].Resped[b.At(u.User)[0]] = true
				callMap[u.Chat.ID].Get++
			}
			if callMap[u.Chat.ID].Get == callMap[u.Chat.ID].Total {
				callMap[u.Chat.ID].Status = false
				b.Reply(u, random.String([]string{
					"调查员都已经聚集好了哦。呵呵，是不是又有什么事情要发生了。",
					"大家都已经在这里啦，让我们开开心心地开始今天的活动吧，诶嘿嘿☆~",
				}))
			}
			return false
		},
		Priority: 1000,
	})
	bm.AddCommand(&botmaid.Command{
		Do: func(u *botmaid.Update, b *botmaid.Bot) bool {
			if botmaid.In(u.Message.Args[1], "status") {
				if _, ok := callMap[u.Chat.ID]; !ok || !callMap[u.Chat.ID].Status {
					b.Reply(u, random.String([]string{
						"好像有什么很热闹的样子呢……我也好想被叫去参加呢。",
						"稍微有点无聊呢……有没有人带阿比玩啊——",
					}))
					return true
				}
				if callMap[u.Chat.ID].Get == 0 {
					b.Reply(u, random.String([]string{
						"还没有谁在哦，果然大家都是鸽子吧（生气）。",
						"大家都不知道去哪里啦——有没有人看到他们呀——",
					}))
					return true
				}
				gu := ""
				for key := range callMap[u.Chat.ID].List {
					if !callMap[u.Chat.ID].Resped[key] {
						gu += key + " "
					}
				}
				gu = gu[:len(gu)-1]
				b.Reply(u, fmt.Sprintf(random.String([]string{
					"已经有%v位伙伴出现啦~\n鸽子名单：%v",
					"刚刚有%v只调查员来过了哦~\n不过%v阿比还没有看见-v-",
				}), callMap[u.Chat.ID].Get, gu))
				return true
			}
			return false
		},
		Menu:       "call",
		Names:      []string{"call"},
		ArgsMinLen: 2,
		ArgsMaxLen: 2,
		Help:       " status - 查看当前点名情况",
	})
	bm.AddCommand(&botmaid.Command{
		Do: func(u *botmaid.Update, b *botmaid.Bot) bool {
			callMap[u.Chat.ID] = &callType{
				Status: true,
				Total:  len(u.Message.Args) - 1,
				Get:    0,
				List:   map[string]bool{},
				Resped: map[string]bool{},
			}
			for i := 1; i < len(u.Message.Args); i++ {
				callMap[u.Chat.ID].List[u.Message.Args[i]] = true
			}
			b.Reply(u, random.String([]string{
				"呵呵呵，看来调查员的召集开始了。",
				"你们又要演出新的戏剧了吗？阿比也来看吧w",
			}))
			return true
		},
		Menu:       "call",
		MenuText:   "点名",
		Names:      []string{"call"},
		ArgsMinLen: 2,
		Help:       " <@其他人> - 进行一次点名",
	})
}
