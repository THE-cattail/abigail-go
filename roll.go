package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/catsworld/abigail/coc"
	"github.com/catsworld/botmaid"
	"github.com/catsworld/nyamath"
	"github.com/catsworld/random"
)

var (
	formatRoll = []string{
		"%s的%s是——%v！",
	}

	hide = map[int64]bool{}
)

func init() {
	bm.AddCommand(botmaid.Command{
		Do: func(u *botmaid.Update, b *botmaid.Bot) bool {
			if botmaid.In(u.Message.Args[1], "final") {
				send(b, botmaid.Update{
					Message: &botmaid.Message{
						Text: fmt.Sprintf(random.String([]string{
							"%s，汝已陷入了疯狂：\n",
						}), u.User.NickName) + coc.RollMadFinal(),
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
		Help:       " final - roll 一次疯狂的总结症状",
	})
	bm.AddCommand(botmaid.Command{
		Do: func(u *botmaid.Update, b *botmaid.Bot) bool {
			send(b, botmaid.Update{
				Message: &botmaid.Message{
					Text: fmt.Sprintf(random.String([]string{
						"%s，汝已陷入了疯狂：\n",
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
	bm.AddCommand(botmaid.Command{
		Do: func(u *botmaid.Update, b *botmaid.Bot) bool {
			c := coc.RollCharacter()
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
			if len(u.Message.Args) > 1 && botmaid.In(u.Message.Args[1], "full") {
				message += "\n"
				/*
					message += random.String([]string{
						"你是这样的：",
					}) + c.Description + "\n"
				*/
				message += random.String([]string{
					"你的信念是这样的：",
				}) + c.Thought + "\n"
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
			send(b, botmaid.Update{
				Message: &botmaid.Message{
					Text: fmt.Sprintf(random.String([]string{
						"这就是%s新的调查员哦。\n",
						"%s家新来的调查员？不过如此，呵呵呵。\n",
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
	bm.AddCommand(botmaid.Command{
		Do: func(u *botmaid.Update, b *botmaid.Bot) bool {
			if botmaid.In(u.Message.Args[1], "hide") {
				hide[u.ID] = true
				u.Message.Text = strings.Replace(u.Message.Text, "hide", "", 1)
			}
			return false
		},
		Names:      []string{"roll", "r", "ww"},
		ArgsMinLen: 1,
		Help:       " hide - 进行一次暗骰",
	})
	bm.AddCommand(botmaid.Command{
		Do: func(u *botmaid.Update, b *botmaid.Bot) bool {
			instruction := ""
			exp := ""
			if b.IsCommand(u, []string{"roll", "r"}) && len(u.Message.Args) == 2 {
				exp = u.Message.Args[1]
			}
			if b.IsCommand(u, []string{"roll", "r"}) && len(u.Message.Args) == 3 {
				instruction = u.Message.Args[1]
				exp = u.Message.Args[2]
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
			if a > 256 {
				a = 256
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
			send(b, botmaid.Update{
				Message: &botmaid.Message{
					Text: fmt.Sprintf(random.String(formatRoll), u.User.NickName, instruction+"检定结果", result),
				},
				Chat: u.Chat,
			}, hide[u.ID], u)
			return true
		},
	})
	bm.AddCommand(botmaid.Command{
		Do: func(u *botmaid.Update, b *botmaid.Bot) bool {
			instruction := "计算结果"
			exp := ""
			if b.IsCommand(u, []string{"roll", "r"}) && len(u.Message.Args) == 2 {
				exp = u.Message.Args[1]
			}
			if b.IsCommand(u, []string{"roll", "r"}) && len(u.Message.Args) == 3 {
				instruction = u.Message.Args[1]
				exp = u.Message.Args[2]
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
			send(b, botmaid.Update{
				Message: &botmaid.Message{
					Text: fmt.Sprintf(random.String(formatRoll), u.User.NickName, instruction, ee.Result()),
				},
				Chat: u.Chat,
			}, hide[u.ID], u)
			return true
		},
	})
	bm.AddCommand(botmaid.Command{
		Do: func(u *botmaid.Update, b *botmaid.Bot) bool {
			instruction := ""
			num := ""
			if b.IsCommand(u, []string{"roll", "r"}) && len(u.Message.Args) == 1 {
				num = "0"
			}
			if b.IsCommand(u, []string{"roll", "r"}) && len(u.Message.Args) == 2 {
				_, err := strconv.Atoi(u.Message.Args[1])
				if err != nil {
					instruction = u.Message.Args[1]
					num = "0"
				} else {
					num = u.Message.Args[1]
				}
			}
			if b.IsCommand(u, []string{"roll", "r"}) && len(u.Message.Args) == 3 {
				instruction = u.Message.Args[1]
				num = u.Message.Args[2]
			}
			n, err := strconv.Atoi(num)
			if err != nil {
				return false
			}
			result := coc.Check(n)
			message := strconv.Itoa(result.Number)
			if n > 0 && n < 100 {
				message = strconv.Itoa(n) + "/" + message
			}
			if result.BigSuccess() {
				message += "！" + random.String([]string{
					"大成功",
					"☆大☆成☆功☆",
				})
			} else if result.BigFailure() {
				message += "！" + random.String([]string{
					"大失败",
					"大失败 Q皿Q",
				})
			} else if result.HardSuccess() {
				message += "！" + random.String([]string{
					"困难成功",
				})
			} else if result.VeryHardSuccess() {
				message += "！" + random.String([]string{
					"极难成功",
				})
			} else if result.Success() {
				message += "！" + random.String([]string{
					"检定成功",
				})
			} else if result.Failure() {
				message += "！" + random.String([]string{
					"检定失败",
				})
			}
			message = fmt.Sprintf(random.String(formatRoll), u.User.NickName, instruction+"检定结果", message)
			if result.Success() {
				message += botmaid.WordSlice{
					botmaid.Word{
						Word:   "",
						Weight: 57,
					},
					botmaid.Word{
						Word:   "(๑•̀ㅂ•́)✧",
						Weight: 1,
					},
					botmaid.Word{
						Word:   "qwq",
						Weight: 1,
					},
					botmaid.Word{
						Word:   "(✪ω✪)",
						Weight: 1,
					},
				}.Random()
			}
			send(b, botmaid.Update{
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
				pkResp(u, b)
			}
			if _, ok := scMap[u.Chat.ID]; ok {
				if _, ok := scMap[u.Chat.ID][u.User.ID]; ok && scMap[u.Chat.ID][u.User.ID].Status && n > 0 && n < 100 {
					scMap[u.Chat.ID][u.User.ID].Result = result
					scResp(u, b)
				}
			}
			return true
		},
		Menu:     "roll",
		MenuText: "roll 点",
		Names:    []string{"roll", "r"},
		Help:     " <说明（可省略）> <表达式（可省略）> - 进行一次表达式计算/检定",
	})
	bm.AddCommand(botmaid.Command{
		Do: func(u *botmaid.Update, b *botmaid.Bot) bool {
			send(b, botmaid.Update{
				Message: &botmaid.Message{
					Text: fmt.Sprintf(random.String(formatRoll), u.User.NickName, "随机结果", u.Message.Args[random.Int(1, len(u.Message.Args)-1)]),
				},
				Chat: u.Chat,
			}, hide[u.ID], u)
			return true
		},
		Menu:  "roll",
		Names: []string{"roll", "r"},
		Help:  " <列表> - 对列表进行随机抽取",
	})
	bm.AddCommand(botmaid.Command{
		Do: func(u *botmaid.Update, b *botmaid.Bot) bool {
			instruction := ""
			exp := ""
			if b.IsCommand(u, []string{"ww"}) && len(u.Message.Args) == 2 {
				exp = u.Message.Args[1]
			}
			if b.IsCommand(u, []string{"ww"}) && len(u.Message.Args) == 3 {
				instruction = u.Message.Args[1]
				exp = u.Message.Args[2]
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
			send(b, botmaid.Update{
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
