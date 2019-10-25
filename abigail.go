package main

import (
	"fmt"
	"log"
	"time"

	"github.com/catsworld/botmaid"
	"github.com/catsworld/random"

	_ "github.com/lib/pq"
)

var (
	rootDir string
	bm      *botmaid.BotMaid
	loc, _  = time.LoadLocation("Asia/Shanghai")
)

func main() {
	err := bm.Start()
	if err != nil {
		log.Fatalf("[Fatal] Read config: %v\n", err)
	}
}

func send(b *botmaid.Bot, u botmaid.Update, hide bool, origin *botmaid.Update) (botmaid.Update, error) {
	if hide {
		b.Reply(&u, fmt.Sprintf(random.String([]string{
			"%v的暗骰。",
			"%v进行了一次暗骰。",
			"%v开始暗骰了，没人知道到底发生了什么~",
			"%v正在瞎编……",
		}), origin.User.NickName))
		u.Message.Text = "[暗骰] " + u.Message.Text
		u.Chat = &botmaid.Chat{
			ID:   origin.User.ID,
			Type: "private",
		}
	}
	return b.API.Push(u)
}
