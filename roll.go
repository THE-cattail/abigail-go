package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/catsworld/abigail/coc"
	"github.com/catsworld/abigail/nyamath"
	"github.com/catsworld/botmaid"
	"github.com/catsworld/botmaid/random"
	"github.com/spf13/pflag"
)

var (
	formatRoll = []string{
		"%v的%v是——%v！",
	}
)

// TODO: hide argument in roll
func init() {
	bm.AddCommand(&botmaid.Command{
		Do: func(u *botmaid.Update, f *pflag.FlagSet) bool {
			if botmaid.In(f.Args()[1], "summary", "sum") {
				send(&botmaid.Update{
					Message: &botmaid.Message{
						Text: fmt.Sprintf(random.String([]string{
							"%v，汝已陷入了临时疯狂：\n",
						}), u.User.NickName) + coc.RollMadSummary(),
					},
					Chat: u.Chat,
				}, hide[u.ID], u)
				return true
			}
			return false
		},
		Menu:       "roll",
		Names:      []string{"mad"},
		ArgsMinLen: 2,
		ArgsMaxLen: 2,
		Help:       " summary - roll 一次疯狂的总结症状",
	})
	bm.AddCommand(&botmaid.Command{
		Do: func(u *botmaid.Update, f *pflag.FlagSet) bool {
			send(&botmaid.Update{
				Message: &botmaid.Message{
					Text: fmt.Sprintf(random.String([]string{
						"%v，汝已陷入了临时疯狂：\n",
					}), u.User.NickName) + coc.RollMad(),
				},
				Chat: u.Chat,
			}, hide[u.ID], u)
			return true
		},
		Menu:       "roll",
		Names:      []string{"mad"},
		ArgsMinLen: 1,
		ArgsMaxLen: 1,
		Help:       " - roll 一次疯狂的即时症状",
	})
	bm.AddCommand(&botmaid.Command{
		Do: func(u *botmaid.Update, f *pflag.FlagSet) bool {
			c := coc.NewCharacter()
			message := random.String([]string{
				"这是你的属性：",
			}) + "\n"
			message += "力量" + strconv.Itoa(c.STR) + " "
			message += "敏捷" + strconv.Itoa(c.DEX) + " "
			message += "意志" + strconv.Itoa(c.POW) + " "
			message += "体质" + strconv.Itoa(c.CON) + " "
			message += "外貌" + strconv.Itoa(c.APP) + " "
			message += "教育" + strconv.Itoa(c.EDU) + " "
			message += "体型" + strconv.Itoa(c.SIZ) + " "
			message += "智力" + strconv.Itoa(c.INT) + " "
			message += "幸运" + strconv.Itoa(c.Luck)
			if len(f.Args()) > 1 && botmaid.In(f.Args()[1], "full") {
				message += "\n"
				/*
					message += random.String([]string{
						"你是这样的：",
					}) + c.Description + "\n"
				*/
				message += c.Thought + "\n"
				message += random.String([]string{
					"你的重要之人是",
				}) + c.Person + random.String([]string{
					"因为",
				}) + c.Reason + "\n"
				message += random.String([]string{
					"你的意义非凡之地是",
				}) + c.Place + "\n"
				message += random.String([]string{
					"你的宝贵之物是",
				}) + c.Treasure + "\n"
				message += random.String([]string{
					"你是一个",
				}) + c.Feature
			}
			send(&botmaid.Update{
				Message: &botmaid.Message{
					Text: fmt.Sprintf(random.String([]string{
						"这就是%v新的调查员哦。\n",
						"%v家新来的调查员？不过如此，呵呵呵。\n",
					}), u.User.NickName) + message,
				},
				Chat: u.Chat,
			}, hide[u.ID], u)
			return true
		},
		Menu:       "roll",
		Names:      []string{"char", "character"},
		ArgsMinLen: 1,
		ArgsMaxLen: 2,
		Help:       " (full) - roll一张人物卡",
	})
	bm.AddCommand(&botmaid.Command{
		Do: func(u *botmaid.Update, f *pflag.FlagSet) bool {
			if botmaid.In(f.Args()[1], "hide") {
				hide[u.ID] = true
				u.Message.Text = strings.Replace(u.Message.Text, "hide", "", 1)
			}
			return false
		},
		Names:      []string{"roll", "r", "ww"},
		ArgsMinLen: 2,
		Help:       " hide - 进行一次暗骰",
	})
	bm.AddCommand(&botmaid.Command{
		Do: func(u *botmaid.Update, f *pflag.FlagSet) bool {
			instruction := ""
			exp := ""
			if botmaid.IsCommand(u, []string{"roll", "r"}) && len(f.Args()) == 2 {
				exp = f.Args()[1]
			}
			if botmaid.IsCommand(u, []string{"roll", "r"}) && len(f.Args()) == 3 {
				instruction = f.Args()[1]
				exp = f.Args()[2]
			}
			if !strings.Contains(exp, "d") {
				return false
			}
			ee := strings.Split(exp, "d")
			if len(ee) > 0 && ee[0] == "" {
				ee[0] = "1"
			}
			if len(ee) != 2 {
				return false
			}
			a, err := strconv.Atoi(ee[0])
			if err != nil {
				return false
			}
			bb, err := strconv.Atoi(ee[1])
			if err != nil {
				return false
			}
			if a > 1000000 {
				a = 1000000
			}
			result := ""
			if a == 1 {
				result = strconv.Itoa(random.Int(1, bb))
			} else {
				result = "("
				sum := 0
				for i := 0; i < a; i++ {
					t := random.Int(1, bb)
					result += strconv.Itoa(t)
					sum += t
					if i < a-1 {
						result += ", "
					}
				}
				result += ") = " + strconv.Itoa(sum)
			}
			send(&botmaid.Update{
				Message: &botmaid.Message{
					Text: fmt.Sprintf(random.String(formatRoll), u.User.NickName, instruction+"检定结果", result),
				},
				Chat: u.Chat,
			}, hide[u.ID], u)
			return true
		},
	})
	bm.AddCommand(&botmaid.Command{
		Do: func(u *botmaid.Update, f *pflag.FlagSet) bool {
			instruction := "计算结果"
			exp := ""
			if botmaid.IsCommand(u, []string{"roll", "r"}) && len(f.Args()) == 2 {
				exp = f.Args()[1]
			}
			if botmaid.IsCommand(u, []string{"roll", "r"}) && len(f.Args()) == 3 {
				instruction = f.Args()[1]
				exp = f.Args()[2]
			}
			if exp == "" {
				return false
			}
			_, err := strconv.Atoi(exp)
			if err == nil {
				return false
			}
			ee, err := nyamath.New(exp)
			if err != nil {
				return false
			}
			send(&botmaid.Update{
				Message: &botmaid.Message{
					Text: fmt.Sprintf(random.String(formatRoll), u.User.NickName, instruction, ee.Result.Value),
				},
				Chat: u.Chat,
			}, hide[u.ID], u)
			return true
		},
	})
	bm.AddCommand(&botmaid.Command{
		Do: func(u *botmaid.Update, f *pflag.FlagSet) bool {
			instruction := ""
			num := ""
			if botmaid.IsCommand(u, []string{"roll", "r"}) && len(f.Args()) == 1 {
				num = "0"
			}
			if botmaid.IsCommand(u, []string{"roll", "r"}) && len(f.Args()) == 2 {
				_, err := strconv.Atoi(f.Args()[1])
				if err != nil {
					instruction = f.Args()[1]
					num = "0"
				} else {
					num = f.Args()[1]
				}
			}
			if botmaid.IsCommand(u, []string{"roll", "r"}) && len(f.Args()) == 3 {
				instruction = f.Args()[1]
				num = f.Args()[2]
			}
			n, err := strconv.Atoi(num)
			if err != nil {
				return false
			}
			result := coc.Check(n)
			message := strconv.Itoa(result.Number)
			if n > 0 {
				message = strconv.Itoa(n) + "/" + message
				if result.Great == coc.GreatSucc {
					message += "！" + random.String([]string{
						"大成功",
						"☆大☆成☆功☆",
					})
				} else if result.Great == coc.GreatFail {
					message += "！" + random.String([]string{
						"大失败",
						"大失败 Q皿Q",
					})
				} else if result.Level == coc.DiffSucc {
					message += "！" + random.String([]string{
						"困难成功",
					})
				} else if result.Level == coc.ExDiffSucc {
					message += "！" + random.String([]string{
						"极难成功",
					})
				} else if result.Succ == coc.Succ {
					message += "！" + random.String([]string{
						"检定成功",
					})
				} else if result.Succ == coc.Fail {
					message += "！" + random.String([]string{
						"检定失败",
					})
				}
			}
			message = fmt.Sprintf(random.String(formatRoll), u.User.NickName, instruction+"检定结果", message)
			if result.Succ == coc.Succ {
				message += random.WordSlice{
					random.Word{
						Word:   "",
						Weight: 57,
					},
					random.Word{
						Word:   "(๑•̀ㅂ•́)✧",
						Weight: 1,
					},
					random.Word{
						Word:   "qwq",
						Weight: 1,
					},
					random.Word{
						Word:   "(✪ω✪)",
						Weight: 1,
					},
				}.Random()
			}
			send(&botmaid.Update{
				Message: &botmaid.Message{
					Text: message,
				},
				Chat: u.Chat,
			}, hide[u.ID], u)
			if _, ok := pkMap[u.Chat.ID]; ok && pkMap[u.Chat.ID].Status && n > 0 && n < 100 {
				pkMap[u.Chat.ID].Results = append(pkMap[u.Chat.ID].Results, pkRollResult{
					User:   u.User,
					Result: result,
				})
				pkResp(u)
			}
			if _, ok := scMap[u.Chat.ID]; ok {
				log.Println(scMap[u.Chat.ID][u.User.ID])
				if _, ok := scMap[u.Chat.ID][u.User.ID]; ok && scMap[u.Chat.ID][u.User.ID].Status && n > 0 && n < 100 {
					scMap[u.Chat.ID][u.User.ID].Result = result
					scResp(u)
				}
			}
			return true
		},
		Menu:     "roll",
		MenuText: "roll 点",
		Names:    []string{"roll", "r"},
		Help:     " <说明（可省略）> <表达式（可省略）> - 进行一次表达式计算/检定",
	})
	bm.AddCommand(&botmaid.Command{
		Do: func(u *botmaid.Update, f *pflag.FlagSet) bool {
			send(&botmaid.Update{
				Message: &botmaid.Message{
					Text: fmt.Sprintf(random.String(formatRoll), u.User.NickName, "随机结果", f.Args()[random.Int(1, len(f.Args())-1)]),
				},
				Chat: u.Chat,
			}, hide[u.ID], u)
			return true
		},
		Menu:  "roll",
		Names: []string{"roll", "r"},
		Help:  " <列表> - 对列表进行随机抽取",
	})
	bm.AddCommand(&botmaid.Command{
		Do: func(u *botmaid.Update, f *pflag.FlagSet) bool {
			instruction := ""
			exp := ""
			if botmaid.IsCommand(u, []string{"ww"}) && len(f.Args()) == 2 {
				exp = f.Args()[1]
			}
			if botmaid.IsCommand(u, []string{"ww"}) && len(f.Args()) == 3 {
				instruction = f.Args()[1]
				exp = f.Args()[2]
			}
			plus := 0
			a := 10
			if strings.Contains(exp, "+") {
				args := strings.Split(exp, "+")
				if len(args) != 2 {
					return true
				}
				exp = args[0]
				tmp, err := strconv.Atoi(args[1])
				if err != nil {
					return true
				}
				plus = tmp
			}
			if strings.Contains(exp, "a") {
				args := strings.Split(exp, "a")
				if len(args) != 2 {
					return true
				}
				exp = args[0]
				tmp, err := strconv.Atoi(args[1])
				if err != nil {
					return true
				}
				a = tmp
			}
			num, err := strconv.Atoi(exp)
			if err != nil {
				return true
			}
			if num < 1 {
				return true
			}
			if num > 256 {
				num = 256
			}
			rollbotmaidlice := [][]int{}
			ans := 0
			remain := num
			if a < 8 {
				a = 8
			}
			for remain > 0 {
				num = remain
				remain = 0
				rollSlice := []int{}
				for i := 0; i < num; i++ {
					result := random.Int(1, 10)
					rollSlice = append(rollSlice, result)
					if result > 7 {
						ans++
					}
					if result >= a {
						remain++
					}
				}
				rollbotmaidlice = append(rollbotmaidlice, rollSlice)
			}
			message := ""
			for i, rollSlice := range rollbotmaidlice {
				if i != 0 {
					message += " + "
				}
				message += "("
				for i, v := range rollSlice {
					if i != 0 {
						message += ", "
					}
					message += strconv.Itoa(v)
				}
				message += ")"
			}
			if plus > 0 {
				message += " + " + strconv.Itoa(plus) + " = " + strconv.Itoa(ans) + " + " + strconv.Itoa(plus)
			}
			message += " = " + strconv.Itoa(ans+plus)
			send(&botmaid.Update{
				Message: &botmaid.Message{
					Text: fmt.Sprintf(random.String(formatRoll), u.User.NickName, instruction+"检定结果", message),
				},
				Chat: u.Chat,
			}, hide[u.ID], u)
			return true
		},
		Menu:       "roll",
		Names:      []string{"ww"},
		ArgsMinLen: 2,
		ArgsMaxLen: 2,
		Help:       " <数值> - 无限恐怖规则中的骰点",
	})
}
