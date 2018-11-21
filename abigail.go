package main

import (
	"database/sql"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/catsworld/api"
	"github.com/catsworld/botmaid"
	"github.com/pelletier/go-toml"

	_ "github.com/lib/pq"
)

var (
	rootDir string

	bm = botmaid.BotMaid{
		Bots: map[string]*botmaid.Bot{},
		HelpMenus: map[string]string{
			"roll":    "Roll点",
			"sc":      "SAN Check",
			"call":    "点名",
			"record":  "记录",
			"pk":      "对抗",
			"cocwiki": "CoC 百科",
			"trpg":    "开一个新团并记录跑团过程",
			"gugugu":  "记录已跑过的团",
		},
		Words: map[string][]string{
			"selfIntro": []string{
				"我是阿比盖尔，他们都叫我塞勒姆的魔女呢，呵呵，可别把我惹急了。%s，要叫出我的话，在命令前敲上“/”、“:”或者“：”就可以了。",
			},
			"undefCommand": []string{
				"%s？那是什么啊，乖孩子是不知道的呢。",
			},
			"invalidMaster": []string{
				"要用 At 和我说哦，不然我也没法知道%s的用户名呢。",
			},
			"masterExisted": []string{
				"%s已经是我的御主了。",
			},
			"masterAdded": []string{
				"%s，你好！我是阿比盖尔——阿比盖尔·威廉姆斯，我是Fo……reigner……你就是御主吗？如果你不介意的话，希望你能叫我阿比。我想我们很快就能成为朋友。",
			},
			"masterNotExisted": []string{
				"%s不是我的御主呢。",
			},
			"masterRemoved": []string{
				"以后%s就不是我的御主了哦。",
			},
			"testPlaceAdded": []string{
				"以后可以在这里练习吗？呵呵，感觉好像挺不错的。",
			},
			"testPlaceRemoved": []string{
				"不能在这里练习了吗……",
			},
		},
		RespTime: time.Now(),
	}

	loc, _ = time.LoadLocation("Asia/Shanghai")
)

func init() {
	bm.AddCommand(botmaid.Command{
		Do:       trpgInit,
		Priority: 200,
	})
}

func main() {
	var err error

	rootDir, err = filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatalf("[Fatal] Read rootdir: %v\n", err)
	}

	go func() {
		http.Handle("/", http.FileServer(http.Dir(rootDir+"/http")))
		http.ListenAndServe(":8570", nil)
	}()

	raw, err := ioutil.ReadFile(rootDir + "/config.toml")
	if err != nil {
		log.Fatalf("[Fatal] Read config: %v\n", err)
	}
	bm.Conf, err = toml.Load(string(raw))
	if err != nil {
		log.Fatalf("[Fatal] Read config: %v\n", err)
	}

	bm.DB, err = sql.Open("postgres", "user="+bm.Conf.Get("Database.User").(string)+" password="+bm.Conf.Get("Database.Password").(string)+" dbname="+bm.Conf.Get("Database.DBName").(string)+" sslmode=disable")
	if err != nil {
		log.Fatalf("[Fatal] Connect database: %v\n", err)
	}

	err = bm.InitBroadcastTable("jrrp")
	if err != nil {
		log.Fatalf("[Fatal] Init jrrp: %v\n", err)
	}

	err = bm.Start()
	if err != nil {
		log.Fatalf("[Fatal] Read config: %v\n", err)
	}
}

func send(e api.Event, b *botmaid.Bot, hide bool) (api.Event, error) {
	e.Sender = &api.User{
		ID:       b.Self.ID,
		NickName: b.Self.NickName,
	}
	if hide {
		e.Message.Text = "[暗骰] " + e.Message.Text
	}
	if _, ok := trpgRecordAgree[e.Place.ID]; ok {
		if _, ok := trpgRecordAgree[e.Place.ID][e.Sender.ID]; ok && trpgRecordAgree[e.Place.ID][e.Sender.ID] {
			trpgRecord(&e, b)
		}
	}
	if hide {
		e.Place = &api.Place{
			ID:   e.Sender.ID,
			Type: "private",
		}
	}
	log.Println(e)
	return b.API.Push(e)
}
