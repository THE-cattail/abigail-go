package main

import (
	"fmt"
	"log"

	"github.com/catsworld/api"
	"github.com/catsworld/botmaid"
	"github.com/catsworld/random"
	"github.com/catsworld/slices"
)

type dbAbiGugugu struct {
	ID      int64
	PlaceID int64
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
	botmaid.AddCommand(&commands, guguguList, 5)
	botmaid.AddCommand(&commands, guguguComplete, 5)
	botmaid.AddCommand(&commands, guguguDelete, 5)
	botmaid.AddCommand(&commands, gugugu, 5)
}

func guguguInit() {
	stmt, err := db.Prepare(`CREATE TABLE abi_gugugu (
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

func gugugu(e *api.Event, b *botmaid.Bot) bool {
	args := botmaid.SplitCommand(e.Message.Text)
	if b.IsCommand(e, "gugugu") && len(args) > 1 {
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
		err := db.QueryRow("SELECT * FROM abi_gugugu WHERE chat_id = $1 AND name = $2", e.Place.ID, args[1]).Scan(&theGugugu.ID, &theGugugu.PlaceID, &theGugugu.Name, &theGugugu.Members, &theGugugu.At, &theGugugu.Status)
		if err != nil {
			stmt, err := db.Prepare("INSERT INTO abi_gugugu(chat_id, name, members, at, status) VALUES($1, $2, $3, $4, $5)")
			if err != nil {
				log.Println(err)
				return true
			}
			stmt.Exec(e.Place.ID, args[1], members, at, 0)
		} else {
			stmt, err := db.Prepare("UPDATE abi_gugugu SET members = $1, at = $2, status = $3 WHERE id = $4")
			if err != nil {
				log.Println(err)
				return true
			}
			stmt.Exec(members, at, 0, theGugugu.ID)
		}
		send(&api.Event{
			Message: &api.Message{
				Text: fmt.Sprintf(random.String(formatGugugu), args[1]),
			},
			Place: e.Place,
		}, b, false)
		return true
	}
	return false
}

func guguguComplete(e *api.Event, b *botmaid.Bot) bool {
	args := botmaid.SplitCommand(e.Message.Text)
	if b.IsCommand(e, "gugugu") && len(args) > 2 && slices.In(args[1], "-complete", "-c") {
		theGugugu := dbAbiGugugu{}
		err := db.QueryRow("SELECT * FROM abi_gugugu WHERE chat_id = $1 AND name = $2", e.Place.ID, args[2]).Scan(&theGugugu.ID, &theGugugu.PlaceID, &theGugugu.Name, &theGugugu.Members, &theGugugu.At, &theGugugu.Status)
		if err != nil || theGugugu.Status == 1 {
			return true
		}
		stmt, err := db.Prepare("UPDATE abi_gugugu SET status = $1 WHERE id = $2")
		if err != nil {
			return true
		}
		stmt.Exec(1, theGugugu.ID)
		send(&api.Event{
			Message: &api.Message{
				Text: fmt.Sprintf(random.String(formatGuguguComplete), args[2]),
			},
			Place: e.Place,
		}, b, false)
		return true
	}
	return false
}

func guguguDelete(e *api.Event, b *botmaid.Bot) bool {
	args := botmaid.SplitCommand(e.Message.Text)
	if b.IsCommand(e, "gugugu") && len(args) > 2 && slices.In(args[1], "-del", "-d") {
		theGugugu := dbAbiGugugu{}
		err := db.QueryRow("SELECT * FROM abi_gugugu WHERE chat_id = $1 AND name = $2", e.Place.ID, args[2]).Scan(&theGugugu.ID, &theGugugu.PlaceID, &theGugugu.Name, &theGugugu.Members, &theGugugu.At, &theGugugu.Status)
		if err != nil {
			return true
		}
		stmt, err := db.Prepare("DELETE FROM abi_gugugu WHERE id = $1")
		if err != nil {
			return true
		}
		stmt.Exec(theGugugu.ID)
		send(&api.Event{
			Message: &api.Message{
				Text: fmt.Sprintf(random.String(formatGuguguDelete), args[2]),
			},
			Place: e.Place,
		}, b, false)
		return true
	}
	return false
}

func guguguList(e *api.Event, b *botmaid.Bot) bool {
	args := botmaid.SplitCommand(e.Message.Text)
	if b.IsCommand(e, "gugugu") && len(args) > 1 && slices.In(args[1], "-list", "-l") {
		rows, err := db.Query("SELECT * FROM abi_gugugu WHERE chat_id = $1 AND status = $2", e.Place.ID, 0)
		if err != nil {
			return true
		}
		if len(args) > 2 && slices.In(args[2], "-all", "-a") {
			rows, err = db.Query("SELECT * FROM abi_gugugu WHERE chat_id = $1", e.Place.ID)
			if err != nil {
				return true
			}
		}
		defer rows.Close()
		list := []string{}
		for rows.Next() {
			theGugugu := dbAbiGugugu{}
			err := rows.Scan(&theGugugu.ID, &theGugugu.PlaceID, &theGugugu.Name, &theGugugu.Members, &theGugugu.At, &theGugugu.Status)
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
			send(&api.Event{
				Message: &api.Message{
					Text: random.String(wordGuguguListEmpty),
				},
				Place: e.Place,
			}, b, false)
			return true
		}
		message := random.String(wordGuguguList)
		for _, v := range list {
			message += "\n" + v
		}
		send(&api.Event{
			Message: &api.Message{
				Text: message,
			},
			Place: e.Place,
		}, b, false)
		return true
	}
	return false
}
