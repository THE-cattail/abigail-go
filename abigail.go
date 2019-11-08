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

func reply(u *botmaid.Update, text string) (*botmaid.Update, error) {
	hide, _ := bm.Flags["roll"].GetBool("hide")

	if hide {
		botmaid.Reply(u, fmt.Sprintf("%v进行了一次暗骰。", botmaid.At(u.User)))

		uu := &(*u)
		uu.Chat = &botmaid.Chat{
			ID:   u.User.ID,
			Type: "private",
		}

		return botmaid.Reply(u, "[暗骰] "+text)
	}

	return botmaid.Reply(u, text)
}
