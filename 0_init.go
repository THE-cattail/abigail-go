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

	bm.Words["selfIntro"] = fmt.Sprintf(`你好！我是%%v——阿比盖尔·威廉姆斯。我是For……eigner……，如果不介意的话，叫我阿比吧。我们应该很快就能成为朋友。

使用方法：

%v（%v）*命令* [参数]
		
命令目录：
%%v
		
使用 ”help [命令]“ 来获得关于单条命令的更多信息。

本程序按照《Call of Cthulhu 7th》秋叶EXODUS翻译版 Version1902 编写`, bm.Conf.CommandPrefix[0], botmaid.ListToString(bm.Conf.CommandPrefix[1:], "“%v”", "、", "或者"))
	bm.Words["undefCommand"] = "%v，命令“%v”不存在，请检查拼写或该命令的帮助条目后重试。"
	bm.Words["unregMaster"] = "%v，%v的 master 权限已被解除。"
	bm.Words["regMaster"] = "%v，%v已获得 master 权限。"
	bm.Words["unbanUser"] = "%v，已解除%v的屏蔽。"
	bm.Words["banUser"] = "%v，%v已被屏蔽。"
	bm.Words["noPermission"] = "%v，你没有使用命令“%v”的权限。"
	bm.Words["invalidParameters"] = "%v，命令“%v”的参数格式输入错误。"
	bm.Words["noHelpText"] = "%v，命令“%v”没有帮助文本。"
	bm.Words["invalidUser"] = "%v，用户“%v”格式非法或不存在。"
	bm.Words["helpHelp"] = "显示帮助"
	bm.Words["masterHelp"] = "增加/移除 master"
	bm.Words["banHelp"] = "屏蔽/解除用户"
	bm.Words["helpHelp"] = "显示帮助"
	bm.Words["helpHelpFull"] = `使用方法：help [命令]

%v`
	bm.Words["masterHelp"] = "增加/移除 master"
	bm.Words["masterHelpFull"] = `使用方法：master @用户

%v`
	bm.Words["banHelp"] = "屏蔽/解除用户"
	bm.Words["banHelpFull"] = `使用方法：ban @用户

%v`
}
