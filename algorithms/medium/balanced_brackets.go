package medium

import (
	"strings"
)

func IsBalanced(s string) string {
	var t []string

	if (!strings.Contains(s, "[") && !strings.Contains(s, "]") &&
		!strings.Contains(s, "{") && !strings.Contains(s, "}") &&
		!strings.Contains(s, "(") && !strings.Contains(s, ")")) ||
		len(s) == 0 {
		return "NO"
	}

	for _, i := range s {
		ii := string(i)

		if ii == "(" || ii == "{" || ii == "[" {
			t = append(t, ii)
			continue
		}

		if ii == ")" || ii == "}" || ii == "]" {
			if len(t) == 0 {
				return "NO"
			}

			tt := t[len(t)-1]

			if (tt == "(" && ii == ")") || (tt == "[" && ii == "]") || (tt == "{" && ii == "}") {
				t = t[:len(t)-1]
				continue
			}

			return "NO"
		}
	}

	if len(t) == 0 {
		return "YES"
	}

	return "NO"
}
