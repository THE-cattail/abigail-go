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
		fmt.Sprintf(`你好！我是%%v——阿比盖尔·威廉姆斯。我是For……eigner……，如果不介意的话，叫我阿比吧。我们应该很快就能成为朋友。

使用方法：

%v（%v）*命令* [参数]
		
命令目录：
%%v
		
使用 ”help [命令]“ 来获得关于单条命令的更多信息。

本程序按照《Call of Cthulhu 7th》秋叶EXODUS翻译版 Version1902 编写`, bm.Conf.CommandPrefix[0], botmaid.ListToString(bm.Conf.CommandPrefix[1:], "“%v”", "、", "或者")),
	}
}
