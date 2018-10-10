package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/catsworld/api"
	"github.com/catsworld/botmaid"
	"github.com/catsworld/coc"
	"github.com/catsworld/expression"
	"github.com/catsworld/random"
	"github.com/catsworld/slices"
)

var (
	formatRollMad = []string{
		"%s，汝已陷入了临时疯狂：\n",
	}

	formatRollCharacter = []string{
		"这就是%s新的调查员哦。\n",
		"%s家新来的调查员？不过如此，呵呵呵。\n",
	}

	wordRollProperty = []string{
		"这是你的属性：",
	}

	wordRollDescription = []string{
		"你是这样的：",
	}

	wordRollThought = []string{
		"你的信念是这样的：",
	}

	wordRollPerson = []string{
		"你的重要之人是",
	}

	wordRollReason = []string{
		"因为",
	}

	wordRollPlace = []string{
		"你的意义非凡之地是",
	}

	wordRollTreasure = []string{
		"你的宝贵之物是",
	}

	wordRollFeature = []string{
		"你是一个",
	}

	formatRoll = []string{
		"%s的%s是——%v！",
	}

	wordCheckSuccess = []string{
		"检定成功",
	}

	wordCheckFailure = []string{
		"检定失败",
	}

	wordCheckHardSuccess = []string{
		"困难成功",
	}

	wordCheckVeryHardSuccess = []string{
		"极难成功",
	}

	wordCheckBigSuccess = []string{
		"☆大☆成☆功☆",
	}

	wordCheckBigFailure = []string{
		"大失败",
		"大失败 Q皿Q",
	}

	wordCheckEgg = []botmaid.Word{
		botmaid.Word{
			Word:   "",
			Weight: 19,
		},
		botmaid.Word{
			Word:   "(๑•̀ㅂ•́)✧",
			Weight: 1,
		},
	}

	hide = map[int64]bool{}
)

func init() {
	botmaid.AddCommand(&commands, rollMad, 5)
	botmaid.AddCommand(&commands, rollCharacter, 5)
	botmaid.AddCommand(&commands, initRoll, 5)
	botmaid.AddCommand(&commands, rollXdy, 5)
	botmaid.AddCommand(&commands, rollCalc, 5)
	botmaid.AddCommand(&commands, roll, 5)
	botmaid.AddCommand(&commands, rollList, 5)
	botmaid.AddCommand(&commands, ww, 5)
}

func roll(e *api.Event, b *botmaid.Bot) bool {
	args := botmaid.SplitCommand(e.Message.Text)
	instruction := ""
	num := ""
	if b.IsCommand(e, "roll", "r") && len(args) == 1 {
		num = "0"
	}
	if b.IsCommand(e, "roll", "r") && len(args) == 2 {
		_, err := strconv.Atoi(args[1])
		if err != nil {
			instruction = args[1]
			num = "0"
		} else {
			num = args[1]
		}
	}
	if b.IsCommand(e, "roll", "r") && len(args) == 3 {
		instruction = args[1]
		num = args[2]
	}
	n, err := strconv.Atoi(num)
	if err != nil {
		return false
	}
	result := coc.Check(int64(n))
	message := strconv.FormatInt(result.Number, 10)
	if n > 0 && n < 100 {
		message = strconv.Itoa(n) + "/" + message
	}
	if result.BigSuccess() {
		message += "！" + random.String(wordCheckBigSuccess)
	} else if result.BigFailure() {
		message += "！" + random.String(wordCheckBigFailure)
	} else if result.HardSuccess() {
		message += "！" + random.String(wordCheckHardSuccess)
	} else if result.VeryHardSuccess() {
		message += "！" + random.String(wordCheckVeryHardSuccess)
	} else if result.Success() {
		message += "！" + random.String(wordCheckSuccess)
	} else if result.Failure() {
		message += "！" + random.String(wordCheckFailure)
	}
	message = fmt.Sprintf(random.String(formatRoll), e.Sender.NickName, instruction+"检定结果", message)
	if result.Success() {
		message += botmaid.RandomWordWithWeight(wordCheckEgg)
	}
	send(&api.Event{
		Message: &api.Message{
			Text: message,
		},
		Place: e.Place,
	}, b, hide[e.ID])
	if _, ok := pkMap[e.Place.ID]; ok && pkMap[e.Place.ID].Status && n > 0 && n < 100 {
		pkMap[e.Place.ID].Results = append(pkMap[e.Place.ID].Results, pkRollResult{
			User:   e.Sender,
			Result: result,
		})
		pkResp(e, b)
	}
	if _, ok := scMap[e.Place.ID]; ok {
		if _, ok := scMap[e.Place.ID][e.Sender.ID]; ok && scMap[e.Place.ID][e.Sender.ID].Status && n > 0 && n < 100 {
			scMap[e.Place.ID][e.Sender.ID].Result = result
			scResp(e, b)
		}
	}
	return true
}

func rollXdy(e *api.Event, b *botmaid.Bot) bool {
	args := botmaid.SplitCommand(e.Message.Text)
	instruction := ""
	exp := ""
	if b.IsCommand(e, "roll", "r") && len(args) == 2 {
		exp = args[1]
	}
	if b.IsCommand(e, "roll", "r") && len(args) == 3 {
		instruction = args[1]
		exp = args[2]
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
		result = strconv.FormatInt(random.Rand(int64(1), int64(bb)), 10)
	} else {
		result = "("
		sum := int64(0)
		for i := 0; i < a; i++ {
			t := random.Rand(int64(1), int64(bb))
			result += strconv.FormatInt(t, 10)
			sum += t
			if i < a-1 {
				result += ", "
			}
		}
		result += ") = " + strconv.FormatInt(sum, 10)
	}
	send(&api.Event{
		Message: &api.Message{
			Text: fmt.Sprintf(random.String(formatRoll), e.Sender.NickName, instruction+"检定结果", result),
		},
		Place: e.Place,
	}, b, hide[e.ID])
	return true
}

func rollCalc(e *api.Event, b *botmaid.Bot) bool {
	args := botmaid.SplitCommand(e.Message.Text)
	instruction := "计算结果"
	exp := ""
	if b.IsCommand(e, "roll", "r") && len(args) == 2 {
		exp = args[1]
	}
	if b.IsCommand(e, "roll", "r") && len(args) == 3 {
		instruction = args[1]
		exp = args[2]
	}
	if exp == "" {
		return false
	}
	_, err := strconv.Atoi(exp)
	if err == nil {
		return false
	}
	ee, err := expression.New(exp)
	if err != nil {
		return false
	}
	send(&api.Event{
		Message: &api.Message{
			Text: fmt.Sprintf(random.String(formatRoll), e.Sender.NickName, instruction, ee.Result()),
		},
		Place: e.Place,
	}, b, hide[e.ID])
	return true
}

func rollList(e *api.Event, b *botmaid.Bot) bool {
	if b.IsCommand(e, "roll", "r") {
		args := botmaid.SplitCommand(e.Message.Text)
		send(&api.Event{
			Message: &api.Message{
				Text: fmt.Sprintf(random.String(formatRoll), e.Sender.NickName, "随机结果", args[random.Rand(1, int64(len(args)-1))]),
			},
			Place: e.Place,
		}, b, hide[e.ID])
		return true
	}
	return false
}

func rollMad(e *api.Event, b *botmaid.Bot) bool {
	if b.IsCommand(e, "mad") {
		send(&api.Event{
			Message: &api.Message{
				Text: fmt.Sprintf(random.String(formatRollMad), e.Sender.NickName) + coc.RollMad(),
			},
			Place: e.Place,
		}, b, hide[e.ID])
		return true
	}
	return false
}

func rollCharacter(e *api.Event, b *botmaid.Bot) bool {
	if b.IsCommand(e, "character", "char") {
		c := coc.RollCharacter()
		message := random.String(wordRollProperty) + "\n"
		message += "力量" + strconv.Itoa(c.STR) + " "
		message += "敏捷" + strconv.Itoa(c.DEX) + " "
		message += "意志" + strconv.Itoa(c.POW) + " "
		message += "体质" + strconv.Itoa(c.CON) + " "
		message += "外貌" + strconv.Itoa(c.APP) + " "
		message += "教育" + strconv.Itoa(c.EDU) + " "
		message += "体型" + strconv.Itoa(c.SIZ) + " "
		message += "智力" + strconv.Itoa(c.INT) + " "
		message += "幸运" + strconv.Itoa(c.Luck)
		args := botmaid.SplitCommand(e.Message.Text)
		if len(args) > 1 && slices.In(args[1], "-full", "-f") {
			message += "\n"
			//message += random.String(wordRollDescription) + c.Description + "\n"
			message += random.String(wordRollThought) + c.Thought + "\n"
			message += random.String(wordRollPerson) + c.Person + random.String(wordRollReason) + c.Reason + "\n"
			message += random.String(wordRollPlace) + c.Place + "\n"
			message += random.String(wordRollTreasure) + c.Treasure + "\n"
			message += random.String(wordRollFeature) + c.Feature
		}
		send(&api.Event{
			Message: &api.Message{
				Text: fmt.Sprintf(random.String(formatRollCharacter), e.Sender.NickName) + message,
			},
			Place: e.Place,
		}, b, hide[e.ID])
		return true
	}
	return false
}

func ww(e *api.Event, b *botmaid.Bot) bool {
	args := botmaid.SplitCommand(e.Message.Text)
	if b.IsCommand(e, "ww") && len(args) > 1 {
		instruction := ""
		exp := ""
		if b.IsCommand(e, "ww") && len(args) == 2 {
			exp = args[1]
		}
		if b.IsCommand(e, "ww") && len(args) == 3 {
			instruction = args[1]
			exp = args[2]
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
		rollSliceSlice := [][]int64{}
		ans := 0
		remain := num
		for remain > 0 {
			num = remain
			remain = 0
			rollSlice := []int64{}
			for i := 0; i < num; i++ {
				result := random.Rand(1, 10)
				rollSlice = append(rollSlice, result)
				if result > 7 {
					ans++
				}
				if result >= int64(a) {
					remain++
				}
			}
			rollSliceSlice = append(rollSliceSlice, rollSlice)
		}
		message := ""
		for i, rollSlice := range rollSliceSlice {
			if i != 0 {
				message += " + "
			}
			message += "("
			for i, v := range rollSlice {
				if i != 0 {
					message += ", "
				}
				message += strconv.FormatInt(v, 10)
			}
			message += ")"
		}
		if plus > 0 {
			message += " + " + strconv.Itoa(plus) + " = " + strconv.Itoa(ans) + " + " + strconv.Itoa(plus)
		}
		message += " = " + strconv.Itoa(ans+plus)
		send(&api.Event{
			Message: &api.Message{
				Text: fmt.Sprintf(random.String(formatRoll), e.Sender.NickName, instruction+"检定结果", message),
			},
			Place: e.Place,
		}, b, hide[e.ID])
		return true
	}
	return false
}

func initRoll(e *api.Event, b *botmaid.Bot) bool {
	args := botmaid.SplitCommand(e.Message.Text)
	if b.IsCommand(e, "roll", "r", "ww") && len(args) > 1 && slices.In(args[1], "-hide", "-h") {
		hide[e.ID] = true
		e.Message.Text = strings.Replace(e.Message.Text, "-h", "", 1)
		e.Message.Text = strings.Replace(e.Message.Text, "-hide", "", 1)
	}
	return false
}
