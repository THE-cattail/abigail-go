package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/the-cattail/botmaid"
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
	bm.Words["noPermission"] = "%v，你没有使用命令“%v”的权限。"
	bm.Words["invalidParameters"] = "%v，命令“%v”的参数格式输入错误。"
	bm.Words["noHelpText"] = "%v，命令“%v”没有帮助文本。"
	bm.Words["invalidUser"] = "%v，用户“%v”格式非法或不存在。"
	bm.Words["fmtVersion"] = "小阿比 %v"
	bm.Words["fmtLog"] = "小阿比 %v：\n\n更新日志：%v"
	bm.Words["versionSet"] = "当前版本号已设置为 %v."
	bm.Words["logAdded"] = "已添加更新日志条目\"%v\"。"
	bm.Words["versionLogHelp"] = "显示当前版本的更新日志"
	bm.Words["versetVerHelp"] = "指定管理的版本"
	bm.Words["versetLogHelp"] = "向更新日志添加一项"
	bm.Words["versetBroadcastHelp"] = "广播更新日志"
	bm.Words["upgraded"] = "版本更新！"
	bm.Words["subscribed"] = "已订阅“%v”"
	bm.Words["unsubscribed"] = "已取消订阅“%v”"
	bm.Words["correctSubEntries"] = "这些条目可以被订阅：%v"
	bm.Words["subEntriesFormat"] = "“%v”"
	bm.Words["subEntriesSeparator"] = "、"
	bm.Words["subEntriesAnd"] = "和"
}
