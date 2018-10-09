package main

import (
	"database/sql"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"sort"
	"time"

	"github.com/catsworld/api"
	"github.com/catsworld/botmaid"
	"github.com/pelletier/go-toml"

	_ "github.com/lib/pq"
)

var (
	apiMaid   = &botmaid.BotMaid{}
	rootDir   string
	conf      *toml.Tree
	db        *sql.DB
	commands  = []botmaid.Command{}
	masters   = map[string][]string{}
	testChats = map[string][]string{}
	help      = botmaid.Help{
		SelfIntro: "我是阿比盖尔，他们都叫我塞勒姆的魔女呢，呵呵，可别把我惹急了。要叫出我的话，在命令前敲上“/”、“:”或者“：”就可以了。",
		HelpMenu: `roll [r] - Roll点
sc [sancheck] - SAN Check
call - 点名
record [rec] - 记录
pk - 对抗
cocwiki - CoC 百科
trpg - 开一个新团并记录跑团过程
gugugu - 记录已跑过的团`,
		UndefCommand: "那是什么啊，乖孩子是不知道的哦。",
		HelpSubMenu: map[string]string{
			"roll": `roll <说明（可省略）> <表达式> - 进行一次表达式计算/检定
roll <说明（可省略）> <数值（可省略）> - 进行一次d100的检定
roll <列表> - roll列表。
roll --hide [-h] <命令> - 暗投
tempmad - roll一次临时疯狂症状
character [/char] (--full [-f]) - roll一张人物卡`,
			"sc": "sc <SANCheck公式> - 进行一次SAN Check",
			"call": `call <@其他人> - 进行一次点名
call --status [-s] - 查看当前点名情况
call --gugugu <名称> - 用记录的成员点名
咕咕咕 - 溜了溜了`,
			"record": `record <名称> <内容> - 进行记录
record <名称> - 查看记录的内容
record --del [-d] <名称> - 删除记录
record --list [-l] - 列出全部记录`,
			"pk":      "pk - 进行一次对抗",
			"cocwiki": "cocwiki <词条> - 在 CoC 百科中查询资料",
			"trpg": `trpg <名称> - 新建一个团或者覆盖之前同名团的存档
load <名称> - 载入之前团的存档
save - 存档
join - 加入当前团
join <昵称> - 加入当前团并设置昵称
review <名称> - 显示之前存档的内容`,
			"gugugu": `gugugu <名称> <成员> - 记录一个新团
gugugu --complete [-c] <名称> - 结一个团
gugugu --del [-d] <名称> - 删除一个团
gugugu --list (--all [-a]) - 查看记录的团`,
		},
	}
	loc *time.Location
)

func init() {
	botmaid.AddCommand(&commands, trpgInit, 200)
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

	loc, err = time.LoadLocation("Asia/Shanghai")
	if err != nil {
		log.Fatalf("[Fatal] Load location: %v\n", err)
	}

	raw, err := ioutil.ReadFile(rootDir + "/config.toml")
	if err != nil {
		log.Fatalf("[Fatal] Read config: %v\n", err)
	}
	conf, err = toml.Load(string(raw))
	if err != nil {
		log.Fatalf("[Fatal] Read config: %v\n", err)
	}

	botmaid.RegHelpCommand(&commands, &help)

	db, err = sql.Open("postgres", "user="+conf.Get("Database.User").(string)+" password="+conf.Get("Database.Password").(string)+" dbname="+conf.Get("Database.DBName").(string)+" sslmode=disable")
	if err != nil {
		log.Fatalf("[Fatal] Connect database: %v\n", err)
	}

	err = apiMaid.Init(conf)
	if err != nil {
		log.Fatalf("[Fatal] Read config: %v\n", err)
	}

	log.Println("加载完成！")

	sort.Sort(botmaid.CommandSlice(commands))

	apiMaid.Run(conf, commands)
}

func send(e *api.Event, b *botmaid.Bot, hide bool) (int64, error) {
	e.Sender = &api.User{
		ID:       b.Self.ID,
		NickName: b.Self.NickName,
	}
	if hide {
		e.Message.Text = "[暗骰] " + e.Message.Text
	}
	if _, ok := trpgRecordAgree[e.Place.ID]; ok {
		if _, ok := trpgRecordAgree[e.Place.ID][e.Sender.ID]; ok && trpgRecordAgree[e.Place.ID][e.Sender.ID] {
			trpgRecord(e, b)
		}
	}
	if hide {
		e.Place = &api.Place{
			ID:   e.Sender.ID,
			Type: "private",
		}
	}
	return b.API.Push(e)
}
