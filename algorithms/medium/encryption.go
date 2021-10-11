package medium

import (
	"math"
	"strings"
)

func resolveMatrixSize(s string) (int, int) {
	square := math.Sqrt(float64(len(s)))
	ceil := math.Ceil(square)
	floor := math.Floor(square)

	var rows, columns int

	if ceil*floor < float64(len(s)) {
		rows = int(ceil)
		columns = int(ceil)
	}

	if ceil*floor >= float64(len(s)) {
		rows = int(ceil)
		columns = int(floor)
	}

	return rows, columns
}

func buildMatrix(s string, rows int, columns int) [][]string {
	encs := make([][]string, rows)
	row := 0

	for _, item := range s {
		if row == rows {
			row = 0
		}

		encs[row] = append(encs[row], string(item))

		row += 1
	}

	return encs
}

func Encryption(s string) string {
	s = strings.ReplaceAll(s, " ", "")

	rows, columns := resolveMatrixSize(s)

	encs := buildMatrix(s, rows, columns)

	result := []string{}
	for _, enc := range encs {
		result = append(result, strings.Join(enc, ""))
	}

	return strings.Join(result, " ")
}
