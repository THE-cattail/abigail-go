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

	bm.Words = map[string][]string{
		"selfIntro": []string{
			fmt.Sprintf("我是阿比盖尔，他们都叫我塞勒姆的魔女呢，呵呵，可别把我惹急了。%%v，要叫出我的话，在命令前敲上%v就可以了。", botmaid.ListToString(bm.Conf.CommandPrefix, "“%v”", "、", "或者")),
			fmt.Sprintf("阿比可不是什么坏孩子哦，嘻嘻，%%v想要呼唤阿比的话，在命令前敲上%v就行啦。", botmaid.ListToString(bm.Conf.CommandPrefix, "“%v”", "、", "或者")),
		},
		"undefCommand": []string{
			"%v？那是什么啊，乖孩子是不知道的呢。",
			"%v？抱歉哦，阿比不知道呢OvO。",
		},
		"invalidMaster": []string{
			"要用 At 和我说哦，不然我也没法知道%v的用户名呢。",
		},
		"masterExisted": []string{
			"%v已经是我的御主了。",
		},
		"masterAdded": []string{
			"%v，你好！我是阿比盖尔——阿比盖尔·威廉姆斯，我是Fo……reigner……你就是御主吗？如果你不介意的话，希望你能叫我阿比。我想我们很快就能成为朋友。",
		},
		"masterNotExisted": []string{
			"%v不是我的御主呢。",
		},
		"masterRemoved": []string{
			"以后%v就不是我的御主了哦。",
		},
	}
}
