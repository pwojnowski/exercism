package parsinglogfiles

import (
	"fmt"
	"regexp"
)

func IsValidLine(text string) bool {
	if text == "" {
		return false
	}
	re := regexp.MustCompile(`^\[(TRC|DBG|INF|WRN|ERR|FTL)\]`)
	return re.MatchString(text)
}

func SplitLogLine(text string) []string {
	re := regexp.MustCompile(`<[~*=-]*?>`)
	return re.Split(text, -1)
}

func CountQuotedPasswords(lines []string) int {
	re := regexp.MustCompile(`(?i)".*?password.*?"`)
	total := 0
	for _, line := range lines {
		if matches := re.FindAllString(line, -1); matches != nil {
			total += len(matches)
		}
	}
	return total
}

func RemoveEndOfLineText(text string) string {
	re := regexp.MustCompile(`end-of-line\d+`)
	return re.ReplaceAllLiteralString(text, "")
}

func TagWithUserName(lines []string) []string {
	re := regexp.MustCompile(`User\s+(\S+)`)
	processed := make([]string, 0)
	for _, line := range lines {
		matches := re.FindStringSubmatch(line)
		if matches != nil {
			line = fmt.Sprintf("[USR] %s %s", matches[1], line)
		}
		processed = append(processed, line)
	}
	return processed
}
