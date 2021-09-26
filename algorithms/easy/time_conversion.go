package easy

func TimeConversion(s string) string {
	hour := s[:2]
	time := s[len(s)-2:]
	conversions := map[string]string{
		"12": "12",
		"01": "13",
		"02": "14",
		"03": "15",
		"04": "16",
		"05": "17",
		"06": "18",
		"07": "19",
		"08": "20",
		"09": "21",
		"10": "22",
		"11": "23",
	}

	t := ""

	if time == "PM" {
		t += conversions[hour]
	} else {
		if hour == "12" {
			t += "00"
		} else {
			t += hour
		}
	}

	t += s[2 : len(s)-2]

	return t
}
