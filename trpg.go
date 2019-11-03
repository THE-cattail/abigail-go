package main

import (
	"crypto/sha256"
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"

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

func trpgRecord(u *botmaid.Update, b *botmaid.Bot) bool {
	if _, ok := trpgRecordFileName[u.Chat.ID]; ok && trpgRecordFileName[u.Chat.ID] != "" {
		if _, ok := trpgRecordAgree[u.Chat.ID]; ok || u.User.ID == b.Self.ID {
			if _, ok := trpgRecordAgree[u.Chat.ID][u.User.ID]; ok && trpgRecordAgree[u.Chat.ID][u.User.ID] || u.User.ID == b.Self.ID {
				f, err := os.OpenFile(rootDir+trpgRecordFileName[u.Chat.ID], os.O_APPEND|os.O_WRONLY, 0777)
				if err != nil {
					return false
				}
				defer f.Close()
				if u.User.ID == b.Self.ID {
					f.Write([]byte("\n\n_" + strings.Replace(u.Message.Text, "\n", "_\n\n_", -1) + "_"))
				} else {
					f.Write([]byte("\n\n__" + u.User.NickName + "：__ " + u.Message.Text))
				}
			}
		}
	}
	return false
}

func startTRPG(u *botmaid.Update, b *botmaid.Bot, name string) {
	trpgRecordFileName[u.Chat.ID] = "TRPGLog/" + strconv.FormatInt(int64(u.Chat.ID), 10) + "_" + name + ".md"
	trpgRecordAgree[u.Chat.ID] = map[int64]bool{}
	trpgRecordNickName[u.Chat.ID] = map[int64]string{}
	b.API.Send(botmaid.Update{
		Message: &botmaid.Message{
			Text: fmt.Sprintf(random.String(formatTRPGStart), name),
		},
		Chat: u.Chat,
	})
}

func trpg(u *botmaid.Update, b *botmaid.Bot) bool {
	args := botmaid.SplitCommand(u.Message.Text)
	if b.IsCommand(u, "trpg") && len(args) > 1 {
		f, err := os.OpenFile(rootDir+"/TRPGLog/"+strconv.FormatInt(int64(u.Chat.ID), 10)+"_"+args[1]+".md", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0777)
		if err != nil {
			return true
		}
		defer f.Close()
		startTRPG(u, b, args[1])
		_, err = f.Write([]byte("# " + args[1]))
		return true
	}
	return false
}

func load(u *botmaid.Update, b *botmaid.Bot) bool {
	args := botmaid.SplitCommand(u.Message.Text)
	if b.IsCommand(u, "load") && len(args) > 1 {
		f, err := os.OpenFile(rootDir+"/TRPGLog/"+strconv.FormatInt(int64(u.Chat.ID), 10)+"_"+args[1]+".md", os.O_WRONLY, 0777)
		if err != nil {
			return true
		}
		defer f.Close()
		startTRPG(u, b, args[1])
		return true
	}
	return false
}

func join(u *botmaid.Update, b *botmaid.Bot) bool {
	args := botmaid.SplitCommand(u.Message.Text)
	if b.IsCommand(u, "join") {
		if _, ok := trpgRecordFileName[u.Chat.ID]; !ok || trpgRecordFileName[u.Chat.ID] == "" {
			return true
		}
		trpgRecordAgree[u.Chat.ID][u.User.ID] = true
		trpgRecordNickName[u.Chat.ID][u.User.ID] = u.User.NickName
		if len(args) > 1 {
			trpgRecordNickName[u.Chat.ID][u.User.ID] = args[1]
		}
		return true
	}
	return false
}

func save(u *botmaid.Update, b *botmaid.Bot) bool {
	if b.IsCommand(u, "save") && trpgRecordFileName[u.Chat.ID] != "" {
		trpgRecordFileName[u.Chat.ID] = ""
		b.API.Send(botmaid.Update{
			Message: &botmaid.Message{
				Text: random.String(wordTRPGSave),
			},
			Chat: u.Chat,
		})
		return true
	}
	return false
}

func review(u *botmaid.Update, b *botmaid.Bot) bool {
	args := botmaid.SplitCommand(u.Message.Text)
	if b.IsCommand(u, "review") && len(args) > 1 {
		h := sha256.New()
		h.Write([]byte(strconv.FormatInt(int64(u.Chat.ID), 10) + "_" + args[1]))
		bs := h.Sum(nil)
		cmd := exec.Command("pandoc", rootDir+"/TRPGLog/"+strconv.FormatInt(int64(u.Chat.ID), 10)+"_"+args[1]+".md", "-o", rootDir+"/http/TRPGLog/"+fmt.Sprintf("%x", bs)+".html")
		cmd.Run()
		send(u, botmaid.Update{
			Message: &botmaid.Message{
				Text: bm.Conf.Get("HTTPServer.Domain").(string) + "TRPGLog/" + fmt.Sprintf("%x", bs) + ".html",
			},
			Chat: u.Chat,
		}, b, false)
		return true
	}
	return false
}

func trpgInit(u *botmaid.Update, b *botmaid.Bot) bool {
	if _, ok := trpgRecordFileName[u.Chat.ID]; !ok || trpgRecordFileName[u.Chat.ID] == "" {
		return false
	}
	if _, ok := trpgRecordNickName[u.Chat.ID][u.User.ID]; ok {
		u.User.NickName = trpgRecordNickName[u.Chat.ID][u.User.ID]
	}
	return false
}
