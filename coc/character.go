package coc

import (
	"github.com/Pallinder/go-randomdata"
	"github.com/catsworld/abigail/nyamath"
	"github.com/catsworld/botmaid/random"
	"github.com/goroom/rand"
	"github.com/mattn/go-gimei"
)

// Character includes some attributes of a character.
type Character struct {
	STR, DEX, POW, CON, APP, EDU, SIZ, INT, Luck      int
	Age                                               int
	AgeMention                                        string
	Name, Class                                       string
	Thought, Person, Reason, Place, Treasure, Feature string
	//Appearance string
}

func improvingCheck(a int) int {
	r := a
	t := random.Int(1, 100)
	if t > r {
		r += random.Int(1, 10)
	}
	if r > 99 {
		r = 99
	}
	return r
}

// NewCharacter returns a random character according to the random list of the CoC rule book.
func NewCharacter(full bool) *Character {
	c := &Character{}

	e, _ := nyamath.New("3d6*5")
	c.STR = e.Result.Value
	e, _ = nyamath.New("3d6*5")
	c.DEX = e.Result.Value
	e, _ = nyamath.New("3d6*5")
	c.POW = e.Result.Value
	e, _ = nyamath.New("3d6*5")
	c.CON = e.Result.Value
	e, _ = nyamath.New("3d6*5")
	c.APP = e.Result.Value
	e, _ = nyamath.New("(2d6+6)*5")
	c.EDU = e.Result.Value
	e, _ = nyamath.New("(2d6+6)*5")
	c.SIZ = e.Result.Value
	e, _ = nyamath.New("(2d6+6)*5")
	c.INT = e.Result.Value
	e, _ = nyamath.New("3d6*5")
	c.Luck = e.Result.Value

	if full {
		c.Age = random.Int(15, 89)
		if c.Age < 20 {
			c.AgeMention = "力量和体型合计减5点"
			c.EDU -= 5
			e, _ := nyamath.New("3d6*5")
			if e.Result.Value > c.Luck {
				c.Luck = e.Result.Value
			}
		} else if c.Age < 40 {
			c.EDU = improvingCheck(c.EDU)
		} else if c.Age < 50 {
			c.EDU = improvingCheck(c.EDU)
			c.EDU = improvingCheck(c.EDU)
			c.AgeMention = "力量体质敏捷合计减5点"
			c.APP -= 5
		} else if c.Age < 60 {
			c.EDU = improvingCheck(c.EDU)
			c.EDU = improvingCheck(c.EDU)
			c.EDU = improvingCheck(c.EDU)
			c.AgeMention = "力量体质敏捷合计减10点"
			c.APP -= 10
		} else if c.Age < 70 {
			c.EDU = improvingCheck(c.EDU)
			c.EDU = improvingCheck(c.EDU)
			c.EDU = improvingCheck(c.EDU)
			c.EDU = improvingCheck(c.EDU)
			c.AgeMention = "力量体质敏捷合计减20点"
			c.APP -= 15
		} else if c.Age < 80 {
			c.EDU = improvingCheck(c.EDU)
			c.EDU = improvingCheck(c.EDU)
			c.EDU = improvingCheck(c.EDU)
			c.EDU = improvingCheck(c.EDU)
			c.AgeMention = "力量体质敏捷合计减40点"
			c.APP -= 20
		} else if c.Age < 90 {
			c.EDU = improvingCheck(c.EDU)
			c.EDU = improvingCheck(c.EDU)
			c.EDU = improvingCheck(c.EDU)
			c.EDU = improvingCheck(c.EDU)
			c.AgeMention = "力量体质敏捷合计减80点"
			c.APP -= 25
		}

		c.Name = random.String([]string{rand.GetRand().ChineseName(), randomdata.FullName(randomdata.RandomGender), gimei.NewName().Kanji()})
		c.Class = random.String([]string{
			"会计师",
			"杂技演员",
			"演员",
			"事务所侦探",
			"精神病医生",
			"动物训练师",
			"文物学家",
			"古董商",
			"考古学家",
			"建筑师",
			"艺术家",
			"精神病院看护",
			"刺客",
			"运动员",
			"作家",
			"特技飞行员",
			"银行劫匪",
			"酒保",
			"猎人",
			"书商",
			"打手",
			"猎人",
			"拳击手",
			"摔跤手",
			"窃贼",
			"管家",
			"仆人",
			"私人司机",
			"神职人员",
			"计算机程序员",
			"计算机工程师",
			"黑客",
			"欺诈师",
			"牛仔",
			"工匠",
			"罪犯",
			"教团首领",
			"除魅师",
			"设计师",
			"业余艺术爱好者",
			"潜水员",
			"医生",
			"流浪者",
			"司机",
			"编辑",
			"政府官员",
			"工程师",
			"艺人",
			"探险家",
			"农民",
			"联邦探员",
			"赃物贩子",
			"消防员",
			"驻外记者",
			"法医",
			"赝造者",
			"赌徒",
			"黑帮",
			"绅士/淑女",
			"游民",
			"勤杂护工",
			"记者",
			"法官",
			"实验室助理",
			"工人",
			"律师",
			"图书馆管理员",
			"伐木工",
			"技师",
			"军官",
			"矿工",
			"传教士",
			"登山家",
			"博物馆管理员",
			"音乐家",
			"护士",
			"神秘学家",
			"旅行家",
			"超心理学家",
			"药剂师",
			"摄影师",
			"摄影记者",
			"飞行员",
			"警察",
			"警探",
			"私家侦探",
			"教授",
			"淘金客",
			"性工作者",
			"精神病学家",
			"心理学家",
			"心理分析学家",
			"通讯记者",
			"研究员",
			"海员",
			"推销员",
			"科学家",
			"秘书",
			"店老板",
			"走私者",
			"士兵",
			"海军陆战队士兵",
			"间谍",
			"混混",
			"学生",
			"实习生",
			"替身演员",
			"出租车司机",
			"暴徒",
			"部落成员",
			"殡葬师",
			"工会活动家",
			"男仆",
			"服务生",
			"白领工人",
			"狂热者",
			"饲养员",
		})

		/*
			charAppearance := []string{
				"结实的",
				"英俊的",
				"笨拙的",
				"机灵的",
				"迷人的",
				"娃娃脸",
				"聪明的",
				"邋遢的",
				"死人脸",
				"肮脏的",
				"耀眼的",
				"书呆子",
				"年轻的",
				"疲倦脸",
				"肥胖的",
				"啤酒肚",
				"长头发",
				"苗条的",
				"优雅的",
				"稀烂的",
				"矮壮的",
				"苍白的",
				"阴沉的",
				"平庸的",
				"乐观的",
				"棕褐色",
				"皱纹的",
				"古板的",
				"狐臭的",
				"狡猾的",
				"健壮的",
				"娇俏的",
				"筋肉人",
				"魁梧的",
				"迟钝的",
				"虚弱的",
			}
			d1 := random.Int(0, len(charAppearance)-1)
			d2 := 0
			d3 := 0
			for {
				d2 = random.Int(0, len(charAppearance)-1)
				if d2 != d1 {
					break
				}
			}
			for {
				d3 = random.Int(0, len(charAppearance)-1)
				if d3 != d1 && d3 != d2 {
					break
				}
			}
		*/

		//c.Appearance = charAppearance[d1] + " " + charAppearance[d2] + " " + charAppearance[d3]

		c.Thought = random.String([]string{
			"信仰并祈并一位大能。（选择一位大能）",
			"认为人类无需上帝。（无神论者或人文学家等）",
			"认为科学是万能的。（选择一种科学理论）",
			"认为命中注定。（选择一种现象）",
			"是社会团体或秘密结社的一员。（选择一个社会团体或秘密结社）",
			"认为社会坏掉了，而自己将成为正义的伙伴。（选择一种应铲除之物）",
			"认为神秘依然在。（选择一种神秘）",
			"喜欢政治。（选择一种党派）",
			"的名言是“金钱就是力量，我的朋友，我将竭尽全力获取我能看到的一切。”",
			"是一名激进主义者。（选择一种激进主义）",
		})

		c.Person = random.String([]string{
			"父母辈。（选择父亲、母亲、继父或继母）",
			"祖父母辈。（选择祖父、祖母、外祖父或外祖母等）",
			"兄弟姐妹。（选择一位兄弟姐妹）",
			"孩子。（选择儿子或女儿）",
			"另一半。（选择男/女朋友、未婚夫/妻或丈夫/妻子等）",
			"那位指引自己人生技能的人。（选择人物并选择一项技能）",
			"青梅竹马。",
			"名人、偶像或者英雄。（选择一位名人、偶像或英雄）",
			"另一名调查员。（选择一名调查员）",
			"一位 NPC。（向 KP 询问并选择一位 NPC）",
		})

		c.Reason = random.String([]string{
			"自己欠了 TA 人情。（选择一个场景）",
			"TA 教会了自己一些东西。（选择学会的内容）",
			"TA 给了自己生命的意义。（选择一个场景）",
			"自己曾害了 TA，而现在寻求救赎。（选择一个场景）",
			"自己曾和 TA 同甘共苦。（选择一个场景）",
			"自己想向 TA 证明自己。（选择证明的内容）",
			"自己崇拜 TA。（选择崇拜的内容）",
			"自己对 TA 有后悔的感觉。（选择一个场景）",
			"自己试图证明自己比 TA 更出色。（选择 TA 的缺点）",
			"TA 扰乱了自己的人生，而自己寻求复仇。（选择一个场景）",
		})

		c.Place = random.String([]string{
			"自己最爱的学府。（选择一所学校）",
			"自己的故乡。（选择一处地方）",
			"相识初恋之处。（选择一处地方）",
			"静思之地。（选择一处地方）",
			"社交之地。（选择一处地方）",
			"联系自己思想信念的场所。（选择一处地方）",
			"重要之人的坟墓。（选择一处地方）",
			"家族所在。（选择一处地方）",
			"生命中最高兴时的所在。（选择一处地方）",
			"工作地点。（选择一处地方）",
		})

		c.Treasure = random.String([]string{
			"与自己得意技相关之物。（选择一件物品）",
			"职业必需品。（选择一件物品）",
			"童年的遗留物。（选择一件物品）",
			"逝者遗物。（选择一件物品）",
			"重要之人给予之物。（选择一件物品）",
			"收藏品。（选择一件物品）",
			"自己发掘而不知真相的东西。（选择一件物品）",
			"体育用品。（选择一件物品）",
			"武器。（选择一件物品）",
			"宠物。（选择一件物品）",
		})

		c.Feature = random.String([]string{
			"慷慨大方的人。",
			"善待动物的人。",
			"梦想家。",
			"享乐主义者。",
			"赌徒，冒险家。",
			"好厨子，好吃货。",
			"女人缘/万人迷。",
			"忠心在我的人。",
			"拥有好名头的人。",
			"雄心壮志的人。",
		})
	}

	return c
}
