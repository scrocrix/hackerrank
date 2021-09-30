package medium

func TimeInWords(h int32, m int32) string {
	mm := map[int32]string{
		60: "o' clock",
		59: "one minute",
		58: "two minutes",
		57: "three minutes",
		56: "four minutes",
		55: "five minutes",
		54: "six minutes",
		53: "seven minutes",
		52: "eight minutes",
		51: "nine minutes",
		50: "ten minutes",
		49: "eleven minutes",
		48: "twelve minutes",
		47: "thirteen minutes",
		46: "fourteen minutes",
		45: "quarter",
		44: "sixteen minutes",
		43: "seventeen minutes",
		42: "eighteen minutes",
		41: "nineteen minutes",
		40: "twenty minutes",
		39: "twenty one minutes",
		38: "twenty two minutes",
		37: "twenty three minutes",
		36: "twenty four minutes",
		35: "twenty five minutes",
		34: "twenty six minutes",
		33: "twenty seven minutes",
		32: "twenty eight minutes",
		31: "twenty nine minutes",
		30: "half",
		29: "twenty nine minutes",
		28: "twenty eight minutes",
		27: "twenty seven minutes",
		26: "twenty six minutes",
		25: "twenty five minutes",
		24: "twenty four minutes",
		23: "twenty three minutes",
		22: "twenty two minutes",
		21: "twenty one minutes",
		20: "twenty minutes",
		19: "nineteen minutes",
		18: "eighteen minutes",
		17: "seventeen minutes",
		16: "sixteen minutes",
		15: "quarter",
		14: "fourteen minutes",
		13: "thirteen minutes",
		12: "twelve minutes",
		11: "eleven minutes",
		10: "ten minutes",
		9:  "nine minutes",
		8:  "eight minutes",
		7:  "seven minutes",
		6:  "six minutes",
		5:  "five minutes",
		4:  "four minutes",
		3:  "three minutes",
		2:  "two minutes",
		1:  "one minute",
	}

	hh := map[int32]string{
		1:  "one",
		2:  "two",
		3:  "three",
		4:  "four",
		5:  "five",
		6:  "six",
		7:  "seven",
		8:  "eight",
		9:  "nine",
		10: "ten",
		11: "eleven",
		12: "twelve",
	}

	md := 60 - m

	time := ""

	var mid string
	if m > 30 {
		mid = "to"
		h = h + 1
	}

	if m != 00 && m <= 30 {
		mid = "past"
	}

	if m == 00 {
		time = hh[h] + " " + mid + "" + mm[md]
	} else {
		time = mm[md] + " " + mid + " " + hh[h]
	}

	return time
}
