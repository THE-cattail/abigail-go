package main

import (
	"fmt"

	"github.com/catsworld/abigail/coc"
	"github.com/catsworld/botmaid"
	"github.com/catsworld/botmaid/random"
)

type pkRollResult struct {
	User   *botmaid.User
	Result coc.CheckResult
}

type pkType struct {
	Status  bool
	Results []pkRollResult
}

var (
	pkMap = map[int64]*pkType{}
)

func init() {
	bm.AddCommand(&botmaid.Command{
		Do: func(u *botmaid.Update) bool {
			pkMap[u.Chat.ID] = &pkType{
				Status: true,
			}
			send(&botmaid.Update{
				Message: &botmaid.Message{
					Text: random.String([]string{
						"开始对抗，请一方 roll 点。",
						"对抗检定~请其中一边先 roll 点~",
					}),
				},
				Chat: u.Chat,
			}, false, u)
			return true
		},
		Menu:     "pk",
		MenuText: "对抗检定",
		Names:    []string{"pk"},
		Help:     " - 进行一次对抗",
	})
}

func pkResp(u *botmaid.Update) {
	if _, ok := pkMap[u.Chat.ID]; !ok || !pkMap[u.Chat.ID].Status {
		return
	}
	if len(pkMap[u.Chat.ID].Results) == 1 {
		send(&botmaid.Update{
			Message: &botmaid.Message{
				Text: random.String([]string{
					"请另一方 roll 点。",
					"接下来请另一边 roll 点~",
				}),
			},
			Chat: u.Chat,
		}, false, u)
		return
	}
	pkResult := coc.PK(pkMap[u.Chat.ID].Results[0].Result, pkMap[u.Chat.ID].Results[1].Result)
	message := ""
	if pkResult == coc.PKDraw {
		message = random.String([]string{
			"对抗结果：平局！",
			"这次对抗平局了哦~",
		})
	} else {
		victor := pkMap[u.Chat.ID].Results[0].User.NickName
		if pkResult == coc.PKBWin {
			victor = pkMap[u.Chat.ID].Results[1].User.NickName
		}
		message = fmt.Sprintf(random.String([]string{
			"对抗结果：%v胜利！",
			"这次对抗由%v取得了胜利~",
		}), victor)
	}
	pkMap[u.Chat.ID] = &pkType{
		Status: false,
	}
	send(&botmaid.Update{
		Message: &botmaid.Message{
			Text: message,
		},
		Chat: u.Chat,
	}, false, u)
}
