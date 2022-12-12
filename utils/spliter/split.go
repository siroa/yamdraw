package spliter

import (
	"regexp"
)

func ProcedureSplit(pr string) []string {
	re := regexp.MustCompile(`->|.>`)
	sp := Split(pr, -1, re)
	return sp
}

func Split(s string, n int, re *regexp.Regexp) []string {

	if n == 0 {
		return nil
	}

	if len(s) == 0 {
		return []string{""}
	}

	matches := re.FindAllStringIndex(s, n)
	matNum := len(matches)
	strings := make([]string, 0, 2*matNum-1)

	beg := 0
	end := 0
	for _, match := range matches {
		// nの回数かつ戻り値の個数がn-1以上だったらブレイク
		if n > 0 && len(strings) >= n-1 {
			break
		}

		end = match[0]
		if match[1] != 0 {
			strings = append(strings, s[beg:end])
		}
		beg = match[1]
		if match[1] != 0 {
			strings = append(strings, s[end:beg])
		}
	}

	if end != len(s) {
		strings = append(strings, s[beg:])
	}

	return strings
}
