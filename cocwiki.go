package main

import (
	"github.com/catsworld/api"
	"github.com/catsworld/botmaid"
	"github.com/catsworld/coc"
)

func init() {
	bm.AddCommand(botmaid.Command{
		Do:       cocWiki,
		Priority: 5,
		Menu:     "cocwiki",
		Names:    []string{"cocwiki"},
		Help:     " <词条> - 在 CoC 百科中查询资料",
	})
}

func cocWiki(e *api.Event, b *botmaid.Bot) bool {
	args := botmaid.SplitCommand(e.Message.Text)
	if b.IsCommand(e, "cocwiki") && len(args) > 1 {
		send(api.Event{
			Message: &api.Message{
				Text: coc.Wiki(args[1]),
			},
			Place: e.Place,
		}, b, false)
		return true
	}
	return false
}
