package Easy

import "strings"

func lengthOfLastWord(s string) int {
	s = strings.TrimSpace(s)

	words := strings.Fields(s)

	return len(words[len(words)-1])
}
