package bob

import (
	"strings"
	"unicode"
)

func Hey(remark string) string {
	remark = strings.TrimSpace(remark)

	if len(remark) == 0 {
		return "Fine. Be that way!"
	}

	if remark == strings.ToUpper(remark) && strings.ContainsFunc(remark, unicode.IsLetter) {
		if strings.HasSuffix(remark, "?") {
			return "Calm down, I know what I'm doing!"
		}
		return "Whoa, chill out!"
	}

	if strings.HasSuffix(remark, "?") {
		return "Sure."
	}

	return "Whatever."
}
