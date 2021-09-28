package easy

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

// parseInput prepares the string input into a workable data structure
func parseInput(inputData string) (int, [][]string) {
	rawInput := []byte(inputData)

	testCases, _ := strconv.Atoi(string(rawInput[0]))

	i := rawInput[1:]

	w := strings.Split(string(i), "\n")

	words := [][]string{}

	t := []string{}

	for _, word := range w {
		if word == "" {
			continue
		}

		t = append(t, word)

		if len(t) == 2 {
			words = append(words, t)
			t = []string{}
		}
	}

	return testCases, words
}

// code_here contains only the solution which runs in O(n)
func code_here(w1 string, w2 string) string {
	l := len(w2)
	i := 0
	s := 0

	if len(w1) < len(w2) || len(w2) < len(w1) {
		s = int(math.Abs(float64(len(w2) - len(w1))))
	}

	for i <= l {
		if i > len(w1)-1 || i > len(w2)-1 {
			break
		}

		if w1[i] != w2[i] {
			s = s + 1
		}

		i = i + 1
	}

	return fmt.Sprintf("%x", s)
}

func StringDistance(inputData string) string {
	_, words := parseInput(inputData)

	d := []string{}

	for _, word := range words {
		d = append(d, code_here(word[0], word[1]))
	}

	return strings.Join(d, "\n")
}
