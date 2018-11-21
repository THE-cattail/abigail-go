package main

import (
	"crypto/sha256"
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"

	"github.com/catsworld/api"
	"github.com/catsworld/botmaid"
	"github.com/catsworld/random"
)

var (
	trpgRecordFileName = make(map[int64]string)
	trpgRecordAgree    = make(map[int64]map[int64]bool)
	trpgRecordNickName = make(map[int64]map[int64]string)
	formatTRPGStart    = []string{
		"%s开团啦！输入“/join <名字>”加入哦～",
	}
	wordTRPGSave = []string{
		"存档咯。",
	}
)

func init() {
	bm.AddCommand(botmaid.Command{
		Do:       save,
		Priority: 10,
	})
	bm.AddCommand(botmaid.Command{
		Do:       trpgRecord,
		Priority: 10,
	})
	bm.AddCommand(botmaid.Command{
		Do:       trpg,
		Priority: 5,
		Menu:     "trpg",
		Names:    []string{"trpg"},
		Help: ` <名称> - 新建一个团或者覆盖之前同名团的存档
		load <名称> - 载入之前团的存档
		save - 存档
		join - 加入当前团
		join <昵称> - 加入当前团并设置昵称
		review <名称> - 显示之前存档的内容`,
	})
	bm.AddCommand(botmaid.Command{
		Do:       load,
		Priority: 5,
	})
	bm.AddCommand(botmaid.Command{
		Do:       join,
		Priority: 5,
	})
	bm.AddCommand(botmaid.Command{
		Do:       review,
		Priority: 5,
	})
}

func trpgRecord(e *api.Event, b *botmaid.Bot) bool {
	if _, ok := trpgRecordFileName[e.Place.ID]; ok && trpgRecordFileName[e.Place.ID] != "" {
		if _, ok := trpgRecordAgree[e.Place.ID]; ok || e.Sender.ID == b.Self.ID {
			if _, ok := trpgRecordAgree[e.Place.ID][e.Sender.ID]; ok && trpgRecordAgree[e.Place.ID][e.Sender.ID] || e.Sender.ID == b.Self.ID {
				f, err := os.OpenFile(rootDir+trpgRecordFileName[e.Place.ID], os.O_APPEND|os.O_WRONLY, 0777)
				if err != nil {
					return false
				}
				defer f.Close()
				if e.Sender.ID == b.Self.ID {
					f.Write([]byte("\n\n_" + strings.Replace(e.Message.Text, "\n", "_\n\n_", -1) + "_"))
				} else {
					f.Write([]byte("\n\n__" + e.Sender.NickName + "：__ " + e.Message.Text))
				}
			}
		}
	}
	return false
}

func startTRPG(e *api.Event, b *botmaid.Bot, name string) {
	trpgRecordFileName[e.Place.ID] = "TRPGLog/" + strconv.FormatInt(int64(e.Place.ID), 10) + "_" + name + ".md"
	trpgRecordAgree[e.Place.ID] = map[int64]bool{}
	trpgRecordNickName[e.Place.ID] = map[int64]string{}
	b.API.Push(api.Event{
		Message: &api.Message{
			Text: fmt.Sprintf(random.String(formatTRPGStart), name),
		},
		Place: e.Place,
	})
}

func trpg(e *api.Event, b *botmaid.Bot) bool {
	args := botmaid.SplitCommand(e.Message.Text)
	if b.IsCommand(e, "trpg") && len(args) > 1 {
		f, err := os.OpenFile(rootDir+"/TRPGLog/"+strconv.FormatInt(int64(e.Place.ID), 10)+"_"+args[1]+".md", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0777)
		if err != nil {
			return true
		}
		defer f.Close()
		startTRPG(e, b, args[1])
		_, err = f.Write([]byte("# " + args[1]))
		return true
	}
	return false
}

func load(e *api.Event, b *botmaid.Bot) bool {
	args := botmaid.SplitCommand(e.Message.Text)
	if b.IsCommand(e, "load") && len(args) > 1 {
		f, err := os.OpenFile(rootDir+"/TRPGLog/"+strconv.FormatInt(int64(e.Place.ID), 10)+"_"+args[1]+".md", os.O_WRONLY, 0777)
		if err != nil {
			return true
		}
		defer f.Close()
		startTRPG(e, b, args[1])
		return true
	}
	return false
}

func join(e *api.Event, b *botmaid.Bot) bool {
	args := botmaid.SplitCommand(e.Message.Text)
	if b.IsCommand(e, "join") {
		if _, ok := trpgRecordFileName[e.Place.ID]; !ok || trpgRecordFileName[e.Place.ID] == "" {
			return true
		}
		trpgRecordAgree[e.Place.ID][e.Sender.ID] = true
		trpgRecordNickName[e.Place.ID][e.Sender.ID] = e.Sender.NickName
		if len(args) > 1 {
			trpgRecordNickName[e.Place.ID][e.Sender.ID] = args[1]
		}
		return true
	}
	return false
}

func save(e *api.Event, b *botmaid.Bot) bool {
	if b.IsCommand(e, "save") && trpgRecordFileName[e.Place.ID] != "" {
		trpgRecordFileName[e.Place.ID] = ""
		b.API.Push(api.Event{
			Message: &api.Message{
				Text: random.String(wordTRPGSave),
			},
			Place: e.Place,
		})
		return true
	}
	return false
}

func review(e *api.Event, b *botmaid.Bot) bool {
	args := botmaid.SplitCommand(e.Message.Text)
	if b.IsCommand(e, "review") && len(args) > 1 {
		h := sha256.New()
		h.Write([]byte(strconv.FormatInt(int64(e.Place.ID), 10) + "_" + args[1]))
		bs := h.Sum(nil)
		cmd := exec.Command("pandoc", rootDir+"/TRPGLog/"+strconv.FormatInt(int64(e.Place.ID), 10)+"_"+args[1]+".md", "-o", rootDir+"/http/TRPGLog/"+fmt.Sprintf("%x", bs)+".html")
		cmd.Run()
		send(api.Event{
			Message: &api.Message{
				Text: bm.Conf.Get("HTTPServer.Domain").(string) + "TRPGLog/" + fmt.Sprintf("%x", bs) + ".html",
			},
			Place: e.Place,
		}, b, false)
		return true
	}
	return false
}

func trpgInit(e *api.Event, b *botmaid.Bot) bool {
	if _, ok := trpgRecordFileName[e.Place.ID]; !ok || trpgRecordFileName[e.Place.ID] == "" {
		return false
	}
	if _, ok := trpgRecordNickName[e.Place.ID][e.Sender.ID]; ok {
		e.Sender.NickName = trpgRecordNickName[e.Place.ID][e.Sender.ID]
	}
	return false
}
