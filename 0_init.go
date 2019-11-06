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
		fmt.Sprintf("我是阿比盖尔，他们都叫我塞勒姆的魔女呢，呵呵，可别把我惹急了。%%v，要叫出我的话，在命令前敲上%v就可以了。", botmaid.ListToString(bm.Conf.CommandPrefix, "“%v”", "、", "或者")),
		fmt.Sprintf("阿比可不是什么坏孩子哦，嘻嘻，%%v想要呼唤阿比的话，在命令前敲上%v就行啦。", botmaid.ListToString(bm.Conf.CommandPrefix, "“%v”", "、", "或者")),
	}
	bm.Words["undefCommand"] = []string{
		"%v？那是什么啊，乖孩子是不知道的呢。",
		"%v？抱歉哦，阿比不知道呢OvO。",
	}
}
