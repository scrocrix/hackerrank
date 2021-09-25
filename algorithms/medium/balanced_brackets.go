package medium

func IsBalanced(s string) string {
	right := len(s) - 1
	left := 0

	if len(s) == 1 || len(s) == 0 {
		return "NO"
	}

	for left < right {
		o := string(s[left])
		c := string(s[right])

		if (o == "{" && c == "}") || (o == "(" && c == ")") || (o == "[" && c == "]") {
			right = right - 1
			left = left + 1
			continue
		}

		return "NO"
	}

	return "YES"
}
