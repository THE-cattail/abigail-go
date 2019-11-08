package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/catsworld/botmaid"
)

func init() {
	var err error

	rootDir, _ = filepath.Abs(filepath.Dir(os.Args[0]))

	bm, err = botmaid.New(rootDir + "/config.toml")
	if err != nil {
		log.Fatalf("[Fatal] Create bot: %v\n", err)
	}

	bm.Words["selfIntro"] = []string{
		fmt.Sprintf("你好！我是阿比盖尔——阿比盖尔·威廉姆斯。我是For……eigner……%%v，如果不介意的话，叫我阿比吧。我们应该很快就能成为朋友。要叫出我的话，在命令前敲上%v就可以啦。", botmaid.ListToString(bm.Conf.CommandPrefix, "“%v”", "、", "或者")),
	}
}
