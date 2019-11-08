package main

import (
	"fmt"
	"log"
	"time"

	"github.com/catsworld/botmaid"

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

func send(u *botmaid.Update, hide bool, origin *botmaid.Update) (*botmaid.Update, error) {
	if hide {
		botmaid.Reply(u, fmt.Sprintf("%v进行了一次暗骰。", botmaid.At(origin.User)))
		u.Message.Text = "[暗骰] " + u.Message.Text
		u.Chat = &botmaid.Chat{
			ID:   origin.User.ID,
			Type: "private",
		}
	}
	return (*origin.Bot.API).Push(u)
}
