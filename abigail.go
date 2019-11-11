package main

import (
	"fmt"
	"log"
	"time"

	"github.com/catsworld/botmaid"
	"github.com/spf13/pflag"

	_ "github.com/lib/pq"
)

var (
	rootDir string
	bm      *botmaid.BotMaid
	loc, _  = time.LoadLocation("Asia/Shanghai")
)

func init() {
	bm.AddCommand(&botmaid.Command{
		Do: func(u *botmaid.Update, f *pflag.FlagSet) bool {
			if bm.IsMaster(u.User) {
				u.User.NickName = "御主"
			}
			return false
		},
		Priority: 10005,
	})

	bm.AddCommand(&botmaid.Command{
		Do: bm.HelpCommandDo,
		Help: &botmaid.Help{
			Menu:  "help",
			Help:  "显示帮助",
			Names: []string{"help"},
			Usage: "使用方法：help [命令]",
		},
		Priority: 10000,
	})

	bm.AddCommand(&botmaid.Command{
		Do:       bm.HelpRespCommandDo,
		Priority: -10000,
	})

	bm.AddCommand(&botmaid.Command{
		Do: bm.MasterCommandDo,
		Help: &botmaid.Help{
			Menu:  "master",
			Help:  "增加/移除 master",
			Names: []string{"master"},
			Usage: "使用方法：master @用户",
		},
	})

	bm.AddCommand(&botmaid.Command{
		Do: bm.VersionCommandDo,
		Help: &botmaid.Help{
			Menu:    "version",
			Help:    "显示版本信息",
			Names:   []string{"version", "ver"},
			Usage:   "使用方法：version [选项]",
			SetFlag: bm.VersionCommandHelpSetFlag,
		},
	})

	bm.AddCommand(&botmaid.Command{
		Do: bm.VersetCommandDo,
		Help: &botmaid.Help{
			Menu:    "verset",
			Help:    "管理版本信息",
			Names:   []string{"verset"},
			Usage:   "使用方法：verset [选项] [内容]",
			SetFlag: bm.VersetCommandHelpSetFlag,
		},
	})

	bm.SubEntries = append(bm.SubEntries, "log")

	bm.AddCommand(&botmaid.Command{
		Do: bm.SubscribeCommandDo,
		Help: &botmaid.Help{
			Menu:  "subscribe",
			Help:  "订阅消息",
			Names: []string{"subscribe", "sub"},
			Usage: "使用方法：subscribe 条目",
		},
	})
}

func main() {
	err := bm.Start()
	if err != nil {
		log.Fatalf("[Fatal] Read config: %v\n", err)
	}
}

func reply(u *botmaid.Update, text string) (*botmaid.Update, error) {
	hide, _ := u.Message.Flags["roll"].GetBool("hide")

	if hide {
		bm.Reply(u, fmt.Sprintf("%v进行了一次暗骰。", bm.At(u.User)))

		uu := &(*u)
		uu.Chat = &botmaid.Chat{
			ID:   u.User.ID,
			Type: "private",
		}

		return bm.Reply(u, "[暗骰] "+text)
	}

	return bm.Reply(u, text)
}
