package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/catsworld/abigail/coc"
	"github.com/catsworld/abigail/nyamath"
	"github.com/catsworld/botmaid"
	"github.com/catsworld/botmaid/random"
	"github.com/spf13/pflag"
)

func init() {
	bm.AddCommand(&botmaid.Command{
		Do: func(u *botmaid.Update, f *pflag.FlagSet) bool {
			sum, _ := f.GetBool("sum")
			if sum {
				reply(u, fmt.Sprintf("%v陷入了临时疯狂，其总结症状为：\n%v", bm.At(u.User), coc.RollMadSummary()))
				return true
			}

			reply(u, fmt.Sprintf("%v陷入了临时疯狂，其即时症状为：\n%v", bm.At(u.User), coc.RollMad()))
			return true
		},
		Help: &botmaid.Help{
			Menu:  "mad",
			Help:  "随机生成疯狂症状",
			Names: []string{"mad"},
			Usage: "使用方法：mad [选项]，默认生成即时症状",
			SetFlag: func(f *pflag.FlagSet) {
				f.BoolP("sum", "s", false, "生成总结症状")
			},
		},
	})

	bm.AddCommand(&botmaid.Command{
		Do: func(u *botmaid.Update, f *pflag.FlagSet) bool {
			full, _ := f.GetBool("full")
			num, _ := f.GetInt("num")
			if full {
				num = 1
			}
			if num > 5 {
				bm.Reply(u, "同时生成的角色卡数量至多为 5。")
			}

			c := &coc.Character{}
			s := ""

			if full {
				s += "属性：\n"
			}

			for i := 0; i < num; i++ {
				c = coc.NewCharacter(full)

				if i != 0 {
					s += "\n"
				}

				s += "力量" + strconv.Itoa(c.STR) + " "
				s += "敏捷" + strconv.Itoa(c.DEX) + " "
				s += "意志" + strconv.Itoa(c.POW) + " "
				s += "体质" + strconv.Itoa(c.CON) + " "
				s += "外貌" + strconv.Itoa(c.APP) + " "
				s += "教育" + strconv.Itoa(c.EDU) + " "
				s += "体型" + strconv.Itoa(c.SIZ) + " "
				s += "智力" + strconv.Itoa(c.INT) + " "
				s += "幸运" + strconv.Itoa(c.Luck)
			}

			if full {
				s += fmt.Sprintf("\n调查员姓名：%v 年龄：%v 职业：%v", c.Name, c.Age, c.Class)

				if c.AgeMention != "" {
					s += "\n（" + c.AgeMention + "）"
				}

				s += "\n背景：\n"
				/*
					s += "你是这样的：" + c.Description + "\n"
				*/
				s += "该调查员" + c.Thought + "\n"
				s += "该调查员的重要之人是" + c.Person + "因为" + c.Reason + "\n"
				s += "该调查员的意义非凡之地是" + c.Place + "\n"
				s += "该调查员的宝贵之物是" + c.Treasure + "\n"
				s += "该调查员是一个" + c.Feature
			}

			bm.Reply(u, s)
			return true
		},
		Help: &botmaid.Help{
			Menu:  "char",
			Help:  "生成随机人物卡",
			Names: []string{"char"},
			Usage: "使用方法：char [选项]，默认仅生成属性",
			SetFlag: func(f *pflag.FlagSet) {
				f.BoolP("full", "f", false, "生成完整角色背景")
				f.Int("num", 1, "生成角色卡数量（仅简略生成时可用）")
			},
		},
	})

	bm.AddCommand(&botmaid.Command{
		Do: func(u *botmaid.Update, f *pflag.FlagSet) bool {
			fmtRoll := "%v的%v是——%v！"

			w := ""
			e := ""

			if len(f.Args()) == 1 {
				e = "1d100"
			}

			if len(f.Args()) == 2 {
				e = f.Args()[1]
			}

			if len(f.Args()) == 3 {
				w = f.Args()[1]
				e = f.Args()[2]
			}

			es := strings.Split(e, "d")
			if len(es) == 2 {
				a, err := strconv.Atoi(es[0])
				if err == nil && a > 1 {
					b, err := strconv.Atoi(es[1])
					if err == nil && b > 0 {
						if a > 256 {
							reply(u, "投掷次数不能超过 256。")
							return true
						}

						s := "("
						sum := 0
						for i := 0; i < a; i++ {
							t := random.Int(1, b)
							s += strconv.Itoa(t)
							sum += t
							if i < a-1 {
								s += ", "
							}
						}
						s += ") = " + strconv.Itoa(sum)

						reply(u, fmt.Sprintf(fmtRoll, bm.At(u.User), w+"骰点结果", s))
						return true
					}
				}
			}

			n, err := strconv.Atoi(e)
			if err != nil || n < 0 {
				ee, err := nyamath.New(e)
				if err != nil && len(f.Args()) == 2 {
					w = f.Args()[1] + "骰点"
					ee, err = nyamath.New("1d100")
				} else {
					if len(es) == 2 {
						_, err := strconv.Atoi(es[1])
						if err == nil && es[0] == "1" {
							w += "骰点"
						} else {
							w += "表达式"
						}
					} else {
						w += "表达式"
					}
				}

				if err == nil {
					reply(u, fmt.Sprintf(fmtRoll, bm.At(u.User), w+"结果", ee.Result.Value))
					return true
				}
			} else {
				diff, _ := f.GetString("diff")
				if diff == "h" {
					n /= 2
				}
				if diff == "e" {
					n /= 5
				}

				bonus, _ := f.GetInt("bonus")
				punish, _ := f.GetInt("punish")
				bp := bonus - punish
				if math.Abs(float64(bp)) > 256 {
					bm.Reply(u, "奖罚骰数量差的绝对值不应超过 256。")
					return true
				}
				r := coc.Check(n, bp)

				s := fmt.Sprintf("%v/%v", n, r.Number)

				if r.Great == coc.GreatSucc {
					s += "，大成功"
				} else if r.Great == coc.GreatFail {
					s += "，大失败"
				} else if r.Level == coc.DiffSucc {
					s += "，困难成功"
				} else if r.Level == coc.ExDiffSucc {
					s += "，极难成功"
				} else if r.Succ == coc.Succ {
					s += "，检定成功"
				} else if r.Succ == coc.Fail {
					s += "，检定失败"
				}

				reply(u, fmt.Sprintf(fmtRoll, bm.At(u.User), w+"检定结果", s))

				if pkMap[u.Chat.ID] != nil {
					pkMap[u.Chat.ID].Results = append(pkMap[u.Chat.ID].Results, pkRollResult{
						User:   u.User,
						Result: r,
					})
					pkResp(u)
				}

				if scMap[u.Chat.ID] != nil {
					if scMap[u.Chat.ID][u.User.ID] != nil {
						scMap[u.Chat.ID][u.User.ID].Result = r
						scResp(u)
					}
				}

				return true
			}

			reply(u, fmt.Sprintf(fmtRoll, u.User.NickName, "随机结果", random.Slice(f.Args()[1:]).(string)))
			return true
		},
		Help: &botmaid.Help{
			Menu:  "roll",
			Help:  "roll 点功能",
			Names: []string{"roll", "r"},
			Usage: "使用方法：roll（或 r） [选项] [说明文字] [数值/表达式/列表]，默认返回 1d100 的结果",
			SetFlag: func(f *pflag.FlagSet) {
				f.BoolP("hide", "h", false, "暗骰")
				f.String("diff", "n", "设置检定的困难程度（普通=n，困难=h，极难=e")
				f.Int("bonus", 0, "设置奖励骰的数量")
				f.Int("punish", 0, "设置惩罚骰的数量")
			},
		},
	})

	/*
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
						Text: fmt.Sprintf(random.String(fmtRoll), u.User.NickName, instruction+"检定结果", message),
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
	*/
}
