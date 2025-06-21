//go:build !solution

package varjoin

import (
		"strings"
)

func Join(sep string, args ...string) string {
	if len(args) == 0 {
		return ""
	}

	output := strings.Builder{}
	output.Grow(TotalLen(sep, args...))

	for i, s := range args {
		if i > 0 {
			output.WriteString(sep)
		}
		output.WriteString(s)
	}

	return output.String()
}

func TotalLen(sep string, args ...string) int {
	sep_len := len(sep)
	var total_len int
	for _, s := range args {
		total_len += len(s) + sep_len
	}

	return total_len - sep_len
}
