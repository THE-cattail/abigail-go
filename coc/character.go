package coc

import (
	"github.com/catsworld/abigail/nyamath"
	"github.com/catsworld/botmaid/random"
)

// Character includes some attributes of a character.
type Character struct {
	STR, DEX, POW, CON, APP, EDU, SIZ, INT, Luck      int
	Thought, Person, Reason, Place, Treasure, Feature string
	//Appearance string
}

// NewCharacter returns a random character according to the random list of the CoC rule book.
func NewCharacter() Character {
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
	*/

	c := Character{}

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

	/*
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
		"你信仰并祈并一位大能。（选择一位大能）",
		"你认为人类无需上帝。（无神论者或人文学家等）",
		"你认为科学是万能的。（选择一种科学理论）",
		"你认为命中注定。（选择一种现象）",
		"你是社会团体或秘密结社的一员。（选择一个社会团体或秘密结社）",
		"你认为社会坏掉了，而你将成为正义的伙伴。（选择一种应铲除之物）",
		"你认为神秘依然在。（选择一种神秘）",
		"你喜欢政治。（选择一种党派）",
		"你的名言是“金钱就是力量，我的朋友，我将竭尽全力获取我能看到的一切。”",
		"你是一名激进主义者。（选择一种激进主义）",
	})

	c.Person = random.String([]string{
		"父母辈。（选择父亲、母亲、继父或继母）",
		"祖父母辈。（选择祖父、祖母、外祖父或外祖母等）",
		"兄弟姐妹。（选择一位兄弟姐妹）",
		"孩子。（选择儿子或女儿）",
		"另一半。（选择男/女朋友、未婚夫/妻或丈夫/妻子等）",
		"那位指引你人生技能的人。（选择人物并选择一项技能）",
		"青梅竹马。",
		"名人、偶像或者英雄。（选择一位名人、偶像或英雄）",
		"另一名调查员。（选择一名调查员）",
		"一位 NPC。（向 KP 询问并选择一位 NPC）",
	})

	c.Reason = random.String([]string{
		"你欠了 TA 人情。（选择一个场景）",
		"TA 教会了你一些东西。（选择学会的内容）",
		"TA 给了你生命的意义。（选择一个场景）",
		"你曾害了 TA，而现在寻求救赎。（选择一个场景）",
		"你曾和 TA 同甘共苦。（选择一个场景）",
		"你想向 TA 证明自己。（选择证明的内容）",
		"你崇拜 TA。（选择崇拜的内容）",
		"你对 TA 有后悔的感觉。（选择一个场景）",
		"你试图证明你比 TA 更出色。（选择 TA 的缺点）",
		"TA 扰乱了你的人生，而你寻求复仇。（选择一个场景）",
	})

	c.Place = random.String([]string{
		"你最爱的学府。（选择一所学校）",
		"你的故乡。（选择一处地方）",
		"相识初恋之处。（选择一处地方）",
		"静思之地。（选择一处地方）",
		"社交之地。（选择一处地方）",
		"联系你思想信念的场所。（选择一处地方）",
		"重要之人的坟墓。（选择一处地方）",
		"家族所在。（选择一处地方）",
		"生命中最高兴时的所在。（选择一处地方）",
		"工作地点。（选择一处地方）",
	})

	c.Treasure = random.String([]string{
		"与你得意技相关之物。（选择一件物品）",
		"职业必需品。（选择一件物品）",
		"童年的遗留物。（选择一件物品）",
		"逝者遗物。（选择一件物品）",
		"重要之人给予之物。（选择一件物品）",
		"收藏品。（选择一件物品）",
		"你发掘而不知真相的东西。（选择一件物品）",
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

	return c
}
