package main

import (
	"fmt"
	"log"

	"github.com/catsworld/botmaid"
	"github.com/catsworld/random"
	"github.com/catsworld/slices"
)

type dbAbiGugugu struct {
	ID      int64
	ChatID  int64
	Name    string
	Members string
	At      string
	Status  int64
}

var (
	wordGuguguStatus = []string{"未结", "已结"}
	formatGugugu     = []string{
		"%s开始咯！",
	}
	formatGuguguComplete = []string{
		"%s结团咯！",
	}
	formatGuguguDelete = []string{
		"%s删除咯！",
	}
	wordGuguguListEmpty = []string{
		"NULL",
	}
	wordGuguguList = []string{
		"鸽子的离去是因为风的追求，还是树的不挽留：",
	}
)

func init() {
	bm.AddCommand(botmaid.Command{
		Do:       guguguList,
		Priority: 5,
	})
	bm.AddCommand(botmaid.Command{
		Do:       guguguComplete,
		Priority: 5,
	})
	bm.AddCommand(botmaid.Command{
		Do:       guguguDelete,
		Priority: 5,
	})
	bm.AddCommand(botmaid.Command{
		Do:       gugugu,
		Priority: 5,
		Menu:     "gugugu",
		Names:    []string{"gugugu"},
		Help: ` <名称> <成员> - 记录一个新团
gugugu -complete [-c] <名称> - 结一个团
gugugu -del [-d] <名称> - 删除一个团
gugugu -list (-all [-a]) - 查看记录的团`,
	})
}

func guguguInit() {
	stmt, err := bm.DB.Prepare(`CREATE TABLE abi_gugugu (
		id SERIAL primary key,
		chat_id bigint not null,
		name text,
		members text,
		at text,
		status integer
	)`)
	if err != nil {
		log.Panic(err)
	}
	stmt.Exec()
}

func gugugu(u *botmaid.Update, b *botmaid.Bot) bool {
	args := botmaid.SplitCommand(u.Message.Text)
	if b.IsCommand(u, "gugugu") && len(args) > 1 {
		members := ""
		at := ""
		if len(args) > 2 && args[2] == "-at" {
			for i := 3; i+1 < len(args); i += 2 {
				if members != "" {
					members += "、"
				}
				members += args[i]
				if at != "" {
					at += " "
				}
				at += args[i+1]
			}
		} else {
			for i := 2; i < len(args); i++ {
				if members != "" {
					members += "、"
				}
				members += args[i]
			}
		}
		log.Println(members, at)
		theGugugu := dbAbiGugugu{}
		err := bm.DB.QueryRow("SELECT * FROM abi_gugugu WHERE chat_id = $1 AND name = $2", u.Chat.ID, args[1]).Scan(&theGugugu.ID, &theGugugu.ChatID, &theGugugu.Name, &theGugugu.Members, &theGugugu.At, &theGugugu.Status)
		if err != nil {
			stmt, err := bm.DB.Prepare("INSERT INTO abi_gugugu(chat_id, name, members, at, status) VALUES($1, $2, $3, $4, $5)")
			if err != nil {
				log.Println(err)
				return true
			}
			stmt.Exec(u.Chat.ID, args[1], members, at, 0)
		} else {
			stmt, err := bm.DB.Prepare("UPDATE abi_gugugu SET members = $1, at = $2, status = $3 WHERE id = $4")
			if err != nil {
				log.Println(err)
				return true
			}
			stmt.Exec(members, at, 0, theGugugu.ID)
		}
		send(u, botmaid.Update{
			Message: &botmaid.Message{
				Text: fmt.Sprintf(random.String(formatGugugu), args[1]),
			},
			Chat: u.Chat,
		}, b, false)
		return true
	}
	return false
}

func guguguComplete(u *botmaid.Update, b *botmaid.Bot) bool {
	args := botmaid.SplitCommand(u.Message.Text)
	if b.IsCommand(u, "gugugu") && len(args) > 2 && slices.In(args[1], "-complete", "-c") {
		theGugugu := dbAbiGugugu{}
		err := bm.DB.QueryRow("SELECT * FROM abi_gugugu WHERE chat_id = $1 AND name = $2", u.Chat.ID, args[2]).Scan(&theGugugu.ID, &theGugugu.ChatID, &theGugugu.Name, &theGugugu.Members, &theGugugu.At, &theGugugu.Status)
		if err != nil || theGugugu.Status == 1 {
			return true
		}
		stmt, err := bm.DB.Prepare("UPDATE abi_gugugu SET status = $1 WHERE id = $2")
		if err != nil {
			return true
		}
		stmt.Exec(1, theGugugu.ID)
		send(u, botmaid.Update{
			Message: &botmaid.Message{
				Text: fmt.Sprintf(random.String(formatGuguguComplete), args[2]),
			},
			Chat: u.Chat,
		}, b, false)
		return true
	}
	return false
}

func guguguDelete(u *botmaid.Update, b *botmaid.Bot) bool {
	args := botmaid.SplitCommand(u.Message.Text)
	if b.IsCommand(u, "gugugu") && len(args) > 2 && slices.In(args[1], "-del", "-d") {
		theGugugu := dbAbiGugugu{}
		err := bm.DB.QueryRow("SELECT * FROM abi_gugugu WHERE chat_id = $1 AND name = $2", u.Chat.ID, args[2]).Scan(&theGugugu.ID, &theGugugu.ChatID, &theGugugu.Name, &theGugugu.Members, &theGugugu.At, &theGugugu.Status)
		if err != nil {
			return true
		}
		stmt, err := bm.DB.Prepare("DELETE FROM abi_gugugu WHERE id = $1")
		if err != nil {
			return true
		}
		stmt.Exec(theGugugu.ID)
		send(u, botmaid.Update{
			Message: &botmaid.Message{
				Text: fmt.Sprintf(random.String(formatGuguguDelete), args[2]),
			},
			Chat: u.Chat,
		}, b, false)
		return true
	}
	return false
}

func guguguList(u *botmaid.Update, b *botmaid.Bot) bool {
	args := botmaid.SplitCommand(u.Message.Text)
	if b.IsCommand(u, "gugugu") && len(args) > 1 && slices.In(args[1], "-list", "-l") {
		rows, err := bm.DB.Query("SELECT * FROM abi_gugugu WHERE chat_id = $1 AND status = $2", u.Chat.ID, 0)
		if err != nil {
			return true
		}
		if len(args) > 2 && slices.In(args[2], "-all", "-a") {
			rows, err = bm.DB.Query("SELECT * FROM abi_gugugu WHERE chat_id = $1", u.Chat.ID)
			if err != nil {
				return true
			}
		}
		defer rows.Close()
		list := []string{}
		for rows.Next() {
			theGugugu := dbAbiGugugu{}
			err := rows.Scan(&theGugugu.ID, &theGugugu.ChatID, &theGugugu.Name, &theGugugu.Members, &theGugugu.At, &theGugugu.Status)
			if err != nil {
				return true
			}
			text := theGugugu.Name + "（" + theGugugu.Members + "）"
			if theGugugu.Status == 0 {
				text = "[！]" + text
			} else {
				text = "[√]" + text
			}
			list = append(list, text)
		}
		if len(list) == 0 {
			send(u, botmaid.Update{
				Message: &botmaid.Message{
					Text: random.String(wordGuguguListEmpty),
				},
				Chat: u.Chat,
			}, b, false)
			return true
		}
		message := random.String(wordGuguguList)
		for _, v := range list {
			message += "\n" + v
		}
		send(u, botmaid.Update{
			Message: &botmaid.Message{
				Text: message,
			},
			Chat: u.Chat,
		}, b, false)
		return true
	}
	return false
}
