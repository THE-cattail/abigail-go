package main

import (
	"fmt"
	"log"
	"strconv"

	"github.com/catsworld/api"
	"github.com/catsworld/botmaid"
	"github.com/catsworld/random"
	"github.com/catsworld/slices"
)

type dbAbiRecord struct {
	ID      int64
	PlaceID int64
	Name    string
	Content string
}

var (
	wordRecordList = []string{
		"这是我记住的所有东西哦：\n",
	}
	wordRecordListEmpty = []string{
		"OvO我什么都不知道哦。",
	}
	formatRecordAdded = []string{
		"我记下了%s。",
	}
	formatRecordDeleted = []string{
		"%s被删除咯。",
	}
	formatRecordViewErr = []string{
		"%s？那是什么啊。",
	}
)

func init() {
	botmaid.AddCommand(&commands, recordView, 5)
	botmaid.AddCommand(&commands, recordDelete, 5)
	botmaid.AddCommand(&commands, recordList, 5)
	botmaid.AddCommand(&commands, record, 5)
}

func recordInit() {
	stmt, err := db.Prepare(`CREATE TABLE abi_record (
		id SERIAL primary key,
		chat_id bigint not null,
		name text,
		content text
	)`)
	if err != nil {
		log.Panic(err)
	}
	stmt.Exec()
}

func record(e *api.Event, b *botmaid.Bot) bool {
	args := botmaid.SplitCommand(e.Message.Text)
	if b.IsCommand(e, "record", "rec") && len(args) > 2 {
		theRecord := dbAbiRecord{}
		err := db.QueryRow("SELECT * FROM abi_record WHERE chat_id = $1 AND name = $2", e.Place.ID, args[1]).Scan(&theRecord.ID, &theRecord.PlaceID, &theRecord.Name, &theRecord.Content)
		if err != nil {
			stmt, err := db.Prepare("INSERT INTO abi_record(chat_id, name, content) VALUES($1, $2, $3)")
			if err != nil {
				return true
			}
			stmt.Exec(e.Place.ID, args[1], args[2])
		} else {
			stmt, err := db.Prepare("UPDATE abi_record SET content = $1 WHERE id = $2")
			if err != nil {
				return true
			}
			stmt.Exec(args[2], theRecord.ID)
		}
		send(&api.Event{
			Message: &api.Message{
				Text: fmt.Sprintf(random.String(formatRecordAdded), args[1]),
			},
			Place: e.Place,
		}, b, false)
		return true
	}
	return false
}

func recordView(e *api.Event, b *botmaid.Bot) bool {
	args := botmaid.SplitCommand(e.Message.Text)
	if b.IsCommand(e, "record", "rec") && len(args) == 2 {
		theRecord := dbAbiRecord{}
		err := db.QueryRow("SELECT * FROM abi_record WHERE chat_id = $1 AND name = $2", e.Place.ID, args[1]).Scan(&theRecord.ID, &theRecord.PlaceID, &theRecord.Name, &theRecord.Content)
		if err != nil {
			send(&api.Event{
				Message: &api.Message{
					Text: fmt.Sprintf(random.String(formatRecordViewErr), args[1]),
				},
				Place: e.Place,
			}, b, false)
			return true
		}
		send(&api.Event{
			Message: &api.Message{
				Text: args[1] + "：\n" + theRecord.Content,
			},
			Place: e.Place,
		}, b, false)
		return true
	}
	return false
}

func recordDelete(e *api.Event, b *botmaid.Bot) bool {
	args := botmaid.SplitCommand(e.Message.Text)
	if b.IsCommand(e, "record", "rec") && len(args) > 2 && slices.In(args[1], "--del", "-d") {
		theRecord := dbAbiRecord{}
		err := db.QueryRow("SELECT * FROM abi_record WHERE chat_id = $1 AND name = $2", e.Place.ID, args[2]).Scan(&theRecord.ID, &theRecord.PlaceID, &theRecord.Name, &theRecord.Content)
		if err != nil {
			send(&api.Event{
				Message: &api.Message{
					Text: fmt.Sprintf(random.String(formatRecordViewErr), args[2]),
				},
				Place: e.Place,
			}, b, false)
			return true
		}
		stmt, err := db.Prepare("DELETE FROM abi_record WHERE chat_id = $1 AND name = $2")
		if err != nil {
			return true
		}
		if len(args) > 3 && args[3] == "--id" {
			stmt, err = db.Prepare("DELETE FROM abi_record WHERE chat_id = $1 AND id = $2")
			if err != nil {
				return true
			}
		}
		stmt.Exec(e.Place.ID, args[2])
		send(&api.Event{
			Message: &api.Message{
				Text: fmt.Sprintf(random.String(formatRecordDeleted), args[2]),
			},
			Place: e.Place,
		}, b, false)
		return true
	}
	return false
}

func recordList(e *api.Event, b *botmaid.Bot) bool {
	args := botmaid.SplitCommand(e.Message.Text)
	if b.IsCommand(e, "record", "rec") && len(args) > 1 && slices.In(args[1], "--list", "-l") {
		rows, err := db.Query("SELECT * FROM abi_record WHERE chat_id = $1", e.Place.ID)
		if err != nil {
			return true
		}
		defer rows.Close()
		list := []string{}
		for rows.Next() {
			theRecord := dbAbiRecord{}
			err := rows.Scan(&theRecord.ID, &theRecord.PlaceID, &theRecord.Name, &theRecord.Content)
			if err != nil {
				return true
			}
			if len(args) > 2 && args[2] == "--id" {
				list = append(list, theRecord.Name+"["+strconv.FormatInt(theRecord.ID, 10)+"]")
			} else {
				list = append(list, theRecord.Name)
			}
		}
		if len(list) == 0 {
			send(&api.Event{
				Message: &api.Message{
					Text: random.String(wordRecordListEmpty),
				},
				Place: e.Place,
			}, b, false)
			return true
		}
		message := random.String(wordRecordList)
		for i := 0; i < len(list)-1; i++ {
			message += list[i] + "、"
		}
		message += list[len(list)-1] + "。"
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
